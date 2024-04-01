package filter

import (
	"strings"

	u "github.com/txix-open/isp-mdb-lib/utils"
)

type Filter struct {
	wlWildcard bool
	blWildcard bool
	wl         []string
	bl         []string
}

func (f *Filter) Apply(ss []string) []string {
	filtered := make([]string, 0, len(ss))
	for _, s := range ss {
		if f.Check(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func (f *Filter) Check(field string) bool {
	if f.blWildcard {
		return false
	}

	if f.wlWildcard {
		if len(f.bl) == 0 {
			return true
		}
	}

	return checkRules(field, f.wl) && !checkRules(field, f.bl)
}

func (f *Filter) IsWhitelistWildcard() bool {
	return f.wlWildcard && !f.blWildcard
}

func NewFilter(availableFields, excludedFields []string) *Filter {
	return &Filter{
		wl:         availableFields,
		bl:         excludedFields,
		wlWildcard: hasWildcard(availableFields),
		blWildcard: hasWildcard(excludedFields),
	}
}

func MatchPath(requestPath string, availableField string) bool {
	if availableField == u.AttrAcceptAny {
		return true
	}
	reqPathArray := strings.Split(requestPath, ".")
	availablePathArray := strings.Split(availableField, ".")
	availPathLastPartIndex := len(availablePathArray) - 1
	for i := range reqPathArray {
		reqPart := reqPathArray[i]
		availPart := availablePathArray[i]
		isLastPart := availPathLastPartIndex == i
		if isLastPart && reqPart == availPart {
			return true
		}
		if reqPart != availPart {
			return false
		}
	}
	return false
}

func checkRules(requestPath string, availableFields []string) bool {
	for _, availablePath := range availableFields {
		if MatchPath(requestPath, availablePath) {
			return true
		}
	}
	return false
}

func hasWildcard(ss []string) bool {
	for _, path := range ss {
		if path == u.AttrAcceptAny {
			return true
		}
	}
	return false
}
