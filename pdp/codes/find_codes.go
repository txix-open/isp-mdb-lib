package codes

type FindPDPCode int32

const (
	FindOK                     FindPDPCode = 0
	FindNotFound               FindPDPCode = 4
	FindNoAttributeMapping     FindPDPCode = 7
	FindWrongAttributeValue    FindPDPCode = 10
	FindNoAttributeParams      FindPDPCode = 11
	FindNoAttributeProfile     FindPDPCode = 12
	FindNotSupported           FindPDPCode = 14
	FindNoAttributeValue       FindPDPCode = 104
	FindExtSystemNotRegistered FindPDPCode = 111
	FindWrongObjectType        FindPDPCode = 112
	FindNoAnyCond              FindPDPCode = 113
	FindInternalError          FindPDPCode = -1
	FindNotAccessError         FindPDPCode = -2
)

var findErrorDescList = map[FindPDPCode]string{
	FindOK:                     "OK",
	FindNotFound:               "Запись не найдена в каталоге",
	FindNoAttributeMapping:     "Атрибут %s не поддерживается",
	FindWrongAttributeValue:    "Атрибут %s имеет некорректное значение",
	FindNoAttributeParams:      "Для атрибута %s не указаны параметры",
	FindNoAttributeProfile:     "Для атрибута %s не указан профиль",
	FindNotSupported:           "Метод не поддерживается",
	FindNoAttributeValue:       "Нет значения для атрибута %s",
	FindExtSystemNotRegistered: "Внешняя ИС (%s) не зарегистрирована",
	FindWrongObjectType:        "Неверный тип объекта %s",
	FindNoAnyCond:              "Не указано ни одного условия поиска",
	FindInternalError:          "Внутренняя ошибка сервиса",
	FindNotAccessError:         "Сервис временно недоступен",
}

func GetDescByFindCode(code FindPDPCode) (string, bool) {
	if desc, ok := findErrorDescList[code]; ok {
		return desc, ok
	}

	return "", false
}
