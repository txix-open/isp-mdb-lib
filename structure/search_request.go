package structure

import (
	"github.com/integration-system/isp-mdb-lib/query"
	"time"
)

type SearchRequest struct {
	Limit     int
	Offset    int
	IsTech    bool
	Condition query.Term
}

type CountRequest struct {
	IsTech    bool
	Condition query.Term
}

type SearchWithScrollRequest struct {
	IsTech    bool
	Condition query.Term
	BatchSize int `valid:"required~Required"`
	ScrollId  string
	ScrollTTL time.Duration `valid:"required~Required"`
	Slicing   *struct {
		SliceId   int
		MaxSlices int
	}
}

type PreferredSearchSlicesRequest struct {
	IsTech bool
}
