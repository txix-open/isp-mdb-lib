package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
				Path: "documents[1].primary_id",
			}, {
				NewValue: "1", Operation: Add,
				Path: "documents[2].primary_id",
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
	} {
		actual := ReplaceArray(e.d)
		for _, descriptor := range actual {
			a.Equal(e.newPath, descriptor.Path)
		}
	}
}
