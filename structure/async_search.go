package structure

import "github.com/integration-system/isp-mdb-lib/query"

type AsyncSearchRequest struct {
	PackageSize int
	Callback    string
	Protocol    ProtocolVersion
}

type ExternalAsyncSearchRequest struct {
	Query map[string]interface{}
	AsyncSearchRequest
}

type InternalAsyncSearchRequest struct {
	Query         query.Term
	ApplicationId int
	AsyncSearchRequest
}
