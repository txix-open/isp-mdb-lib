package pdp_diff

import (
	"regexp"

	"github.com/txix-open/isp-mdb-lib/diff"
)

const (
	itemIdPath = "ITEM_ID"
)

var mapIndexRegexp = regexp.MustCompile(`\.\[[a-zA-Z0-9]+]`)

func EvalDiff(oldData map[string]any, newData map[string]any) (diff.Delta, error) {
	return diffObjects("", oldData, newData)
}

func ReplaceArray(delta diff.Delta) diff.Delta {
	result := make([]*diff.DiffDescriptor, len(delta))
	for i, descriptor := range delta {
		descriptor.Path = mapIndexRegexp.ReplaceAllStringFunc(descriptor.Path, func(s string) string { return "" })
		result[i] = descriptor
	}
	return result
}
