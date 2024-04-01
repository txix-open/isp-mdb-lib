//nolint:forcetypeassert
package mutator_test

import (
	"github.com/txix-open/isp-mdb-lib/delta"
)

func (t *TestServiceApply) Test_ChangeField_Object_NewBlock() {
	data := t.makeData([]byte(`{}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
				IsArray:   false,
				ItemId:    "",
			},
			Field:     "LastName",
			Operation: delta.ChangeField,
			NewValue:  "234",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("234", data["FIO"].(map[string]any)["LastName"])
}

func (t *TestServiceApply) Test_ChangeField_Object_NewItem() {
	data := t.makeData([]byte(`{"FIO":{}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
				IsArray:   false,
				ItemId:    "",
			},
			Field:     "LastName",
			Operation: delta.ChangeField,
			NewValue:  "234",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("234", data["FIO"].(map[string]any)["LastName"])
}

func (t *TestServiceApply) Test_ChangeField_Object_ExistItem() {
	data := t.makeData([]byte(`{"FIO":{"LastName": "123"}}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "FIO",
				IsArray:   false,
				ItemId:    "",
			},
			Field:     "LastName",
			Operation: delta.ChangeField,
			NewValue:  "234",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("234", data["FIO"].(map[string]any)["LastName"])
}

func (t *TestServiceApply) Test_ChangeField_Array_NewBlock() {
	data := t.makeData([]byte(`{}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Field:     "VIN",
			Operation: delta.ChangeField,
			NewValue:  "234",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("234", data["VEHICLES"].([]any)[0].(map[string]any)["VIN"])
}

func (t *TestServiceApply) Test_ChangeField_Array_NewItem() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Field:     "VIN",
			Operation: delta.ChangeField,
			NewValue:  "123456",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("123456", data["VEHICLES"].([]any)[0].(map[string]any)["VIN"])
}

func (t *TestServiceApply) Test_ChangeField_Array_ExistItem() {
	data := t.makeData([]byte(`{"VEHICLES":[{"ITEM_ID":"234","VIN":"123456"}]}`))
	changelogs := []delta.Changelog{
		{
			Path: delta.Path{
				BlockName: "VEHICLES",
				IsArray:   true,
				ItemId:    "234",
			},
			Field:     "VIN",
			Operation: delta.ChangeField,
			NewValue:  "234567",
		},
	}
	err := t.service.Apply(data, changelogs)
	t.Require().NoError(err)
	t.Equal("234567", data["VEHICLES"].([]any)[0].(map[string]any)["VIN"])
}
