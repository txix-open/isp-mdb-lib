package entity

import (
	"encoding/json"

	"github.com/integration-system/isp-mdb-lib/diff"
)

type BaseNotification struct {
	ExternalId  string
	OperationId string
	Type        string
	Version     int64
	IsTech      bool
	Delete      bool
	Delta       diff.Delta
}

type Notification struct {
	*BaseNotification
	Data       map[string]interface{}
	CustomData map[string]interface{}
}

type TransitNotification struct {
	*BaseNotification
	Data       json.RawMessage
	CustomData json.RawMessage
}
