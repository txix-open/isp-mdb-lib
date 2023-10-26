package entity

// nolint:unused
type DataTechRecord struct {
	tableName string `pg:"?db_schema.data_tech_records" json:"-"`

	*DataRecord
}

// nolint:unused
type TransitDataTechRecord struct {
	tableName string `pg:"?db_schema.data_tech_records" json:"-"`

	*TransitDataRecord
}
