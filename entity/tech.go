package entity

type DataTechRecord struct {
	tableName string `pg:"?db_schema.data_tech_records" json:"-"`

	*DataRecord
}

type TransitDataTechRecord struct {
	tableName string `pg:"?db_schema.data_tech_records" json:"-"`

	*TransitDataRecord
}
