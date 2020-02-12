package service

import (
	"github.com/integration-system/isp-lib/http"
	"github.com/integration-system/isp-mdb-lib/structure"
)

const (
	urlApiServiceSyncGetByIds        = "/sync/get_by_ids"
	urlApiServiceSyncFindIds         = "/sync/find_ids"
	urlApiServiceSyncFindIdsExtended = "/sync/find_ids_extended"
	urlApiServiceSyncFind            = "/sync/find"
	urlApiServiceSyncFindExtended    = "/sync/find_extended"
)

type ApiService struct {
	client  http.RestClient
	headers map[string]string
}

func (s ApiService) SyncGetByIdList(request structure.IdListApiRequest) (*structure.ApiResponse, error) {
	response := new(structure.ApiResponse)
	err := s.client.Invoke("POST", urlApiServiceSyncGetByIds, s.headers, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s ApiService) SyncFindIdList(request structure.SyncSearchApiRequest) (*structure.ApiResponse, error) {
	response := new(structure.ApiResponse)
	err := s.client.Invoke("POST", urlApiServiceSyncFindIds, s.headers, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s ApiService) SyncFindIdListExtended(request structure.SyncSearchExtendedApiRequest) (*structure.ApiResponse, error) {
	response := new(structure.ApiResponse)
	err := s.client.Invoke("POST", urlApiServiceSyncFindIdsExtended, s.headers, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s ApiService) SyncFind(request structure.SyncSearchApiRequest) (*structure.ApiResponse, error) {
	response := new(structure.ApiResponse)
	err := s.client.Invoke("POST", urlApiServiceSyncFind, s.headers, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s ApiService) SyncFindExtended(request structure.SyncSearchExtendedApiRequest) (*structure.ApiResponse, error) {
	response := new(structure.ApiResponse)
	err := s.client.Invoke("POST", urlApiServiceSyncFindExtended, s.headers, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewApiService(client http.RestClient, headers map[string]string) ApiService {
	return ApiService{
		client:  client,
		headers: headers,
	}
}
