// nolint:funlen
package pdp_diff_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/txix-open/isp-mdb-lib/diff"
	"github.com/txix-open/isp-mdb-lib/pdp_diff"
)

func TestEvalDiff(t *testing.T) {
	t.Parallel()
	left := map[string]any{
		"array": map[string]any{
			"remove_1": "1",
			"remove_2": "2",
			"ITEM_ID":  "remove_array",
		},
		"change": []any{
			map[string]any{
				"ITEM_ID": "remove_change",
				"remove":  "remove",
				"changed": "1",
			},
		},
		"change3": []any{
			map[string]any{
				"changed": "1",
				"ITEM_ID": "change3",
				"a":       "a",
			},
			map[string]any{
				"remove":  "remove",
				"ITEM_ID": "remove",
			},
			map[string]any{
				"save":    "save",
				"ITEM_ID": "save",
			},
		},
		"saved": []any{
			map[string]any{
				"1":       "1",
				"2":       "2",
				"ITEM_ID": "saved",
			},
		},
	}

	right := map[string]any{
		"array": map[string]any{
			"add_1":   "1",
			"add_2":   "2",
			"ITEM_ID": "add_array",
		},
		"change": []any{
			map[string]any{
				"add":     "add",
				"changed": "2",
				"ITEM_ID": "remove_change",
			},
		},
		"change3": []any{
			map[string]any{
				"changed": "2",
				"ITEM_ID": "change3",
				"a":       "a",
			},
			map[string]any{
				"save":    "save",
				"ITEM_ID": "save",
			},
			map[string]any{
				"add":     "add",
				"ITEM_ID": "add",
			},
		},
		"saved": []any{
			map[string]any{
				"1":       "1",
				"2":       "2",
				"ITEM_ID": "saved",
			},
		},
	}

	expectedEvalDiffPath := map[string]bool{
		"array.ITEM_ID":                      true,
		"array.add_1":                        true,
		"array.add_2":                        true,
		"array.remove_1":                     true,
		"array.remove_2":                     true,
		"change.[remove_change]":             true,
		"array.[add_array].ITEM_ID":          true,
		"array.[add_array].add_1":            true,
		"array.[add_array].add_2":            true,
		"change3.[change3].changed":          true,
		"change3.[remove]":                   true,
		"change3.[remove].remove":            true,
		"change3.[remove].ITEM_ID":           true,
		"change3.[add]":                      true,
		"change3.[add].add":                  true,
		"change3.[add].ITEM_ID":              true,
		"change.[remove_change].changed":     true,
		"change.[remove_change].save.add":    true,
		"change.[remove_change].save.remove": true,
		"change.[remove_change].remove":      true,
		"change.[remove_change].add":         true,
	}
	delta, err := pdp_diff.EvalDiff(left, right)
	require.NoError(t, err)
	for _, descriptor := range delta {
		fmt.Println(descriptor.Path)
		require.True(t, expectedEvalDiffPath[descriptor.Path])
	}

	expectedExtension := map[string]*diff.DiffDescriptor{
		"array.ITEM_ID": {
			Operation: diff.Change, Path: "array.ITEM_ID", OldValue: "remove_array", NewValue: "add_array",
		},
		"array.add_1": {
			Operation: diff.Add, Path: "array.add_1", OldValue: nil, NewValue: "1",
		},
		"array.add_2": {
			Operation: diff.Add, Path: "array.add_2", OldValue: nil, NewValue: "2",
		},
		"array.remove_1": {
			Operation: diff.Delete, Path: "array.remove_1", OldValue: "1", NewValue: nil,
		},
		"array.remove_2": {
			Operation: diff.Delete, Path: "array.remove_2", OldValue: "2", NewValue: nil,
		},
		"change.[remove_change].add": {
			Operation: diff.Add, Path: "change.[remove_change].add", OldValue: nil, NewValue: "add",
		},
		"change.[remove_change].changed": {
			Operation: diff.Change, Path: "change.[remove_change].changed", OldValue: "1", NewValue: "2",
		},
		"change.[remove_change].remove": {
			Operation: diff.Delete, Path: "change.[remove_change].remove", OldValue: "remove", NewValue: nil,
		},
		"change3.[change3].changed": {
			Operation: diff.Change, Path: "change3.[change3].changed", OldValue: "1", NewValue: "2",
		},
		"change3.[remove].remove": {
			Operation: diff.Delete, Path: "change3.[remove].remove", OldValue: "remove", NewValue: nil,
		},
		"change3.[remove].ITEM_ID": {
			Operation: diff.Delete, Path: "change3.[remove].ITEM_ID", OldValue: "remove", NewValue: nil,
		},
		"change3.[add].add": {
			Operation: diff.Add, Path: "change3.[add].add", OldValue: nil, NewValue: "add",
		},
		"change3.[add].ITEM_ID": {
			Operation: diff.Add, Path: "change3.[add].ITEM_ID", OldValue: nil, NewValue: "add",
		},
	}
	for _, descriptor := range delta {
		require.Equal(t, expectedExtension[descriptor.Path], descriptor)
		delete(expectedExtension, descriptor.Path)
	}
	require.Empty(t, expectedExtension)
}

func TestReplaceArray(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	type example struct {
		d       diff.Delta
		newPath string
	}
	for _, e := range []example{
		{
			d: diff.Delta{{
				NewValue: "1", Operation: diff.Add,
				Path: "documents.[1].primary_id",
			}, {
				NewValue: "1", Operation: diff.Add,
				Path: "documents.[2].primary_id",
			}},
			newPath: "documents.primary_id",
		},
		{
			d: diff.Delta{{
				NewValue: "1", Operation: diff.Add,
				Path: "$$cards.9.[0].mdm_obj_id",
			}},
			newPath: "$$cards.9.mdm_obj_id",
		},
		{
			d: diff.Delta{{
				NewValue: "1", Operation: diff.Add,
				Path: "USER.[FIO].first_name",
			}, {
				NewValue: "1", Operation: diff.Add,
				Path: "USER.[SOMEID123].first_name",
			}},
			newPath: "USER.first_name",
		},
		{
			d: diff.Delta{{
				NewValue: "1", Operation: diff.Add,
				Path: "USER.[--123fio].first_name",
			}},
			newPath: "USER.[--123fio].first_name",
		},
	} {
		actual := pdp_diff.ReplaceArray(e.d)
		for _, descriptor := range actual {
			a.Equal(e.newPath, descriptor.Path)
		}
	}
}

func TestTooNestedArrayOfObjects(t *testing.T) {
	_, err := pdp_diff.EvalDiff(
		map[string]any{"change": []any{
			map[string]any{
				"add":     "add",
				"changed": "2",
				"ITEM_ID": "remove_change",
				"nested": map[string]any{
					"add":     "add",
					"changed": "2",
				},
			},
		}},
		map[string]any{"FIO": []map[string]any{}},
	)
	require.Error(t, err)
}

func TestTooNestedObject(t *testing.T) {
	_, err := pdp_diff.EvalDiff(
		map[string]any{"array": map[string]any{
			"remove_1": "1",
			"remove_2": "2",
			"ITEM_ID":  "remove_array",
			"nested": map[string]any{
				"add": "add",
			},
		}},
		map[string]any{"FIO": []map[string]any{}},
	)
	require.Error(t, err)
}

func TestEvalDiff_AddItem(t *testing.T) {
	_, err := pdp_diff.EvalDiff(
		map[string]any{},
		map[string]any{"FIO": []map[string]any{
			{"ITEM_ID": "1", "a": "1", "b": "2", "c": "3"},
		}},
	)
	require.Error(t, err)
}

func TestEvalDiff_ChangeItem(t *testing.T) {
	_, err := pdp_diff.EvalDiff(
		map[string]any{"FIO": []map[string]any{
			{"ITEM_ID": "1", "a": "2", "b": "2", "c": "3"},
		}},
		map[string]any{"FIO": []map[string]any{
			{"ITEM_ID": "1", "a": "1", "b": "2", "c": "3"},
		}},
	)
	require.Error(t, err)
}
