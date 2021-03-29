package diff

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalDiff(t *testing.T) {
	oldValue := []byte(`
{"cards":{},"users":{"SSO":[{"id_type":"SSO","last_update_dt":"2020-11-09 20:13:18.000","is_confirmed_offline":false}]},"mdm_id":"123456","contacts":{"7":[{"ref_num":"1","etalon_id":"1","mdm_obj_id":"1","last_update_dt":"2020-11-09 20:13:18.000","sourcechannels":[4],"cont_meth_cat_cd":"7"}],"8":[{"ref_num":"1","etalon_id":"1","mdm_obj_id":"1","last_update_dt":"2020-11-09 20:13:18.000","sourcechannels":[4],"cont_meth_cat_cd":"8"}]},"addresses":{"1":[{"unad":"1","unom":"1","house_no":"1","city_name":"","corpus_no":"2","etalon_id":"2","street_id":"2","confirming":false,"mdm_obj_id":"2","sourcechannels":[4]}],"8":[{"workplace":{"sourcechannels":[4],"mdm_workplaces_address_rel":"C599A50C1D674960B0D00FE214A3C225"},"confirming":false,"mdm_obj_id":"C599A50C1D674960B0D00FE214A3C225","street_omk":"22545","validation":false,"hpsm_number":"","stroenie_no":"","chronicle_id":"C599A50C1D674960B0D00FE214A3C225","hpsm_comment":"","temporary_id":"C599A50C1D674960B0D00FE214A3C225","addr_line_one":"Пятницкое шоссе, д.дом 23, стр., корп., кв.","addr_line_two":"Северо-Западный административный округ, район муниципальный округ Митино","building_name":"31984.1","residence_num":"","last_update_dt":"2020-04-17 15:22:50.000","sourcechannels":[4]}]},"mdm_version":1616771016040586500,"escredentials":{},"citizen_relatives":null}
`)
	newValue := []byte(`
{"cards":{"9":[{"del_sign":false,"card_type":"9","etalon_id":"1","mdm_obj_id":"1","primary_id":"1","card_number":"0490000657517","card_species":"03","last_update_dt":"2021-02-11 15:41:13.000","sourcechannels":[17]}]},"users":{"SSO":[{"id_type":"SSO","last_update_dt":"2021-02-11 15:41:13.000","is_confirmed_offline":false}],"ABIS":[{"id_type":"ABIS","id_value":"1","etalon_id":"1","mdm_obj_id":"1","primary_id":"1","chronicle_id":"1","last_update_dt":"2021-02-11 15:41:13.000","is_confirmed_offline":false}]},"mdm_id":"123456","contacts":{"7":[{"ref_num":"1","etalon_id":"1","mdm_obj_id":"1","last_update_dt":"2020-11-09 20:13:18.000","sourcechannels":[4],"cont_meth_cat_cd":"7"}],"8":[{"ref_num":"1","etalon_id":"1","mdm_obj_id":"1","last_update_dt":"2020-11-09 20:13:18.000","sourcechannels":[4],"cont_meth_cat_cd":"8"}]},"addresses":{"1":[{"unad":"1","unom":"1","house_no":"1","city_name":"","corpus_no":"2","etalon_id":"2","street_id":"2","confirming":false,"mdm_obj_id":"2","primary_id":"1","sourcechannels":[4]}],"8":[{"workplace":{"primary_id":"1","sourcechannels":[4],"mdm_workplaces_address_rel":"C599A50C1D674960B0D00FE214A3C225"},"primary_id":"1","sourcechannels":[4]}]},"mdm_version":1617008663078477300,"escredentials":{},"citizen_relatives":null}
`)
	a := assert.New(t)
	var left, right map[string]interface{}
	err := json.Unmarshal(oldValue, &left)
	a.NoError(err)
	err = json.Unmarshal(newValue, &right)
	a.NoError(err)

	expectedPath := map[string]bool{
		"addresses.1.[0].primary_id":   true,
		"addresses.8.[0]":              true,
		"cards.9":                      true,
		"mdm_version":                  true,
		"users.ABIS":                   true,
		"users.SSO.[0].last_update_dt": true,
	}
	_, delta := EvalDiff(left, right)
	for _, descriptor := range delta {
		a.True(expectedPath[descriptor.Path])
	}

	expectedExtensionPath := map[string]bool{
		"addresses.1.[0].primary_id":          true,
		"addresses.8.[0]":                     true,
		"cards.9.[0].etalon_id":               true,
		"cards.9.[0].mdm_obj_id":              true,
		"cards.9.[0].last_update_dt":          true,
		"cards.9.[0].card_type":               true,
		"cards.9.[0].sourcechannels.[0]":      true,
		"cards.9.[0].del_sign":                true,
		"cards.9.[0].card_species":            true,
		"cards.9.[0].primary_id":              true,
		"cards.9.[0].card_number":             true,
		"mdm_version":                         true,
		"users.ABIS.[0].mdm_obj_id":           true,
		"users.ABIS.[0].primary_id":           true,
		"users.ABIS.[0].chronicle_id":         true,
		"users.ABIS.[0].last_update_dt":       true,
		"users.ABIS.[0].is_confirmed_offline": true,
		"users.ABIS.[0].id_type":              true,
		"users.ABIS.[0].id_value":             true,
		"users.ABIS.[0].etalon_id":            true,
		"users.SSO.[0].last_update_dt":        true,
	}
	delta = ExtensionDelta(delta)
	for _, descriptor := range delta {
		a.True(expectedExtensionPath[descriptor.Path])
	}
}

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
						map[string]interface{}{"1": 1, "2": 2, "3": 3},
						map[string]interface{}{"4": 4, "5": 5},
					},
				},
				{
					Path:      "pets",
					Operation: Delete,
					OldValue:  "delete_pets",
				},
				{
					Path:      "pets.del_sign",
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
				"addr": {
					Path: "addr", Operation: Add, NewValue: []interface{}{
						map[string]interface{}{"1": 1, "2": 2, "3": 3},
						map[string]interface{}{"4": 4, "5": 5},
					},
				},
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
				"pets.del_sign": {
					Path: "pets.del_sign", Operation: Delete, OldValue: "delete_pets",
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

func TestReplaceArray(t *testing.T) {
	a := assert.New(t)
	type example struct {
		d       Delta
		newPath string
	}
	for _, e := range []example{
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "documents.[1].primary_id",
			}, {
				NewValue: "1", Operation: Add,
				Path: "documents.[2].primary_id",
			}},
			newPath: "documents.primary_id",
		},
		{
			d: Delta{{
				NewValue: "1", Operation: Add,
				Path: "$$cards.9.[0].mdm_obj_id",
			}},
			newPath: "$$cards.9.mdm_obj_id",
		},
	} {
		actual := ReplaceArray(e.d)
		for _, descriptor := range actual {
			a.Equal(e.newPath, descriptor.Path)
		}
	}
}
