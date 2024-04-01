package structure

import "github.com/txix-open/isp-mdb-lib/entity"

type UpsertResponse struct {
	OldValue *entity.DataRecord
	NewValue *entity.DataRecord
	Inserted bool
	Updated  bool
}
