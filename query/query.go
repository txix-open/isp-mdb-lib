package query

type Operator string

type LogicOperator string

const (
	Equal        Operator = "="
	NotEqual     Operator = "!="
	Lt           Operator = "<"
	Gt           Operator = ">"
	Lte          Operator = "<="
	Gte          Operator = ">="
	Constains    Operator = "contains"
	NotContains  Operator = "not contains"
	StartsWith   Operator = "starts with"
	NotStartWith Operator = "not starts with"
	EndWith      Operator = "ends with"
	NotEndWith   Operator = "not ends with"

	And LogicOperator = "AND"
	Or  LogicOperator = "OR"
)

type Term struct {
	LogicOperation  *LogicOperation
	BinaryOperation *BinaryOperation
}

func (t Term) IsLogic() bool {
	return t.LogicOperation != nil
}

func (t Term) IsBinary() bool {
	return t.BinaryOperation != nil
}

func (t Term) IsValid() bool {
	return t.BinaryOperation != nil || t.LogicOperation != nil
}

type BinaryOperation struct {
	Operator       Operator
	Value          string
	Field          string
	SearchInCustom bool
	IsPrimaryKey   bool
}

func (bo BinaryOperation) IsValid() bool {
	return bo.Field != "" && bo.Operator != "" && bo.Value != ""
}

type LogicOperation struct {
	LogicOperator LogicOperator
	Terms         []Term
}

func (lo LogicOperation) IsValid() bool {
	return lo.LogicOperator == And || lo.LogicOperator == Or
}
