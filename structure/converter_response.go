package structure

import (
	"fmt"
	"github.com/integration-system/isp-mdb-lib/query"
	"github.com/integration-system/isp-mdb-lib/stubsV1/erl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	emptyList = make([]ConvertResponse, 0)
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
	GetResult(appId int32) []ConvertResponse
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

func (r BatchListConvertForSudirResponse) GetResult(appId int32) []ConvertResponse {
	if arr, ok := r[appId]; ok {
		res := make([]ConvertResponse, len(arr))
		for i, v := range arr {
			res[i] = v
		}
		return res
	} else {
		return emptyList
	}
}

type BatchConvertForFindResponse map[int32]*FindConvertResult
type BatchListConvertForFindResponse map[int32][]FindConvertResult

func (r BatchListConvertForFindResponse) GetResult(appId int32) []ConvertResponse {
	if arr, ok := r[appId]; ok {
		res := make([]ConvertResponse, len(arr))
		for i, v := range arr {
			res[i] = v
		}
		return res
	} else {
		return emptyList
	}
}

type BatchConvertAnyResponse map[int32]*AnyConvertResult
type BatchListConvertAnyResponse map[int32][]AnyConvertResult

func (r BatchListConvertAnyResponse) GetResult(appId int32) []ConvertResponse {
	if arr, ok := r[appId]; ok {
		res := make([]ConvertResponse, len(arr))
		for i, v := range arr {
			res[i] = v
		}
		return res
	} else {
		return emptyList
	}
}

type BatchConvertErlResponse map[int32]*ErlConvertResult
type BatchListConvertErlResponse map[int32][]ErlConvertResult

func (r BatchListConvertErlResponse) GetResult(appId int32) []ConvertResponse {
	if arr, ok := r[appId]; ok {
		res := make([]ConvertResponse, len(arr))
		for i, v := range arr {
			res[i] = v
		}
		return res
	} else {
		return emptyList
	}
}

type BatchFilterDataResponse map[int32]*FilterDataResult
type BatchListFilterDataResponse map[int32][]FilterDataResult

func (r BatchListFilterDataResponse) GetResult(appId int32) []ConvertResponse {
	if arr, ok := r[appId]; ok {
		res := make([]ConvertResponse, len(arr))
		for i, v := range arr {
			res[i] = v
		}
		return res
	} else {
		return emptyList
	}
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
	Error     *ConvertSearchError
}

type ConvertAnySearchResponse struct {
	Condition         *query.Term
	UnavailableFields []string
}
