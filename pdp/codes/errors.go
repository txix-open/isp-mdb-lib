package codes

import (
	"fmt"
)

type PdpError struct {
	Code SudirPDPCode
	Desc string
	Args []interface{}
}

func (e *PdpError) Error() string {
	return fmt.Sprintf(e.Desc, e.Args...)
}

func (e *PdpError) GetCode() SudirPDPCode {
	return e.Code
}

func (e *PdpError) GetMessage() string {
	return e.Error()
}

func NewError(code SudirPDPCode, args ...interface{}) error {
	if desc, ok := errorDescList[code]; ok && code != -1 {
		return &PdpError{code, desc, args}
	}

	return &PdpError{Code: InternalError, Desc: errorDescList[InternalError]}
}
