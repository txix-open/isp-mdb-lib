package structure

import "github.com/integration-system/isp-mdb-lib/query"

type AsyncSearchRequest struct {
	Query       query.Term
	PackageSize int
	Callback    string
	Protocol    ProtocolVersion
}
