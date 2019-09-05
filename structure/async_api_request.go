package structure

import "github.com/integration-system/isp-mdb-lib/query"

type Amqp struct {
	ExchangeName string `valid:"required~Required" schema:"Название точки маршрутизации"`
	ExchangeKind string `schema:"Тип точки маршрутизации,'direct', 'topic', 'fanout' (по умолчанию - 'direct')"`
	QueueName    string `schema:"Название очереди"`
	RoutingKey   string `valid:"required~Required" schema:"Ключ маршрутизации,для публикации напрямую в очередь, указывается название очереди"`
	Declare      bool   `schema:"Автоматическое объявление очереди,точки маршрутизации,привязки"`
}

type AsyncSearchRequest struct {
	PackageSize int
	TechEntries bool
	Callback    string
	Protocol    ProtocolVersion
	Amqp        *Amqp
}

type ExternalAsyncSearchRequest struct {
	Query map[string]interface{} `valid:"required~Required"`
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
