package structure

import (
	"fmt"

	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/query"
	"github.com/integration-system/isp-mdb-lib/stubsV1/erl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConvertError struct {
	Code  codes.Code
	Error string
}

func (e *ConvertError) ToGrpcError() error {
	return status.Error(e.Code, e.Error)
}

type ConvertResponse interface {
	GetMetadata() AbstractConvertResult
	GetResult() interface{}
}

type BatchConvertResponse interface {
	GetResult() []ConvertResponse
}

type AbstractConvertResult struct {
	Target     NotificationTarget
	Error      *ConvertError
	ExternalId string
	Id         uint64
	Version    int64
	Protocol   ProtocolVersion
	IsTech     bool
}

type SudirConvertResult struct {
	Result *EntryType
	AbstractConvertResult
}

func (r SudirConvertResult) GetMetadata() AbstractConvertResult {
	return r.AbstractConvertResult
}

func (r SudirConvertResult) GetResult() interface{} {
	return r.Result
}

type FindConvertResult struct {
	Result []*MdmObject
	AbstractConvertResult
}

func (r FindConvertResult) GetMetadata() AbstractConvertResult {
	return r.AbstractConvertResult
}

func (r FindConvertResult) GetResult() interface{} {
	return r.Result
}

type AnyConvertResult struct {
	Result map[string]interface{}
	AbstractConvertResult
}

func (r AnyConvertResult) GetMetadata() AbstractConvertResult {
	return r.AbstractConvertResult
}

func (r AnyConvertResult) GetResult() interface{} {
	return r.Result
}

type ErlConvertResult struct {
	Result *erl.PersonsIncoming
	AbstractConvertResult
}

func (r ErlConvertResult) GetMetadata() AbstractConvertResult {
	return r.AbstractConvertResult
}

func (r ErlConvertResult) GetResult() interface{} {
	return r.Result
}

type FilterDataResult struct {
	Result *Record
	AbstractConvertResult
}

func (r FilterDataResult) GetMetadata() AbstractConvertResult {
	return r.AbstractConvertResult
}

func (r FilterDataResult) GetResult() interface{} {
	return r.Result
}

type BatchConvertForSudirResponse map[int32]*SudirConvertResult
type BatchListConvertForSudirResponse map[int32][]SudirConvertResult

func (r BatchListConvertForSudirResponse) GetResult() []ConvertResponse {
	result := make([]ConvertResponse, 0, len(r)*2)
	for _, arr := range r {
		for _, resp := range arr {
			result = append(result, resp)
		}
	}
	return result
}

type BatchConvertForFindResponse map[int32]*FindConvertResult
type BatchListConvertForFindResponse map[int32][]FindConvertResult

func (r BatchListConvertForFindResponse) GetResult() []ConvertResponse {
	result := make([]ConvertResponse, 0, len(r)*2)
	for _, arr := range r {
		for _, resp := range arr {
			result = append(result, resp)
		}
	}
	return result
}

type BatchConvertAnyResponse map[int32]*AnyConvertResult
type BatchListConvertAnyResponse map[int32][]AnyConvertResult

func (r BatchListConvertAnyResponse) GetResult() []ConvertResponse {
	result := make([]ConvertResponse, 0, len(r)*2)
	for _, arr := range r {
		for _, resp := range arr {
			result = append(result, resp)
		}
	}
	return result
}

type BatchConvertErlResponse map[int32]*ErlConvertResult
type BatchListConvertErlResponse map[int32][]ErlConvertResult

func (r BatchListConvertErlResponse) GetResult() []ConvertResponse {
	result := make([]ConvertResponse, 0, len(r)*2)
	for _, arr := range r {
		for _, resp := range arr {
			result = append(result, resp)
		}
	}
	return result
}

type BatchFilterDataResponse map[int32]*FilterDataResult
type BatchListFilterDataResponse map[int32][]FilterDataResult

func (r BatchListFilterDataResponse) GetResult() []ConvertResponse {
	result := make([]ConvertResponse, 0, len(r)*2)
	for _, arr := range r {
		for _, resp := range arr {
			result = append(result, resp)
		}
	}
	return result
}

type Reason int

const (
	ReasonUnsupported Reason = iota
	ReasonWrongValue
	ReasonNotAccepted
	ReasonEmptyCond
	ReasonNoValue
)

type ConvertSearchError struct {
	Field  string
	Reason Reason
}

func (e ConvertSearchError) Error() string {
	return fmt.Sprintf("%s: %d", e.Field, e.Reason)
}

type ConvertSearchResponse struct {
	Condition *query.Term
	Type      string
	Error     *ConvertSearchError
}

type ConvertAnySearchResponse struct {
	Condition         *query.Term
	Type              string
	UnavailableFields []string
}

type MappingTypeResponse struct {
	Type string
}

type SudirUpdateRecordRequest struct {
	TechRecord       bool
	Record           *entity.DataRecord `valid:"required~Required"`
	SoftDelete       bool
	DeleteOperations map[string]map[string]string
	Error            *ConvertSearchError
}

type ConvertPayloadResponse struct {
	ConvertRequest *ConvertRequestPayload
	Error          error
}
