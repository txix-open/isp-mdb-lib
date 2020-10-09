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
						map[string]interface{}{
							"1": 1,
							"2": 2,
							"3": 3,
						},
						map[string]interface{}{
							"4": 4,
							"5": 5,
						},
					},
				},
				{
					Path:      "pets",
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
