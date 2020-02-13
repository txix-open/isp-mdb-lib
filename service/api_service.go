package service

import (
	"github.com/integration-system/isp-lib/v2/http"
	"github.com/integration-system/isp-lib/v2/utils"
	"github.com/integration-system/isp-mdb-lib/structure"
)

const (
	urlApiServiceSyncGetByIds        = "/api/mdm-api/sync/get_by_ids"
	urlApiServiceSyncFindIds         = "/api/mdm-api/sync/find_ids"
	urlApiServiceSyncFindIdsExtended = "/api/mdm-api/sync/find_ids_extended"
	urlApiServiceSyncFind            = "/api/mdm-api/sync/find"
	urlApiServiceSyncFindExtended    = "/api/mdm-api/sync/find_extended"
)

type ApiService struct {
	client  http.RestClient
	address string
	headers map[string]string
}

func (s *ApiService) SyncGetByIdList(request structure.IdListApiRequest, response *structure.ApiResponse) error {
	return s.doRequest(urlApiServiceSyncGetByIds, request, response)
}

func (s *ApiService) SyncFindIdList(request structure.SyncSearchApiRequest, response *structure.ApiResponse) error {
	return s.doRequest(urlApiServiceSyncFindIds, request, response)
}

func (s *ApiService) SyncFindIdListExtended(request structure.SyncSearchExtendedApiRequest, response *structure.ApiResponse) error {
	return s.doRequest(urlApiServiceSyncFindIdsExtended, request, response)
}

func (s *ApiService) SyncFind(request structure.SyncSearchApiRequest, response *structure.ApiResponse) error {
	return s.doRequest(urlApiServiceSyncFind, request, response)
}

func (s *ApiService) SyncFindExtended(request structure.SyncSearchExtendedApiRequest, response *structure.ApiResponse) error {
	return s.doRequest(urlApiServiceSyncFindExtended, request, response)
}

func (s *ApiService) doRequest(method string, request interface{}, response *structure.ApiResponse) error {
	err := s.client.Invoke("POST", s.address+method, s.headers, request, response)
	if err != nil {
		return err
	} else if response.GetError() != nil {
		return err
	}
	return nil
}

func NewApiService(client http.RestClient, address string, applicationToken string) *ApiService {
	return &ApiService{
		client:  client,
		address: address,
		headers: map[string]string{
			utils.ApplicationTokenHeader: applicationToken,
		},
	}
}
