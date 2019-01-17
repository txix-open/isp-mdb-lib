package entity

import "time"

type Attribute struct {
	TableName string `sql:"?db_schema.attributes" json:"-"`
	Id        int32
	Path      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
