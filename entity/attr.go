package entity

import "time"

type Attribute struct {
	TableName  string `pg:"?db_schema.attributes" json:"-"`
	Id         int32
	Path       string
	Type       string
	RecordType string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
