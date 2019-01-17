package diff

import (
	"github.com/integration-system/go-cmp/cmp"
)

type Operation string

const (
	Add    Operation = "ADD"
	Delete Operation = "DELETE"
	Change Operation = "CHANGE"

	ArrayAdd    Operation = "ARRAY_ADD"
	ArrayDelete Operation = "ARRAY_DELETE"
	ArrayChange Operation = "ARRAY_CHANGE"
	ArraySwap   Operation = "ARRAY_SWAP"
)

type DiffDescriptor struct {
	OldValue  interface{} `json:"oldValue,omitempty"`
	NewValue  interface{} `json:"newValue,omitempty"`
	Path      string      `json:"path"`
	Operation Operation   `json:"operation"`
	OldIndex  *int        `json:"oldIndex,omitempty"`
	NewIndex  *int        `json:"newIndex,omitempty"`
}

type Delta []*DiffDescriptor

func EvalDiff(left, right map[string]interface{}) (bool, Delta) {
	c := new(diffCollector)
	c.Delta = make(Delta, 0)
	return cmp.Equal(left, right, c), c.Delta
}
