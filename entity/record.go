package entity

import (
	"time"
)

type DataRecord struct {
	TableName string `sql:"?db_schema.data_records" json:"-"`

	Id         int64                  `json:"id"`
	ExternalId string                 `valid:"required~Required" json:"externalId"`
	Data       map[string]interface{} `valid:"required~Required" json:"data"`
	CustomData map[string]interface{} `json:"customData"`
	Version    int64                  `valid:"required~Required" json:"version"`
	UpdatedAt  time.Time              `json:"updatedAt"`
	CreatedAt  time.Time              `json:"createdAt"`
}
