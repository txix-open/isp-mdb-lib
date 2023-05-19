//nolint:forcetypeassert
package mutator_test

import (
	"github.com/integration-system/isp-mdb-lib/delta"
)

func (t *TestServiceApply) Test_AddArrayItem_NewObject() {
	data := t.makeData([]byte(`{}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "123",
			},
			Operation: delta.AddArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("123", data["VEHICLES"].([]any)[0].(map[string]any)["ITEM_ID"])
}

func (t *TestServiceApply) Test_AddArrayItem_EmptyObject() {
	data := t.makeData([]byte(`{"VEHICLES":[]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "123",
			},
			Operation: delta.AddArrayItem,
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

func (t *TestServiceApply) Test_AddArrayItem_NewItem() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234","VIN":"123456"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "123",
			},
			Operation: delta.AddArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)

	result := data["VEHICLES"].([]any)
	t.Equal(2, len(result))
	found := false
	for _, m := range result {
		if m.(map[string]any)["ITEM_ID"] == "123" {
			found = true
		}
	}
	t.Equal(true, found)
}

func (t *TestServiceApply) Test_AddArrayItem_NotArrayData() {
	data := t.makeData([]byte(`{"VEHICLES":{"ITEM_ID":"234","VIN":"123456"}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   false,
				ItemId:    "123",
			},
			Operation: delta.AddArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().Error(err)
}

func (t *TestServiceApply) Test_AddArrayItem_PathIsArrayFalse() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234","VIN":"123456"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   false,
				ItemId:    "123",
			},
			Operation: delta.AddArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().Error(err)
}

func (t *TestServiceApply) Test_AddArrayItem_AlreadyExistItem() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234","VIN":"123456"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Operation: delta.AddArrayItem,
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().Error(err)
}
