package structure

import "github.com/integration-system/isp-mdb-lib/query"

type OneOfCondition struct {
	Logic  *LogicCondition  //логическая условие
	Binary *BinaryCondition //бинарное условия
}

type BinaryCondition struct {
	Field    string              `valid:"required~Required"`                                                                                                                //путь до атрибута
	Operator query.LogicOperator `valid:"required~Required,in(=|!=|<|>|<=|>=|contains|not contains|starts with|not starts with|ends with|not ends with|exists|not exists)"` //условный оператор
	Value    string              //значение атрибута, обязательно для всех операторов, кроме (exists, not exists)
}

type LogicCondition struct {
	Operator   query.LogicOperator `valid:"required~Required,in(AND|OR)"` //логический оператор
	Conditions []OneOfCondition    `valid:"required~Required"`            //список условий для объединения
}
