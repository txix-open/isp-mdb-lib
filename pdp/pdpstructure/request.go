package pdpstructure

import (
	"github.com/integration-system/isp-mdb-lib/pdp"
	"github.com/integration-system/isp-mdb-lib/structure"
)

type PdpUpsertRequest struct {
	structure.SudirUpdateRecordRequest
	UpdateMethod pdp.HandleMethod
}
