//nolint:forcetypeassert
package mutator_test

import (
	"github.com/txix-open/isp-mdb-lib/delta"
)

func (t *TestServiceApply) Test_DeleteArrayItem_HappyPath() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234","VIN":"123456"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.DeleteArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)

	result := data["VEHICLES"].([]any)
	t.Equal(0, len(result))
}

func (t *TestServiceApply) Test_DeleteArrayItem_MultiplyItems() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"123","VIN":"123456"},{"ITEM_ID":"234","VIN":"234567"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.DeleteArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)

	result := data["VEHICLES"].([]any)
	t.Equal(1, len(result))
	found := false
	for _, m := range result {
		if m.(map[string]any)["ITEM_ID"] == "123" {
			found = true
		}
	}
	t.Equal(true, found)
}

func (t *TestServiceApply) Test_DeleteArrayItem_EmptyArray() {
	data := t.makeData([]byte(`{"VEHICLES":[]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.DeleteArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}

func (t *TestServiceApply) Test_DeleteArrayItem_NotFoundBlock() {
	data := t.makeData([]byte(`{"FIO":{}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.DeleteArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}

func (t *TestServiceApply) Test_DeleteArrayItem_NotFoundItemId() {
	data := t.makeData([]byte(`{"VEHICLES":[{"VIN":"123456"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.DeleteArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().Error(err)
}

func (t *TestServiceApply) Test_DeleteArrayItem_UnknownItemId() {
	data := t.makeData([]byte(`{"VEHICLES":[{"VIN":"123456", "ITEM_ID": "123"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.DeleteArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}
