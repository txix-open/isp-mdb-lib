package diff

import (
	"reflect"
	"strings"

	"github.com/txix-open/go-cmp/cmp"
)

var arrayReplacer = strings.NewReplacer(`[`, `.[`)

type AdditionalDataMaker func(op Operation, x, y reflect.Value, path string, step cmp.PathStep) interface{}

type Option func(dc *diffCollector)

type diffCollector struct {
	adm AdditionalDataMaker
	cmp.Option
	Delta
}

func (dc *diffCollector) Report(x, y reflect.Value, eq bool, p cmp.Path) {
	if eq {
		return
	}

	newPath := make(cmp.Path, 0, len(p))
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
	path = arrayReplacer.Replace(path)

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

		if dc.adm != nil {
			desc.AdditionalData = dc.adm(desc.Operation, x, y, path, lastStep)
		}

		l := len(dc.Delta)
		if l > 0 {
			prev := dc.Delta[l-1]
			if desc.Path == prev.Path && prev.Operation == ArrayDelete && desc.Operation == ArrayAdd {
				desc.Operation = ArrayChange
				desc.OldValue = prev.OldValue
				desc.OldIndex = prev.OldIndex
				dc.Delta[l-1] = desc
			} else {
				dc.Delta = append(dc.Delta, desc)
			}
		} else {
			dc.Delta = append(dc.Delta, desc)
		}
	}
}

func NewDiffCollector(opts ...Option) *diffCollector {
	c := new(diffCollector)
	c.Delta = make(Delta, 0)

	for _, v := range opts {
		v(c)
	}

	return c
}

func MakeAdditionalData(f AdditionalDataMaker) Option {
	return func(dc *diffCollector) {
		dc.adm = f
	}
}
