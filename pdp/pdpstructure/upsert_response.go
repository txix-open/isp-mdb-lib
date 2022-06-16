package pdpstructure

import (
	"github.com/integration-system/isp-mdb-lib/pdp/pdpentity"
)

type UpsertResponse struct {
	OldValue *pdpentity.DataRecord
	NewValue *pdpentity.DataRecord
	Inserted bool
	Updated  bool
}
