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
	Exists       Operator = "exists"
	NotExists    Operator = "not exists"

	And LogicOperator = "AND"
	Or  LogicOperator = "OR"
)

type Term struct {
	LogicOperation  *LogicOperation  `schema:"Логическая операция"`
	BinaryOperation *BinaryOperation `schema:"Бинарная операция"`
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
	Operator        Operator `schema:"Оператор,[=,!=,<,>,<=,>=,contains,not contains, starts with, not starts with, ends with, not ends with, exists, not exists]"`
	Value           string   `schema:"Ожидаемое значение"`
	Field           string   `schema:"Операнд,название поля к которому применяется оператор"`
	MappingSrcField string   `schema:"Исходное название поля маппинга"`
	SearchInCustom  bool     `schema:"Поиск по расширенным данным,если используется поле из массива атрибутов ('documents.100016.ref_num'), то должно быть значение 'true'"`
	IsPrimaryKey    bool     `schema:"Первичны ключ,если поле является первичным ключом для профиля, то рекомендуется значение true, для оптимизации поиска"`
}

func (bo BinaryOperation) IsValid() bool {
	ok := bo.Field != "" && bo.Operator != ""
	if !ok {
		return false
	}
	return bo.Value != "" || (bo.Operator == Exists || bo.Operator == NotExists)
}

type LogicOperation struct {
	LogicOperator LogicOperator `schema:"Логический оператор,['AND', 'OR']"`
	Terms         []Term        `schema:"Список условий"`
}

func (lo LogicOperation) IsValid() bool {
	return lo.LogicOperator == And || lo.LogicOperator == Or
}
