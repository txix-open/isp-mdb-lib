package entity

import (
	"encoding/json"
	"time"
)

const (
	RecordsTableName     = "data_records"
	TechRecordsTableName = "data_tech_records"
)

type BaseRecord struct {
	Id         int64
	ExternalId string `valid:"required~Required"`
	Version    int64  `valid:"required~Required"`
	Type       string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

type DataRecord struct {
	tableName string `pg:"?db_schema.data_records" json:"-"`
	*BaseRecord
	Data       map[string]interface{} `valid:"required~Required"`
	CustomData map[string]interface{}
}

type TransitDataRecord struct {
	tableName string `pg:"?db_schema.data_records" json:"-"`
	*BaseRecord
	Data       json.RawMessage `valid:"required~Required"`
	CustomData json.RawMessage
}
