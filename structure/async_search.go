package structure

import "github.com/integration-system/isp-mdb-lib/query"

type Amqp struct {
	ExchangeName string `valid:"required~Required" schema:"Exchange name"`
	ExchangeKind string `schema:"Exchange kind,'direct','topic','fanout'. Default:'direct'"`
	QueueName    string `schema:"Queue name"`
	RoutingKey   string `valid:"required~Required" schema:"Routing key"`
	Declare      bool   `schema:"Auto declaration,If enable it automatically declares queue, exchange and binding"`
}

type AsyncSearchRequest struct {
	PackageSize int `valid:"required~Required"`
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
	ApplicationId int32 `valid:"required~Required"`
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
