//nolint:forcetypeassert
package mutator_test

import (
	"github.com/integration-system/isp-mdb-lib/delta"
)

func (t *TestServiceApply) Test_DeleteField_Object_EmptyBlock() {
	data := t.makeData([]byte(`{}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
				IsArray:   false,
				ItemId:    "",
			},
			Field:     "LastName",
			Operation: delta.DeleteField,
			NewValue:  "",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}

func (t *TestServiceApply) Test_DeleteField_Object_ItemNotExist() {
	data := t.makeData([]byte(`{"FIO":{}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
				IsArray:   false,
				ItemId:    "",
			},
			Field:     "LastName",
			Operation: delta.DeleteField,
			NewValue:  "",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}

func (t *TestServiceApply) Test_DeleteField_Object_ExistItem() {
	data := t.makeData([]byte(`{"FIO":{"LastName":"123","FirstName":"234"}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
				IsArray:   false,
				ItemId:    "",
			},
			Field:     "LastName",
			Operation: delta.DeleteField,
			NewValue:  "",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal(nil, data["FIO"].(map[string]any)["LastName"])
	t.Equal("234", data["FIO"].(map[string]any)["FirstName"])
}

func (t *TestServiceApply) Test_DeleteField_Array_NewBlock() {
	data := t.makeData([]byte(`{}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Field:     "LastName",
			Operation: delta.DeleteField,
			NewValue:  "",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}

func (t *TestServiceApply) Test_DeleteField_Array_EmptyItem() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Field:     "VIN",
			Operation: delta.DeleteField,
			NewValue:  "",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
}

func (t *TestServiceApply) Test_DeleteField_Array_ExistItem() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234","VIN":"123456","B":"5432"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Field:     "VIN",
			Operation: delta.DeleteField,
			NewValue:  "",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal(nil, data["VEHICLES"].([]any)[0].(map[string]any)["VIN"])
	t.Equal("5432", data["VEHICLES"].([]any)[0].(map[string]any)["B"])
}
