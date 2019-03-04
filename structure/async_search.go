package structure

import "github.com/integration-system/isp-mdb-lib/query"

type Amqp struct {
	ExchangeName string `valid:"required~Required"`
	ExchangeKind string
	QueueName    string
	RoutingKey   string `valid:"required~Required"`
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
	ApplicationId int `valid:"required~Required"`
	JobCountLimit int
	RequestId     string `valid:"required~Required"`
	AsyncSearchRequest
}

type GetJobStatusRequest struct {
	RequestId string `valid:"required~Required"`
}

type GetAsyncResultRequest struct {
	RequestId string `valid:"required~Required"`
	Limit     int
	Offset    int
}
