package structure

import "fmt"

type ExecuteByIdRequest struct {
	Id  string `valid:"required~Required"`
	Arg interface{}
}

type BatchExecuteByIdsRequest struct {
	Ids []string `valid:"required~Required"`
	Arg interface{}
}

type ExecuteRequest struct {
	Script string `valid:"required~Required"`
	Arg    interface{}
}

const (
	ErrorTypeCompile = "Compilation"
	ErrorTypeRunTime = "Runtime"
)

type ScriptResponse struct {
	Result interface{}
	Error  *ScriptResponseError
}

type ScriptResponseError struct {
	Type        string
	Description string
}

func (err ScriptResponseError) Error() string {
	return fmt.Sprintf("%s: %s", err.Type, err.Description)
}
