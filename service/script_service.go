package service

import (
	"github.com/integration-system/isp-lib/v2/backend"
	"github.com/integration-system/isp-mdb-lib/structure"
)

const (
	ScriptServiceExecuteByIdAddr  = "script/script/execute_by_id"
	ScriptServiceExecuteAddr      = "script/script/execute"
	ScriptServiceBatchExecute     = "script/script/batch_execute"
	ScriptServiceBatchExecuteById = "script/script/batch_execute_by_ids"
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

func (s *ScriptService) BatchExecute(req []structure.ExecuteByIdRequest, responsePtr interface{}) error {
	err := s.client.Invoke(
		ScriptServiceBatchExecute,
		s.callerId,
		req,
		responsePtr,
	)
	return err
}

func (s *ScriptService) BatchExecuteById(req structure.BatchExecuteByIdsRequest, responsePtr interface{}) error {
	err := s.client.Invoke(
		ScriptServiceBatchExecuteById,
		s.callerId,
		req,
		responsePtr,
	)
	return err
}

func NewScriptService(client *backend.RxGrpcClient, callerId int) *ScriptService {
	return &ScriptService{
		client:   client,
		callerId: callerId,
	}
}
