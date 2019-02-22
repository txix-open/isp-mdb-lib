package structure

import (
	"fmt"
	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/query"
	"github.com/integration-system/isp-mdb-lib/stubsV1/erl"
	"google.golang.org/grpc/codes"
)

type ConvertError struct {
	Code  codes.Code
	Error string
}

type ConvertResponse interface {
	GetMetadata() AbstractConvertResult
	GetResult() interface{}
}

type AbstractConvertResult struct {
	ExternalId string
	Id         uint64
	Version    int64
	Protocol   ProtocolVersion
	IsTech     bool
	Error      *ConvertError
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

type BatchConvertForFindResponse map[int32]*FindConvertResult
type BatchListConvertForFindResponse map[int32][]FindConvertResult

type BatchConvertAnyResponse map[int32]*AnyConvertResult
type BatchListConvertAnyResponse map[int32][]AnyConvertResult

type BatchConvertErlResponse map[int32]*ErlConvertResult
type BatchListConvertErlResponse map[int32][]ErlConvertResult

type BatchFilterDataResponse map[int32]*FilterDataResult
type BatchListFilterDataResponse map[int32][]FilterDataResult

type UpsertResponse struct {
	OldValue *entity.DataRecord
	NewValue *entity.DataRecord
	Inserted bool
	Updated  bool
}

type SearchResponse struct {
	Items      []entity.DataRecord
	TotalCount int64
}

type UuidSearchResponse struct {
	Items      []string
	TotalCount int64
}

type Reason int

const (
	ReasonUnsupported = iota
	ReasonWrongValue
	ReasonNotAccepted
	ReasonEmptyCond
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

func GetPkFilterValue(cond *query.Term) string {
	if cond != nil && cond.IsValid() &&
		cond.IsLogic() && cond.LogicOperation.IsValid() &&
		len(cond.LogicOperation.Terms) == 1 {
		if cond.LogicOperation.Terms[0].IsBinary() && cond.LogicOperation.Terms[0].BinaryOperation.IsPrimaryKey { //and((id == ?))
			return cond.LogicOperation.Terms[0].BinaryOperation.Value
		} else if cond.LogicOperation.Terms[0].IsLogic() { //and(or((id1 == id), id2 == id)) case
			or := cond.LogicOperation.Terms[0].LogicOperation
			if or.IsValid() && len(or.Terms) > 0 && or.Terms[0].IsBinary() && or.Terms[0].BinaryOperation.IsPrimaryKey {
				pk := or.Terms[0].BinaryOperation.Value
				for i := 1; i < len(or.Terms); i++ {
					t := or.Terms[i]
					if t.IsBinary() && t.BinaryOperation.IsPrimaryKey && pk == t.BinaryOperation.Value {
						continue
					} else {
						return ""
					}
				}
				return pk
			}
		}
	}

	return ""
}
