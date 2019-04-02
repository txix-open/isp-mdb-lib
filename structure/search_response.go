package structure

import "github.com/integration-system/isp-mdb-lib/entity"

type SearchResponse struct {
	Items      []entity.DataRecord
	TotalCount int64
}

type CountResponse struct {
	TotalCount int64
}

type SearchIdResponse struct {
	Items      []string
	TotalCount int64
}

type SearchIdWithScrollResponse struct {
	SearchIdResponse
	ScrollId string
}

type PreferredSearchSlicesResponse struct {
	MaxSlices int
}
