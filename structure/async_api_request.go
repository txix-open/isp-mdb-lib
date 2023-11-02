package structure

import "github.com/integration-system/isp-mdb-lib/query"

type Amqp struct {
	ExchangeName string `valid:"required~Required" schema:"Название точки маршрутизации"`
	ExchangeKind string `schema:"Тип точки маршрутизации,'direct', 'topic', 'fanout' (по умолчанию - 'direct')"`
	QueueName    string `schema:"Название очереди"`
	RoutingKey   string `valid:"required~Required" schema:"Ключ маршрутизации,для публикации напрямую в очередь, указывается название очереди"`
	// автоматическое объявление очереди, точки маршрутизации, привязки
	Declare bool `schema:"Автоматическое объявление очереди,точки маршрутизации,привязки"`
}

type AsyncSearchRequest struct {
	// размер пакета с результатами поиска
	PackageSize int
	// поиск по техническим записям
	TechEntries bool
	// если поле не заданно используется внутренее хранилище результатов поиска
	// Поддерживается rabbit (начинается с amqp://) и http (c http:// или https://)
	// amqp://user:pass@host:10000/vhost
	Callback string
	Protocol ProtocolVersion
	// обязательный параметр, если указан callback RabbitMQ
	Amqp *Amqp
}

type ExternalAsyncSearchRequest struct {
	//название атрибута -> значение, все элементы объекта объединяются логическим 'И'
	Query map[string]interface{} `valid:"required~Required"`
	AsyncSearchRequest
}

type ExternalExtendedAsyncSearchRequest struct {
	Query *OneOfCondition `valid:"required~Required"`
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
