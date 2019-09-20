package structure

import (
	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/stubsV2/findV2"
)

type AbstractConvertBatchRequest struct {
	ExternalId string
	Id         uint64
	Version    int64
	Protocol   ProtocolVersion
	IsTech     bool
	AppIdList  []int32 `valid:"required~Required"`
}

type ConvertRequestPayload struct {
	Data                  map[string]interface{} `valid:"required~Required"`
	CustomData            map[string]interface{}
	AttachedObjectTypes   []string
	FilterByAttachedTypes bool
}

type ConvertDataRequest struct {
	*ConvertRequestPayload `valid:"required~Required"`
	ApplicationId          int32 `valid:"required~Required"`
	Protocol               ProtocolVersion
}

type BatchConvertDataRequest struct {
	*ConvertRequestPayload `valid:"required~Required"`
	*AbstractConvertBatchRequest
}

type Record struct {
	Data       map[string]interface{} `valid:"required~Required"`
	CustomData map[string]interface{}
}

type ConvertForFindServiceRequestPayload struct {
	Records             []*Record
	AttachedObjectTypes []string
	AttachedRefType     []string
	ObjectType          string
}

type ConvertForFindServiceRequest struct {
	*ConvertForFindServiceRequestPayload `valid:"required~Required"`
	ApplicationId                        int32           `valid:"required~Required"`
	Protocol                             ProtocolVersion `valid:"required~Required"`
}

type ConvertSearchForFindServiceRequest struct {
	ApplicationId int32           `valid:"required~Required"`
	Protocol      ProtocolVersion `valid:"required~Required"`
	Source        *findV2.Filter
}

type ConvertSearchRequest struct {
	ApplicationId int32           `valid:"required~Required"`
	Protocol      ProtocolVersion `valid:"required~Required"`
	Source        *EntryType
}

type BatchConvertForFindServiceRequest struct {
	*ConvertForFindServiceRequestPayload `valid:"required~Required"`
	*AbstractConvertBatchRequest
}

type ConvertAnyRequest struct {
	Record        *Record         `valid:"required~Required"`
	ApplicationId int32           `valid:"required~Required"`
	Protocol      ProtocolVersion `valid:"required~Required"`
}

type BatchConvertAnyRequest struct {
	Record *Record `valid:"required~Required"`
	*AbstractConvertBatchRequest
}

type ConvertErlRequest struct {
	Record        *Record `valid:"required~Required"`
	ApplicationId int32   `valid:"required~Required"`
}

type BatchConvertErlRequest struct {
	Record *Record `valid:"required~Required"`
	*AbstractConvertBatchRequest
}

type FilterDataRequest struct {
	Record        *Record `valid:"required~Required"`
	ApplicationId int32   `valid:"required~Required"`
}

type BatchFilterDataRequest struct {
	Record *Record `valid:"required~Required"`
	*AbstractConvertBatchRequest
}

type FilterAttributeRequest struct {
	Attributes    []entity.Attribute `valid:"required~Required"`
	ApplicationId int32              `valid:"required~Required"`
}

type FilterSearchRequest struct {
	Query         map[string]interface{}
	ApplicationId int32 `valid:"required~Required"`
}

type ConvertAnySearchRequest struct {
	Condition     *OneOfCondition `valid:"required~Required"`
	ApplicationId int32           `valid:"required~Required"`
	Protocol      ProtocolVersion
}
