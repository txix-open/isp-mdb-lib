package structure

import (
	"github.com/txix-open/isp-mdb-lib/query"
)

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
