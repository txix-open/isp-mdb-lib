package structure

import (
	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/stubsV2/findV2"
)

type NotificationTarget struct {
	AppId    int32
	Notifier string
}

type AbstractConvertBatchRequest struct {
	ExternalId string
	Id         uint64
	Version    int64
	Protocol   ProtocolVersion
	IsTech     bool
	AppIdList  []NotificationTarget `valid:"required~Required"`
}

type ConvertRequestPayload struct {
	Data                  interface{} `valid:"required~Required"`
	CustomData            interface{}
	AttachedObjectTypes   []string
	FilterByAttachedTypes bool
}

func (p ConvertRequestPayload) CastToMaps() (map[string]interface{}, map[string]interface{}) {
	data, _ := p.Data.(map[string]interface{})
	customData, _ := p.CustomData.(map[string]interface{})
	return data, customData
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
	Data       interface{} `valid:"required~Required"`
	CustomData interface{}
}

func (r Record) CastToMaps() (map[string]interface{}, map[string]interface{}) {
	data, _ := r.Data.(map[string]interface{})
	customData, _ := r.CustomData.(map[string]interface{})
	return data, customData
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

type MappingTypeRequest struct {
	ApplicationId int32           `valid:"required~Required"`
	Protocol      ProtocolVersion `valid:"required~Required"`
}

type MappingRequest struct {
	ApplicationId int32              `valid:"required~Required"`
	Protocol      ProtocolVersion    `valid:"required~Required"`
	Attributes    []entity.Attribute `valid:"required~Required"`
}
