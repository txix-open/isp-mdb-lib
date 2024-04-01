package structure

import "github.com/txix-open/isp-mdb-lib/query"

type OneOfCondition struct {
	Logic  *LogicCondition  //логическое условие
	Binary *BinaryCondition //бинарное условие
}

type BinaryCondition struct {
	Field    string         `valid:"required~Required"`                                                                                                                //путь до атрибута
	Operator query.Operator `valid:"required~Required,in(=|!=|<|>|<=|>=|contains|not contains|starts with|not starts with|ends with|not ends with|exists|not exists)"` //условный оператор
	Value    interface{}    //(строка, число, булево значение) значение атрибута, обязательно для всех операторов, кроме (exists, not exists)
}

type LogicCondition struct {
	Operator   query.LogicOperator `valid:"required~Required,in(AND|OR)"` //логический оператор
	Conditions []OneOfCondition    `valid:"required~Required"`            //список условий для объединения
}
