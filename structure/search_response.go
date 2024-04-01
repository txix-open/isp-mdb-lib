package structure

import "github.com/txix-open/isp-mdb-lib/entity"

type SearchResponse struct {
	Items      []entity.TransitDataRecord
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
