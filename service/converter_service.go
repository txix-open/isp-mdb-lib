package service

import (
	"github.com/integration-system/isp-lib/v2/backend"
	"github.com/integration-system/isp-lib/v2/modules"
	"github.com/integration-system/isp-mdb-lib/structure"
)

const (
	convertAnySearchMethod = "mdm-converter/any/convert_search_request"
)

type ConverterService struct {
	client   *backend.RxGrpcClient
	callerId int
}

func (s *ConverterService) ConvertToSudirFindEntry(req []structure.BatchConvertForFindServiceRequest) (
	structure.BatchListConvertForFindResponse, error,
) {

	res := make(structure.BatchListConvertForFindResponse)
	return res, s.convertFind(req, &res)
}

func (s *ConverterService) ConvertToSudir(req []structure.BatchConvertDataRequest) (structure.BatchListConvertForSudirResponse, error) {
	res := make(structure.BatchListConvertForSudirResponse)
	return res, s.convert(req, &res)
}

func (s *ConverterService) ConvertToJson(req []structure.BatchConvertAnyRequest) (structure.BatchListConvertAnyResponse, error) {
	res := make(structure.BatchListConvertAnyResponse)
	return res, s.convertJson(req, &res)
}

func (s *ConverterService) ConvertToErl(req []structure.BatchConvertErlRequest) (structure.BatchListConvertErlResponse, error) {
	res := make(structure.BatchListConvertErlResponse)
	return res, s.convertErl(req, &res)
}

func (s *ConverterService) FilterData(req []structure.BatchFilterDataRequest) (structure.BatchListFilterDataResponse, error) {
	res := make(structure.BatchListFilterDataResponse)
	return res, s.filterData(req, &res)
}

func (s *ConverterService) FilterSearchRequest(req structure.FilterSearchRequest) (*structure.ConvertAnySearchResponse, error) {
	res := new(structure.ConvertAnySearchResponse)
	return res, s.filterSearchRequest(req, res)
}

func (s *ConverterService) ConvertAnySearchRequest(req structure.ConvertAnySearchRequest) (*structure.ConvertAnySearchResponse, error) {
	res := new(structure.ConvertAnySearchResponse)
	err := s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(convertAnySearchMethod, s.callerId, req, res)
	})
	return res, err
}

func (s *ConverterService) convertFind(req []structure.BatchConvertForFindServiceRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmDumperLinks.MdmConverterService.ConvertToFindBatchList,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *ConverterService) convert(req []structure.BatchConvertDataRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmDumperLinks.MdmConverterService.ConvertToSudirBatchList,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *ConverterService) convertJson(req []structure.BatchConvertAnyRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmDumperLinks.MdmConverterService.ConvertAnyBatchList,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *ConverterService) convertErl(req []structure.BatchConvertErlRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmDumperLinks.MdmConverterService.ConvertErlBatchList,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *ConverterService) filterData(req []structure.BatchFilterDataRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmDumperLinks.MdmConverterService.FilterBatchList,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func (s *ConverterService) filterSearchRequest(req structure.FilterSearchRequest, resPtr interface{}) error {
	return s.client.Visit(func(c *backend.InternalGrpcClient) error {
		return c.Invoke(
			modules.MdmApiLinks.MdmConverterService.FilterSearchRequest,
			s.callerId,
			req,
			resPtr,
		)
	})
}

func NewConverterService(client *backend.RxGrpcClient, callerId int) ConverterService {
	return ConverterService{
		client:   client,
		callerId: callerId,
	}
}
