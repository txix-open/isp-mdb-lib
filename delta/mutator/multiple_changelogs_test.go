package mutator_test

import (
	"encoding/json"

	"github.com/txix-open/isp-mdb-lib/delta"
)

//nolint:funlen
func (t *TestServiceApply) Test_MultipleChangelogs_HappyPath() {
	data := t.makeData([]byte(`
{
  "SNILS": {"Snils": "-------"},
  "RELATIVES": [{"ITEM_ID": "112"}],
  "PETS": [{"ITEM_ID": "228"},{"ITEM_ID": "911"}]
}
`))
	changelogs := make([]delta.Changelog, 0)
	err := json.Unmarshal([]byte(`
[
	{
	  "path": {"blockName": "FIO","isArray": false,"itemId": ""},
	  "field": "",
	  "operation": "DELETE_BLOCK",
	  "newValue": ""
	},
    {
      "path": {"blockName": "USER","isArray": false,"itemId": ""},
      "field": "LoginName",
      "operation": "CHANGE_FIELD",
      "newValue": "ID"
    },
    {
      "path": {"blockName": "FIO","isArray": false,"itemId": ""},
      "field": "LastName",
      "operation": "CHANGE_FIELD",
      "newValue": "LASTNAME"
    },
    {
      "path": {"blockName": "PERSON","isArray": false,"itemId": ""},
      "field": "Birthday",
      "operation": "CHANGE_FIELD",
      "newValue": "10.12.1995"
    },
    {
      "path": {"blockName": "SNILS","isArray": false,"itemId": ""},
      "field": "Snils",
      "operation": "DELETE_FIELD",
      "newValue": null
    },
    {
      "path": {"blockName": "VEHICLE","isArray": true,"itemId": "F08A648641AE40959A83C22FF9829855"},
      "field": "",
      "operation": "ADD_ARRAY_ITEM",
      "newValue": null
    },
    {
      "path": {"blockName": "VEHICLE","isArray": true,"itemId": "F08A648641AE40959A83C22FF9829855"},
      "field": "VehicleDescription",
      "operation": "CHANGE_FIELD",
      "newValue": "FORD"
    },
    {
      "path": {"blockName": "RELATIVES","isArray": true,"itemId": "112"},
      "field": "RltvBirthDate",
      "operation": "CHANGE_FIELD",
      "newValue": "11.09.2021"
    },
    {
      "path": {"blockName": "RELATIVES","isArray": true,"itemId": "112"},
      "field": "RltvSnils",
      "operation": "CHANGE_FIELD",
      "newValue": "778"
    },
    {
      "path": {"blockName": "PETS","isArray": true,"itemId": "228"},
      "field": "",
      "operation": "DELETE_ARRAY_ITEM",
      "newValue": null
    },
    {
      "path": {"blockName": "PETS","isArray": true,"itemId": "911"},
      "field": "",
      "operation": "DELETE_ARRAY_ITEM",
      "newValue": null
    }
  ]
`), &changelogs)
	t.Require().NoError(err)

	err = t.service.Apply(data, changelogs)
	t.Require().NoError(err)

	_, err = json.Marshal(data)
	t.Require().NoError(err)
}
