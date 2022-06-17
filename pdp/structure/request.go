package structure

import (
	"github.com/integration-system/isp-mdb-lib/pdp"
	"github.com/integration-system/isp-mdb-lib/pdp/entity"
	"github.com/integration-system/isp-mdb-lib/structure"
)

type PdpUpsertRequest struct {
	SudirUpdateRecordRequest
	UpdateMethod pdp.HandleMethod
}

type PdpUpsertBatchRequest struct {
	Requests     []SudirUpdateRecordRequest
	UpdateMethod pdp.HandleMethod
	AppSource    int32
}

type SudirUpdateRecordRequest struct {
	TechRecord       bool
	Record           *entity.DataRecord `valid:"required~Required"`
	SoftDelete       bool
	DeleteOperations map[string]map[string]string
	Error            *structure.ConvertSearchError
}

type BatchConvertPdpRequest struct {
	*structure.ConvertRequestPayload `valid:"required~Required"`
	*structure.AbstractConvertBatchRequest
	Delta PdpDelta
}

/*
	DeletedObject e.q.
	{
		Id: "62D1FA909140BABC0C10DA90A13B041E30",
		ObjectGroup: "VEHICLE",
	}
*/
type DeletedObject struct {
	Id          string
	ObjectGroup string
}

/*
	DeltaInfo e.q.
	{
		AttributeDeltaData: {
			"FIO": {
				"MIDNAME": "IVANOVICH",
				"SURNAME": "",
			}
		}
		ObjectDeltaData: {
			"VEHICLES": [
				{
					"ITEM_ID": "1234567890_1",
					"VEHICLE_DESC": "TEST VEH 1"
				},
				{
					"ITEM_ID": "1234567890_2",
					"VEHICLE_DESC": ""
				}
			]
		}
	}
*/
type DeltaInfo struct {
	AttributeDeltaData map[string]map[string]interface{}
	ObjectDeltaData    map[string][]map[string]interface{}
}

type PdpDelta struct {
	DeltaInfo      DeltaInfo
	DeletedObjects []DeletedObject
}

type PdpNotification struct {
	entity.Notification

	PdpDelta PdpDelta
}

type PdpNotificationMessage struct {
	Notifications []PdpNotification
	AppSource     int32
}
