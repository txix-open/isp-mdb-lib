package diff

import (
	"reflect"
	"regexp"

	"github.com/integration-system/bellows"
	"github.com/integration-system/go-cmp/cmp"
)

const (
	Add    Operation = "ADD"
	Delete Operation = "DELETE"
	Change Operation = "CHANGE"

	ArrayAdd    Operation = "ARRAY_ADD"
	ArrayDelete Operation = "ARRAY_DELETE"
	ArrayChange Operation = "ARRAY_CHANGE"
	ArraySwap   Operation = "ARRAY_SWAP"
)

var mapIndexRegexp = regexp.MustCompile(`\.\[\d+]`)

type (
	Operation string

	DiffDescriptor struct {
		OldValue       interface{} `json:"oldValue,omitempty"`
		NewValue       interface{} `json:"newValue,omitempty"`
		Path           string      `json:"path"`
		Operation      Operation   `json:"operation"`
		AdditionalData interface{} `json:"additionalData,omitempty"`
		OldIndex       *int        `json:"oldIndex,omitempty"`
		NewIndex       *int        `json:"newIndex,omitempty"`
	}

	Delta         []*DiffDescriptor
	resultHandler func(diff *DiffDescriptor)
)

func EvalDiff(left, right map[string]interface{}, opts ...Option) (bool, Delta) {
	c := NewDiffCollector(opts...)
	return cmp.Equal(left, right, c), c.Delta
}

func FlattenDelta(delta Delta) map[string]*DiffDescriptor {
	result := make(map[string]*DiffDescriptor, len(delta)*2)
	resultHandler := resultHandler(func(diff *DiffDescriptor) {
		result[diff.Path] = diff
	})

	extensionDelta(delta, resultHandler)
	return result
}

func ExtensionDelta(delta Delta) Delta {
	result := make([]*DiffDescriptor, 0)
	resultHandler := resultHandler(func(diff *DiffDescriptor) {
		result = append(result, diff)
	})
	extensionDelta(delta, resultHandler)
	return result
}

func ReplaceArray(delta Delta) Delta {
	result := make([]*DiffDescriptor, len(delta))
	for i, descriptor := range delta {
		descriptor.Path = mapIndexRegexp.ReplaceAllStringFunc(descriptor.Path, func(s string) string { return "" })
		result[i] = descriptor
	}
	return result
}

func extensionDelta(delta Delta, resultHandler resultHandler) {
	for _, desc := range delta {
		switch {
		case (desc.Operation == Add || desc.Operation == ArrayAdd) && desc.NewValue != nil:
			rt := reflect.TypeOf(desc.NewValue)
			switch rt.Kind() {
			case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
				flattenValue := bellows.Flatten(desc.NewValue)
				for path, value := range flattenValue {
					newPath := getNewPath(desc.Path, path, rt.Kind())
					resultHandler(&DiffDescriptor{NewValue: value, Path: newPath, Operation: Add})
				}
			default:
				resultHandler(desc)
			}
		case (desc.Operation == Delete || desc.Operation == ArrayDelete) && desc.OldValue != nil:
			rt := reflect.TypeOf(desc.OldValue)
			switch rt.Kind() {
			case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
				flattenValue := bellows.Flatten(desc.OldValue)
				for path, value := range flattenValue {
					newPath := getNewPath(desc.Path, path, rt.Kind())
					resultHandler(&DiffDescriptor{OldValue: value, Path: newPath, Operation: Delete})
				}
			default:
				resultHandler(desc)
			}
		case desc.Operation == ArrayChange || desc.Operation == Change:
			oldValueByPath := make(map[string]interface{})
			if desc.OldValue != nil {
				rt := reflect.TypeOf(desc.OldValue)
				switch rt.Kind() {
				case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
					flattenValue := bellows.Flatten(desc.OldValue)
					for path, value := range flattenValue {
						newPath := getNewPath(desc.Path, path, rt.Kind())
						oldValueByPath[newPath] = value
					}
				default:
					resultHandler(desc)
					continue
				}
			}
			if desc.NewValue != nil {
				rt := reflect.TypeOf(desc.NewValue)
				switch rt.Kind() {
				case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
					flattenValue := bellows.Flatten(desc.NewValue)
					for path, newValue := range flattenValue {
						newPath := getNewPath(desc.Path, path, rt.Kind())
						oldValue, found := oldValueByPath[newPath]
						if found {
							delete(oldValueByPath, newPath)
							if oldValue == newValue {
								continue
							}
							resultHandler(&DiffDescriptor{NewValue: newValue, OldValue: oldValue, Path: newPath, Operation: Change})
						} else {
							resultHandler(&DiffDescriptor{NewValue: newValue, Path: newPath, Operation: Add})
						}
					}
				default:
					resultHandler(desc)
					continue
				}
			}
			for path, oldValue := range oldValueByPath {
				resultHandler(&DiffDescriptor{OldValue: oldValue, Path: path, Operation: Delete})
			}
		default:
			resultHandler(desc)
		}
	}
}

func getNewPath(base string, path string, kind reflect.Kind) string {
	switch kind {
	case reflect.Slice, reflect.Array:
		return base + path
	default:
		return base + "." + path
	}
}
