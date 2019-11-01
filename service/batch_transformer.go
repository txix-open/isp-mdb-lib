package service

import (
	"errors"
	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/structure"
	"github.com/integration-system/isp-mdb-lib/utils"
	"reflect"
)

var (
	unknownError = errors.New("results for specified system identity not found int response")
)

type convertFunc func(list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error)

type BatchTransformer struct {
	service    ConverterService
	converters map[structure.ProtocolVersion]convertFunc
}

func (t BatchTransformer) ConvertBatch(protocol structure.ProtocolVersion, list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
	if convert, ok := t.converters[protocol]; ok {
		return convert(list, identity, techRecord)
	} else {
		return convertAny(t.service, protocol, list, identity, techRecord)
	}
}

func (t BatchTransformer) FilterDataBatch(list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
	request := make([]structure.BatchFilterDataRequest, len(list))
	for i, record := range list {
		request[i] = structure.BatchFilterDataRequest{
			Record:                      &structure.Record{Data: record.Data, CustomData: record.CustomData},
			AbstractConvertBatchRequest: makeAbstractRequest("", record, identity, techRecord),
		}
	}

	response, err := t.service.FilterData(request)
	if err != nil {
		return nil, err
	}

	return castResultArray(response[identity])
}

func NewBatchTransformer(service ConverterService) BatchTransformer {
	bt := BatchTransformer{service: service}
	bt.converters = map[structure.ProtocolVersion]convertFunc{
		structure.SudirV1:     makeSudirConverter(service, structure.SudirV1),
		structure.SudirV2:     makeSudirConverter(service, structure.SudirV2),
		structure.SudirV1Find: makeSudirFindConverter(service, structure.SudirV1Find),
		structure.SudirV2Find: makeSudirFindConverter(service, structure.SudirV2Find),
		structure.ErlProtocol: makeErlConverter(service),
	}
	return bt
}

func convertSudir(service ConverterService, protocol structure.ProtocolVersion, list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
	request := make([]structure.BatchConvertDataRequest, len(list))
	for i, record := range list {
		request[i] = structure.BatchConvertDataRequest{
			ConvertRequestPayload: &structure.ConvertRequestPayload{
				Data:                  record.Data,
				CustomData:            record.CustomData,
				AttachedObjectTypes:   []string{utils.AllObjectTypes},
				FilterByAttachedTypes: protocol == structure.SudirV2,
			},
			AbstractConvertBatchRequest: makeAbstractRequest(protocol, record, identity, techRecord),
		}
	}

	response, err := service.ConvertToSudir(request)
	if err != nil {
		return nil, err
	}

	return castResultArray(response[identity])
}

func convertSudirFind(service ConverterService, protocol structure.ProtocolVersion, list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
	request := make([]structure.BatchConvertForFindServiceRequest, len(list))
	for i, record := range list {
		request[i] = structure.BatchConvertForFindServiceRequest{
			ConvertForFindServiceRequestPayload: &structure.ConvertForFindServiceRequestPayload{
				Records:             []*structure.Record{{Data: record.Data, CustomData: record.CustomData}},
				AttachedObjectTypes: []string{utils.AllObjectTypes},
			},
			AbstractConvertBatchRequest: makeAbstractRequest(protocol, record, identity, techRecord),
		}
	}

	response, err := service.ConvertToSudirFindEntry(request)
	if err != nil {
		return nil, err
	}

	return castResultArray(response[identity])
}

func convertErl(service ConverterService, list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
	request := make([]structure.BatchConvertErlRequest, len(list))
	for i, record := range list {
		request[i] = structure.BatchConvertErlRequest{
			Record:                      &structure.Record{Data: record.Data, CustomData: record.CustomData},
			AbstractConvertBatchRequest: makeAbstractRequest(structure.ErlProtocol, record, identity, techRecord),
		}
	}

	response, err := service.ConvertToErl(request)
	if err != nil {
		return nil, err
	}

	return castResultArray(response[identity])
}

func convertAny(service ConverterService, protocol structure.ProtocolVersion, list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
	request := make([]structure.BatchConvertAnyRequest, len(list))
	for i, record := range list {
		request[i] = structure.BatchConvertAnyRequest{
			Record:                      &structure.Record{Data: record.Data, CustomData: record.CustomData},
			AbstractConvertBatchRequest: makeAbstractRequest(protocol, record, identity, techRecord),
		}
	}

	response, err := service.ConvertToJson(request)
	if err != nil {
		return nil, err
	}

	return castResultArray(response[identity])
}

func makeErlConverter(service ConverterService) convertFunc {
	return func(list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
		return convertErl(service, list, identity, techRecord)
	}
}

func makeSudirConverter(service ConverterService, protocol structure.ProtocolVersion) convertFunc {
	return func(list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
		return convertSudir(service, protocol, list, identity, techRecord)
	}
}

func makeSudirFindConverter(service ConverterService, protocol structure.ProtocolVersion) convertFunc {
	return func(list []entity.TransitDataRecord, identity int32, techRecord bool) ([]structure.ConvertResponse, error) {
		return convertSudirFind(service, protocol, list, identity, techRecord)
	}
}

func makeAbstractRequest(protocol structure.ProtocolVersion, record entity.TransitDataRecord, identity int32, techRecord bool) *structure.AbstractConvertBatchRequest {
	return &structure.AbstractConvertBatchRequest{
		AppIdList:  []structure.NotificationTarget{{AppId: identity}},
		Id:         uint64(record.Id),
		ExternalId: record.ExternalId,
		Version:    record.Version,
		IsTech:     techRecord,
		Protocol:   protocol,
	}
}

func castResultArray(result interface{}) ([]structure.ConvertResponse, error) {
	if result == nil {
		return nil, unknownError
	}
	rv := reflect.ValueOf(result)
	if rv.Kind() == reflect.Slice {
		l := rv.Len()
		arr := make([]structure.ConvertResponse, l)
		for i := 0; i < l; i++ {
			value := rv.Index(i)
			if value.IsValid() {
				if convertResponse, ok := value.Interface().(structure.ConvertResponse); ok {
					arr[i] = convertResponse
				} else {
					return nil, unknownError
				}
			} else {
				return nil, unknownError
			}
		}
		return arr, nil
	} else {
		return nil, unknownError
	}
}
