package service

import (
	"github.com/integration-system/isp-lib/v2/backend"
	"github.com/integration-system/isp-mdb-lib/structure"
)

const (
	ScriptServiceExecuteByIdAddr = "script/script/execute_by_id"
	ScriptServiceExecuteAddr     = "script/script/execute"
)

type ScriptService struct {
	client   *backend.RxGrpcClient
	callerId int
}

func (s *ScriptService) ExecuteById(req structure.ExecuteByIdRequest, responseResPtr interface{}) (*structure.ScriptResponse, error) {
	res := new(structure.ScriptResponse)
	res.Result = responseResPtr
	err := s.client.Invoke(
		ScriptServiceExecuteByIdAddr,
		s.callerId,
		req,
		res,
	)
	return res, err
}

func (s *ScriptService) Execute(req structure.ExecuteRequest, responseResPtr interface{}) (*structure.ScriptResponse, error) {
	res := new(structure.ScriptResponse)
	res.Result = responseResPtr
	err := s.client.Invoke(
		ScriptServiceExecuteAddr,
		s.callerId,
		req,
		res,
	)
	return res, err
}

func NewScriptService(client *backend.RxGrpcClient, callerId int) *ScriptService {
	return &ScriptService{
		client:   client,
		callerId: callerId,
	}
}
