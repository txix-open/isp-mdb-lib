package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalDiff(t *testing.T) {
	left := map[string]interface{}{
		"remove_value":   "1",
		"changed_value":  1617008663078477300,
		"save_old_value": "3",
		"array": map[string]interface{}{
			"remove": []interface{}{
				0: map[string]interface{}{
					"remove_1": "1",
					"remove_2": "2",
				},
			},
			"change": []interface{}{
				0: map[string]interface{}{
					"remove":  "remove",
					"changed": "1",
					"save": map[string]interface{}{
						"i": []int{
							4,
						},
						"remove": "remove",
					},
				},
			},
			// TODO get `array.change2.changed path for this indexes
			// when must be `array.change2.[1].changed` or ADD/DELETE values
			// bad works only for 1 field in array and 1 -> 0 indexes
			//
			//"change2": []interface{}{
			//	0: map[string]interface{}{
			//		"remove": "remove",
			//	},
			//	1: map[string]interface{}{
			//		"changed":  1,
			//	},
			//	2: map[string]interface{}{
			//		"save": "save",
			//	},
			//},
			"change3": []interface{}{
				0: map[string]interface{}{
					"changed": "1",
					"a":       1,
				},
				1: map[string]interface{}{
					"remove": "remove",
				},
				2: map[string]interface{}{
					"save": "save",
				},
			},
			"saved": []interface{}{
				0: map[string]interface{}{
					"1": "1",
					"2": "2",
				},
			},
		},
		"empty":     struct{}{},
		"undefined": nil,
	}

	right := map[string]interface{}{
		"add_value":      "2",
		"changed_value":  1616771016040586500,
		"save_old_value": "3",
		"array": map[string]interface{}{
			"add": []interface{}{
				0: map[string]interface{}{
					"add_1": "1",
					"add_2": "2",
				},
			},
			"change": []interface{}{
				0: map[string]interface{}{
					"save": map[string]interface{}{
						"i": []int{
							4,
						},
						"add": "add",
					},
					"add":     "add",
					"changed": "2",
				},
			},
			//"change2": []interface{}{
			//	0: map[string]interface{}{
			//		"changed": 2,
			//	},
			//	1: map[string]interface{}{
			//		"save": "save",
			//	},
			//	2: map[string]interface{}{
			//		"add":  "add",
			//	},
			//},
			"change3": []interface{}{
				0: map[string]interface{}{
					"changed": "2",
					"a":       1,
				},
				1: map[string]interface{}{
					"save": "save",
				},
				2: map[string]interface{}{
					"add": "add",
				},
			},
			"saved": []interface{}{
				0: map[string]interface{}{
					"1": "1",
					"2": "2",
				},
			},
		},
		"empty":     struct{}{},
		"undefined": nil,
	}

	a := assert.New(t)
	expectedEvalDiffPath := map[string]bool{
		"array.remove":     true,
		"array.add":        true,
		"array.change.[0]": true,
		//"array.change2.[0]":         true,
		//"array.change2.[1].changed": true,
		//"array.change2.[2]":         true,
		"array.change3.[0].changed": true,
		"array.change3.[1]":         true,
		"array.change3.[2]":         true,
		"changed_value":             true,
		"remove_value":              true,
		"add_value":                 true,
	}
	_, delta := EvalDiff(left, right)
	for _, descriptor := range delta {
		a.True(expectedEvalDiffPath[descriptor.Path])
	}

	expectedExtension := map[string]*DiffDescriptor{
		"add_value": {
			Operation: Add, Path: "add_value", OldValue: nil, NewValue: "2",
		},
		"changed_value": {
			Operation: Change, Path: "changed_value", OldValue: 1617008663078477300, NewValue: 1616771016040586500,
		},
		"remove_value": {
			Operation: Delete, Path: "remove_value", OldValue: "1", NewValue: nil,
		},
		"array.add.[0].add_1": {
			Operation: Add, Path: "array.add.[0].add_1", OldValue: nil, NewValue: "1",
		},
		"array.add.[0].add_2": {
			Operation: Add, Path: "array.add.[0].add_2", OldValue: nil, NewValue: "2",
		},
		"array.remove.[0].remove_1": {
			Operation: Delete, Path: "array.remove.[0].remove_1", OldValue: "1", NewValue: nil,
		},
		"array.remove.[0].remove_2": {
			Operation: Delete, Path: "array.remove.[0].remove_2", OldValue: "2", NewValue: nil,
		},
		"array.change.[0].add": {
			Operation: Add, Path: "array.change.[0].add", OldValue: nil, NewValue: "add",
		},
		"array.change.[0].changed": {
			Operation: Change, Path: "array.change.[0].changed", OldValue: "1", NewValue: "2",
		},
		"array.change.[0].save.add": {
			Operation: Add, Path: "array.change.[0].save.add", OldValue: nil, NewValue: "add",
		},
		"array.change.[0].save.remove": {
			Operation: Delete, Path: "array.change.[0].save.remove", OldValue: "remove", NewValue: nil,
		},
		"array.change.[0].remove": {
			Operation: Delete, Path: "array.change.[0].remove", OldValue: "remove", NewValue: nil,
		},
		//"array.change2.[0].remove": {
		//	Operation: Delete, Path: "array.change2.[0].remove", OldValue: "remove", NewValue: nil,
		//},
		//"array.change2.[1].changed": {
		//	Operation: Change, Path: "array.change2.[1].changed", OldValue: "1", NewValue: "2",
		//},
		//"array.change2.[2].add": {
		//	Operation: Add, Path: "array.change2.[2].add", OldValue: nil, NewValue: "add",
		//},
		"array.change3.[0].changed": {
			Operation: Change, Path: "array.change3.[0].changed", OldValue: "1", NewValue: "2",
		},
		"array.change3.[1].remove": {
			Operation: Delete, Path: "array.change3.[1].remove", OldValue: "remove", NewValue: nil,
		},
		"array.change3.[2].add": {
			Operation: Add, Path: "array.change3.[2].add", OldValue: nil, NewValue: "add",
		},
	}
	delta = ExtensionDelta(delta)
	for _, descriptor := range delta {
		a.Equal(expectedExtension[descriptor.Path], descriptor)
		delete(expectedExtension, descriptor.Path)
	}
	a.Len(expectedExtension, 0)
}

func TestExtensionDelta(t *testing.T) {
	a := assert.New(t)
	type example struct {
		request  Delta
		expected map[string]*DiffDescriptor
	}

	for _, e := range []example{
		{
			request: Delta{
				{
					Path: "add_empty_value", Operation: Add,
				},
			},
			expected: map[string]*DiffDescriptor{
				"add_empty_value": {Path: "add_empty_value", Operation: Add},
			},
		},
		{
			request: Delta{
				{
					Path:      "addr",
					Operation: Add,
					NewValue: []interface{}{
						map[string]interface{}{"1": 1, "2": 2, "3": 3},
						map[string]interface{}{"4": 4, "5": 5},
					},
				},
				{
					Path:      "pets",
					Operation: Delete,
					OldValue:  "delete_pets",
				},
				{
					Path:      "pets.del_sign",
					Operation: Delete,
					OldValue:  "delete_pets",
				},
				{
					Path:      "delete_empty_array",
					Operation: Delete,
					OldValue:  []interface{}{},
				},
			},
			expected: map[string]*DiffDescriptor{
				"addr": {
					Path: "addr", Operation: Add, NewValue: []interface{}{
						map[string]interface{}{"1": 1, "2": 2, "3": 3},
						map[string]interface{}{"4": 4, "5": 5},
					},
				},
				"addr.[0].1": {
					Path: "addr.[0].1", Operation: Add, NewValue: 1,
				},
				"addr.[0].2": {
					Path: "addr.[0].2", Operation: Add, NewValue: 2,
				},
				"addr.[0].3": {
					Path: "addr.[0].3", Operation: Add, NewValue: 3,
				},
				"addr.[1].4": {
					Path: "addr.[1].4", Operation: Add, NewValue: 4,
				},
				"addr.[1].5": {
					Path: "addr.[1].5", Operation: Add, NewValue: 5,
				},
				"pets": {
					Path: "pets", Operation: Delete, OldValue: "delete_pets",
				},
				"pets.del_sign": {
					Path: "pets.del_sign", Operation: Delete, OldValue: "delete_pets",
				},
				"delete_empty_array": {
					Path: "delete_empty_array", Operation: Delete, OldValue: []interface{}{},
				},
			},
		},
	} {
		response := ExtensionDelta(e.request)
		for _, diffDescriptor := range response {
			a.Equal(e.expected[diffDescriptor.Path], diffDescriptor)
		}
	}
}

func TestReplaceArray(t *testing.T) {
	a := assert.New(t)
	type example struct {
		d       Delta
		newPath string
	}
	for _, e := range []example{
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "documents.[1].primary_id",
			}, {
				NewValue: "1", Operation: Add,
				Path: "documents.[2].primary_id",
			}},
			newPath: "documents.primary_id",
		},
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "$$cards.9.[0].mdm_obj_id",
			}},
			newPath: "$$cards.9.mdm_obj_id",
		},
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "$$cards.9.[B7579F7C21554937ADE6E2721FF304DF].mdm_obj_id",
			}},
			newPath: "$$cards.9.mdm_obj_id",
		},
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "$$cards.9.[5E13DEE3-8953-4B83-8F3A-A238D19DF108].mdm_obj_id",
			}},
			newPath: "$$cards.9.mdm_obj_id",
		},
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "$$cards.9.[аЯвФ-ia51-r1--двЁ].mdm_obj_id",
			}},
			newPath: "$$cards.9.mdm_obj_id",
		},
	} {
		actual := ReplaceArray(e.d)
		for _, descriptor := range actual {
			a.Equal(e.newPath, descriptor.Path)
		}
	}
}
