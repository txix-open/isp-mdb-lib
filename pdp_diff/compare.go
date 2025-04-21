package pdp_diff

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/txix-open/bellows"
	"github.com/txix-open/isp-mdb-lib/diff"
)

func diffObjects(path string, oldObj map[string]any, newObj map[string]any) (diff.Delta, error) {
	out := make(diff.Delta, 0)
	keys := make(map[string]bool)
	for k := range oldObj {
		keys[k] = true
	}
	for k := range newObj {
		keys[k] = true
	}

	for key := range keys {
		fullPath := key
		if path != "" {
			fullPath = fmt.Sprintf("%s.%s", path, key)
		}
		if len(strings.Split(fullPath, ".")) > 3 {
			return nil, errors.Errorf("too nested data %s", fullPath)
		}
		oldVal, oldOk := oldObj[key]
		newVal, newOk := newObj[key]

		switch {
		case !oldOk && newOk:
			addDiff, err := handleAdd(fullPath, newVal)
			if err != nil {
				return nil, errors.WithMessage(err, "unable to handle add")
			}
			out = append(out, addDiff...)
		case oldOk && !newOk:
			deleteDiff, err := handleDelete(fullPath, oldVal)
			if err != nil {
				return nil, errors.WithMessage(err, "unable to handle delete")
			}
			out = append(out, deleteDiff...)
		default:
			compareDiff, err := handleCompare(fullPath, oldVal, newVal)
			if err != nil {
				return nil, errors.WithMessage(err, "unable to handle compare")
			}
			out = append(out, compareDiff...)
		}
	}

	return out, nil
}

func diffArrayObjects(path string, oldObjArray []map[string]any, newObjArray []map[string]any) (diff.Delta, error) {
	out := make(diff.Delta, 0)
	oldDataByItemId, oldItemIds, err := toItemIdMap(oldObjArray)
	if err != nil {
		return nil, errors.WithMessage(err, "unable to get item_ids from oldData")
	}
	newDataByItemId, newItemIds, err := toItemIdMap(newObjArray)
	if err != nil {
		return nil, errors.WithMessage(err, "unable to get item_ids from newData")
	}

	for _, itemId := range oldItemIds {
		itemPath := fmt.Sprintf("%s.[%s]", path, itemId)
		oldItem := oldDataByItemId[itemId]
		newItem, exists := newDataByItemId[itemId]

		if exists {
			difference, err := diffObjects(itemPath, oldItem, newItem)
			if err != nil {
				return nil, errors.WithMessage(err, "unable to diff objects")
			}
			out = append(out, difference...)
			continue
		}
		deleteDiff, err := diffFlatDelete(itemPath, oldItem)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to diff flat delete")
		}
		out = append(out, deleteDiff...)
	}

	for _, itemId := range newItemIds {
		if _, exists := oldDataByItemId[itemId]; exists {
			continue
		}
		itemPath := fmt.Sprintf("%s.[%s]", path, itemId)
		addDiff, err := diffFlatAdd(itemPath, newDataByItemId[itemId])
		if err != nil {
			return nil, errors.WithMessage(err, "unable to diff flat add")
		}
		out = append(out, addDiff...)
	}

	return out, nil
}

func diffFlatDelete(basePath string, item map[string]any) (diff.Delta, error) {
	if !isFlatObject(item) {
		return nil, errors.New("object is not flat")
	}
	out := make(diff.Delta, 0)
	for k, v := range bellows.Flatten(item) {
		out = append(out, &diff.DiffDescriptor{
			Path:      basePath + "." + k,
			Operation: diff.Delete,
			OldValue:  v,
		})
	}

	return out, nil
}

func diffFlatAdd(basePath string, item map[string]any) (diff.Delta, error) {
	if !isFlatObject(item) {
		return nil, errors.New("object is not flat")
	}
	out := make(diff.Delta, 0)
	for k, v := range bellows.Flatten(item) {
		out = append(out, &diff.DiffDescriptor{
			Path:      basePath + "." + k,
			Operation: diff.Add,
			NewValue:  v,
		})
	}

	return out, nil
}

func handleAdd(path string, val any) (diff.Delta, error) {
	switch {
	case isArrayOfObjects(val):
		oldData, err := toArrayMap(val)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast array object to map")
		}

		return diffArrayObjects(path, nil, oldData)
	case isObject(val):
		obj, _ := val.(map[string]any)
		if !isFlatObject(obj) {
			return nil, errors.New("new object is not flat")
		}

		return diffObjects(path, nil, obj)
	case isString(val):
		newVal, err := castToString(val)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast newValue to string")
		}
		return diff.Delta{&diff.DiffDescriptor{
			Path:      path,
			Operation: diff.Add,
			NewValue:  newVal,
		}}, nil
	default:
		return nil, errors.New("unable to handle add")
	}
}

func handleDelete(path string, val any) (diff.Delta, error) {
	switch {
	case isArrayOfObjects(val):
		arrData, err := toArrayMap(val)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast array object to map")
		}

		return diffArrayObjects(path, arrData, nil)
	case isObject(val):
		obj, _ := val.(map[string]any)
		if !isFlatObject(obj) {
			return nil, errors.New("old object is not flat")
		}

		return diffObjects(path, obj, nil)
	case isString(val):
		oldVal, err := castToString(val)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast oldValue to string")
		}
		return diff.Delta{&diff.DiffDescriptor{
			Path:      path,
			Operation: diff.Delete,
			OldValue:  oldVal,
		}}, nil
	default:
		return nil, errors.New("unable to handle delete")
	}
}

func handleCompare(path string, oldVal any, newVal any) (diff.Delta, error) {
	switch {
	case isArrayOfObjects(oldVal) && isArrayOfObjects(newVal):
		oldData, err := toArrayMap(oldVal)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast oldData array object to map")
		}
		newData, err := toArrayMap(newVal)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast newData array object to map")
		}

		return diffArrayObjects(path, oldData, newData)
	case isObject(oldVal) && isObject(newVal):
		oldObj, _ := oldVal.(map[string]any)
		if !isFlatObject(oldObj) {
			return nil, errors.New("old object is not flat")
		}
		newObj, _ := oldVal.(map[string]any)
		if !isFlatObject(newObj) {
			return nil, errors.New("new object is not flat")
		}

		return diffObjects(path, oldVal.(map[string]any), newVal.(map[string]any))
	case isString(oldVal) && isString(newVal):
		oldValStr, err := castToString(oldVal)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast oldVal to string")
		}
		newValStr, err := castToString(newVal)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to cast newVal to string")
		}
		if oldValStr == newValStr {
			return nil, nil
		}

		return diff.Delta{&diff.DiffDescriptor{
			Path:      path,
			Operation: diff.Change,
			OldValue:  oldValStr,
			NewValue:  newValStr,
		}}, nil
	default:
		return nil, errors.New("unable to handle compare")
	}
}
