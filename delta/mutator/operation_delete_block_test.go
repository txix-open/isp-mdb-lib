//nolint:forcetypeassert
package mutator_test

import (
	"github.com/integration-system/isp-mdb-lib/delta"
)

func (t *TestServiceApply) Test_DeleteBlock_Exists() {
	data := t.makeData([]byte(`{"FIO":{"LastName":"123","FirstName":"234"}, "REL": {"PROP":"VAL"}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
			},
			Operation: delta.DeleteBlock,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal(nil, data["FIO"])
	t.Equal("VAL", data["REL"].(map[string]any)["PROP"])
}

func (t *TestServiceApply) Test_DeleteBlock_NotExists() {
	data := t.makeData([]byte(`{"REL": {"PROP":"VAL"}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
			},
			Operation: delta.DeleteBlock,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal(nil, data["FIO"])
	t.Equal("VAL", data["REL"].(map[string]any)["PROP"])
}
