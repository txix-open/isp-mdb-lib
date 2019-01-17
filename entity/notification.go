package entity

import (
	"github.com/integration-system/isp-mdb-lib/diff"
)

type Notification struct {
	ExternalId  string
	OperationId string
	Version     int64
	IsTech      bool
	Data        map[string]interface{}
	CustomData  map[string]interface{}
	Delta       diff.Delta
}
