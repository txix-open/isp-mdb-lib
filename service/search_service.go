package service

import (
	"github.com/integration-system/isp-lib/backend"
	"github.com/integration-system/isp-mdb-lib/structure"
)

type SearchService struct {
	client   *backend.RxGrpcClient
	method   string
	callerId int
}

func (s *SearchService) Search(req structure.SearchRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertSearch(req, &res)
}

func (s *SearchService) SearchIdList(req structure.SearchRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertSearchIdList(req, &res)
}

func (s *SearchService) SearchIdWithScroll(req structure.SearchWithScrollRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertSearchIdWithScroll(req, &res)
}

func (s *SearchService) GetPreferredSlicesCount(isTech bool) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.convertGetPreferredSlicesCount(isTech, &res)
}

func (s *SearchService) convertSearch(req structure.SearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			s.method,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertSearchIdList(req structure.SearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			s.method,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertSearchIdWithScroll(req structure.SearchWithScrollRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			s.method,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *SearchService) convertGetPreferredSlicesCount(isTech bool, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			s.method,
			s.callerId,
			structure.PreferredSearchSlicesRequest{IsTech: isTech},
			resPtr,
		)
	})
}

func NewSeachService(client *backend.RxGrpcClient, method string, callerId int) SearchService {
	return SearchService{
		client:   client,
		method:   method,
		callerId: callerId,
	}
}
