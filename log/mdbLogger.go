package log

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
)

type Phase string

type DataType string

const (
	ReceiveData      = Phase("kri_in")
	FindEntrySudirV1 = Phase("kri_findentry_sudir1")
	AddEntrySudirV1  = Phase("kri_addentry_sudir1")
	UpdateSudirV1    = Phase("kri_update_sudir1")
	UpdateSudirV2    = Phase("kri_update_sudir2")
	UpdateErl        = Phase("kri_update_erl")

	ApiFindEntrySudirV1 = Phase("kri_api_findentry_sudir1")
	ApiFindEntrySudirV2 = Phase("kri_api_findentry_sudir2")
	ApiFindSudirV1      = Phase("kri_api_find_objects_sudir1")
	ApiFindSudirV2      = Phase("kri_api_find_objects_sudir2")

	JsonDataType = DataType("json")
	XmlDataType  = DataType("xml")

	componentId = "id_kri"
)

var (
	errUnknownDataType = "mdbLog: unknown data type %s"
	errEmptyData       = errors.New("mdbLog: empty data")

	json = jsoniter.ConfigFastest
)

// Create a new formatted log entry for mdb
// data - can be one of any type, if it is not formatted, it will format to specified dataType
func MakeLogEntry(phase Phase, operationId string, data interface{}, dataType DataType) ([]byte, error) {
	return makeForResponse(make(map[string]interface{}, 4), phase, operationId, data, dataType)
}

func MakeLogEntryWithRequest(
	phase Phase,
	operationId string,
	requestData interface{},
	requestType DataType,
	responseData interface{},
	responseType DataType,
) ([]byte, error) {
	s, err := formatData(requestData, requestType)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{}, 5)
	if requestType == JsonDataType {
		m["requestJson"] = s
	} else if requestType == XmlDataType {
		m["requestXML"] = s
	} else {
		return nil, fmt.Errorf(errUnknownDataType, requestType)
	}

	return makeForResponse(m, phase, operationId, responseData, responseType)
}

func makeForResponse(m map[string]interface{}, phase Phase, operationId string, data interface{}, dataType DataType) ([]byte, error) {
	s, err := formatData(data, dataType)
	if err != nil {
		return nil, err
	}

	m["componentID"] = componentId
	m["phaseID"] = phase
	m["operationID"] = operationId
	if dataType == JsonDataType {
		m["json"] = s
	} else if dataType == XmlDataType {
		m["xml"] = s
	} else {
		return nil, fmt.Errorf(errUnknownDataType, dataType)
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func formatData(data interface{}, dataType DataType) (string, error) {
	var s string

	switch d := data.(type) {
	case []byte:
		s = string(d)
	case string:
		s = d
	default:
		var bytes []byte
		var err error
		if dataType == JsonDataType {
			bytes, err = json.Marshal(data)
		} else if dataType == XmlDataType {
			bytes, err = xml.Marshal(data)
		}

		if err != nil {
			return "", err
		}

		s = string(bytes)
	}

	if s == "" {
		return "", errEmptyData
	}

	return s, nil
}
