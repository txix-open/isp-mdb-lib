package structure

import "github.com/txix-open/isp-mdb-lib/query"

type IdListApiRequest struct {
	//массив идентификаторов профилей
	Ids []string `valid:"required~Required"`
	//название протокола для конвертации найденных профилей, если не указан данные возвращаются в исходной структуре
	Protocol ProtocolVersion
	//запрос на поиск в технических записях
	TechEntries bool
}

type SyncSearchApiRequest struct {
	//название атрибута -> значение, все элементы объекта объединяются логическим 'И'
	Query map[string]interface{} `valid:"required~Required"`
	//название протокола для конвертации найденных профилей, если не указан данные возвращаются в исходной структуре
	Protocol ProtocolVersion
	//запрос на поиск в технических записях
	TechEntries bool
}

func (req *SyncSearchApiRequest) ToExtended() SyncSearchExtendedApiRequest {
	reqExt := SyncSearchExtendedApiRequest{
		Protocol:       "",
		TechEntries:    req.TechEntries,
		Query:          ToExtendedCondition(req.Query),
		resultProtocol: req.Protocol,
	}
	return reqExt
}

type SyncSearchExtendedApiRequest struct {
	//условия поиска
	Query *OneOfCondition `valid:"required~Required"`
	//название протокола для конвертации найденных профилей, если не указан данные возвращаются в исходной структуре
	Protocol ProtocolVersion
	//запрос на поиск в технических записях
	TechEntries bool

	resultProtocol ProtocolVersion //TODO хак, что бы сохранить обратную совместимость
}

func (r *SyncSearchExtendedApiRequest) GetResultProtocol() ProtocolVersion {
	if r.resultProtocol != "" {
		return r.resultProtocol
	}
	return r.Protocol
}

func ToExtendedCondition(m map[string]interface{}) *OneOfCondition {
	conds := make([]OneOfCondition, 0, len(m))
	for key, value := range m {
		conds = append(conds, OneOfCondition{
			Binary: &BinaryCondition{
				Field:    key,
				Operator: query.Equal,
				Value:    value,
			},
		})
	}
	return &OneOfCondition{Logic: &LogicCondition{
		Operator:   query.And,
		Conditions: conds,
	}}
}
