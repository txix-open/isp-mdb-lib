package structure

import "github.com/integration-system/isp-mdb-lib/query"

type Amqp struct {
	ExchangeName string
	ExchangeKind string
	QueueName    string
	RoutingKey   string
	Declare      bool
}

type AsyncSearchRequest struct {
	PackageSize int
	TechEntries bool
	Callback    string
	Protocol    ProtocolVersion
	Amqp        *Amqp
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
