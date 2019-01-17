package diff

import (
	"github.com/integration-system/go-cmp/cmp"
	"reflect"
)

type diffCollector struct {
	cmp.Option
	Delta
}

func (c *diffCollector) Report(x, y reflect.Value, eq bool, p cmp.Path) {
	if eq {
		return
	}

	newPath := make(cmp.Path, 0)
	for _, ps := range p {
		switch ps.(type) {
		case cmp.SliceIndex, cmp.StructField, cmp.MapIndex:
			newPath = append(newPath, ps)
		}
	}
	path := newPath.ToJSON()
	if path == "" {
		return
	}

	xIsValid := x.IsValid()
	yIsValid := y.IsValid()
	lastStep := newPath[len(newPath)-1]
	var desc *DiffDescriptor = nil
	switch ls := lastStep.(type) {
	case cmp.SliceIndex:
		i, j := ls.SplitKeys()
		if i == j {
			desc = &DiffDescriptor{
				Operation: ArrayChange,
				OldIndex:  &i,
				NewIndex:  &j,
			}
		} else if i == -1 {
			desc = &DiffDescriptor{
				Operation: ArrayAdd,
				NewIndex:  &j,
			}
		} else if j == -1 {
			desc = &DiffDescriptor{
				Operation: ArrayDelete,
				OldIndex:  &i,
			}
		} else {
			desc = &DiffDescriptor{
				Operation: ArraySwap,
				OldIndex:  &i,
				NewIndex:  &j,
			}
		}
	case cmp.StructField, cmp.MapIndex:
		if xIsValid && yIsValid {
			desc = &DiffDescriptor{
				Operation: Change,
			}
		} else if xIsValid && !yIsValid {
			desc = &DiffDescriptor{
				Operation: Delete,
			}
		} else if !xIsValid && yIsValid {
			desc = &DiffDescriptor{
				Operation: Add,
			}
		}
	}

	if desc != nil {
		var oldVal interface{} = nil
		var newVal interface{} = nil
		if xIsValid {
			oldVal = x.Interface()
		}
		if yIsValid {
			newVal = y.Interface()
		}
		desc.Path = path
		desc.OldValue = oldVal
		desc.NewValue = newVal

		l := len(c.Delta)
		if l > 0 {
			prev := c.Delta[l-1]
			if desc.Path == prev.Path && prev.Operation == ArrayDelete && desc.Operation == ArrayAdd {
				desc.Operation = ArrayChange
				desc.OldValue = prev.OldValue
				desc.OldIndex = prev.OldIndex
				c.Delta[l-1] = desc
			} else {
				c.Delta = append(c.Delta, desc)
			}
		} else {
			c.Delta = append(c.Delta, desc)
		}
	}
}
