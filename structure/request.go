package structure

import (
	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/query"
	"github.com/integration-system/isp-mdb-lib/stubsV2/findV2"
)

type Mode string

const (
	Initializing Mode = "INITIALIZING"
	Default      Mode = "DEFAULT"
	SwitchedOff  Mode = "OFF"
)

type ChangeModRequest struct {
	Mode Mode `valid:"in(INITIALIZING|DEFAULT|OFF),required~Required"`
}

type DataRecordByExternalId struct {
	ExternalId string `json:"externalId" valid:"required~Required"`
	TechRecord bool
}

type MdmHandleRequest struct {
	TechRecord  bool
	OperationId string
	Record      *entity.DataRecord `valid:"required~Required"`
}

type Identities struct {
	Ids []int64 `json:"ids" valid:"required~Required"`
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
	ExternalId             string
	Id                     uint64
	Version                int64
	Protocol               ProtocolVersion
	IsTech                 bool
	AppIdList              []int32
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
	ExternalId                           string
	Id                                   uint64
	Version                              int64
	Protocol                             ProtocolVersion `valid:"required~Required"`
	IsTech                               bool
	AppIdList                            []int32
}

type ConvertAnyRequest struct {
	Record        *Record         `valid:"required~Required"`
	ApplicationId int32           `valid:"required~Required"`
	Protocol      ProtocolVersion `valid:"required~Required"`
}

type BatchConvertAnyRequest struct {
	Record     *Record `valid:"required~Required"`
	AppIdList  []int32
	ExternalId string
	Id         uint64
	Version    int64
	IsTech     bool
	Protocol   ProtocolVersion `valid:"required~Required"`
}

type ConvertErlRequest struct {
	Record        *Record `valid:"required~Required"`
	ApplicationId int32   `valid:"required~Required"`
}

type BatchConvertErlRequest struct {
	Record     *Record `valid:"required~Required"`
	AppIdList  []int32 `valid:"required~Required"`
	ExternalId string
	Id         uint64
	Version    int64
	IsTech     bool
}

type FilterDataRequest struct {
	Record        *Record `valid:"required~Required"`
	ApplicationId int32   `valid:"required~Required"`
}

type BatchДFilterDataRequest struct {
	Record     *Record `valid:"required~Required"`
	AppIdList  []int32 `valid:"required~Required"`
	ExternalId string
	Id         uint64
	Version    int64
	IsTech     bool
}

type SearchRequest struct {
	Limit     int `valid:"range(0|1000)"`
	Offset    int
	IsTech    bool
	Condition query.Term `valid:"required~Required"`
}
