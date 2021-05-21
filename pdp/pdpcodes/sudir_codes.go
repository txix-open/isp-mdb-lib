package pdpcodes

type SudirPDPCode int32

const (
	OK                        SudirPDPCode = 0
	EntryNotSupported         SudirPDPCode = 2
	NotFound                  SudirPDPCode = 4
	NoAttributeMapping        SudirPDPCode = 7
	WrongAttributeValue       SudirPDPCode = 10
	SSOAlreadyExists          SudirPDPCode = 12
	NotSupported              SudirPDPCode = 14
	AttributeTypeNotSupported SudirPDPCode = 20
	NoAttributeParams         SudirPDPCode = 21
	NoAttributeProfile        SudirPDPCode = 22
	AbsentEntryName           SudirPDPCode = 101
	SSONotFound               SudirPDPCode = 102
	SSONotProvided            SudirPDPCode = 103
	CantCreateUser            SudirPDPCode = 107
	AttributeNotAccepted      SudirPDPCode = 110
	ExtSystemNotRegistered    SudirPDPCode = 111
	UnableExtUpload           SudirPDPCode = 112
	CantDeleteUserWoSNILS     SudirPDPCode = 121
	CantDeletePassWoPass      SudirPDPCode = 122
	CantUpdPassWoSNILS        SudirPDPCode = 123
	CantDeleteSNILS           SudirPDPCode = 124
	CantUpdSNILSWoFIO         SudirPDPCode = 132
	CantUpdSNILSWoBirth       SudirPDPCode = 133
	CantUpdPassWoData         SudirPDPCode = 134
	PhoneExists               SudirPDPCode = 135
	RelativeExists            SudirPDPCode = 136
	IncorrectData             SudirPDPCode = 137
	InternalError             SudirPDPCode = -1
	NotAccessError            SudirPDPCode = -2
)

var errorDescList = map[SudirPDPCode]string{
	OK:                        "OK",
	EntryNotSupported:         "Тип сущности %s не поддерживается",
	NotFound:                  "Запись не найдена в каталоге",
	NoAttributeMapping:        "Атрибут %s не поддерживается",
	WrongAttributeValue:       "Атрибут %s имеет некорректное значение",
	SSOAlreadyExists:          "Запись с уникальным значением LoginName (%s) уже существует в каталоге",
	NotSupported:              "Метод не поддерживается",
	AttributeTypeNotSupported: "Тип атрибута %s не поддерживается",
	NoAttributeParams:         "Для атрибута %s не указаны параметры",
	NoAttributeProfile:        "Для атрибута %s не указаны параметры",
	AbsentEntryName:           "Отсутствует EntryName",
	SSONotFound:               "Не найден LoginName для обновления",
	SSONotProvided:            "Не указан LoginName для обновления",
	CantCreateUser:            "Не удалось создать пользователя значением LoginName (%s)",
	AttributeNotAccepted:      "Атрибут не допустим для клиента (%s)",
	ExtSystemNotRegistered:    "Внешняя ИС (%s) не зарегистрирована",
	UnableExtUpload:           "Для внешней ИС (%s) не разрешено внесение изменений",
	CantDeleteUserWoSNILS:     "ФИО и д.р. не могут быть удалены если есть СНИЛС",
	CantDeletePassWoPass:      "Паспортные данные не могут быть удалены при наличии паспорта",
	CantUpdPassWoSNILS:        "Паспортные данные не могут быть исправлены если нет СНИЛС",
	CantDeleteSNILS:           "СНИЛС не может быть удален",
	CantUpdSNILSWoFIO:         "СНИЛС не может быть исправлен если нет ФИО",
	CantUpdSNILSWoBirth:       "СНИЛС не может быть исправлен если нет Даты рождения",
	CantUpdPassWoData:         "Паспортные данные не могут быть исправлены, т.к. не хватает данных",
	PhoneExists:               "Введёный телефон уже есть в УЗ",
	RelativeExists:            "Аналогичные данные были введены Вами ранее для другого члена семьи.",
	IncorrectData:             "Один или несколько параметров неправильно заполнены",
	InternalError:             "Внутренняя ошибка сервиса",
	NotAccessError:            "Сервис временно недоступен",
}

func GetDescByCode(code SudirPDPCode) (string, bool) {
	if desc, ok := errorDescList[code]; ok {
		return desc, ok
	}

	return "", false
}
