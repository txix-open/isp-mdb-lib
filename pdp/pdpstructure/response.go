package pdpstructure

import (
	"errors"
	"fmt"

	"github.com/integration-system/isp-mdb-lib/pdp/pdpcodes"
	"github.com/integration-system/isp-mdb-lib/structure"
	"github.com/integration-system/isp-mdb-lib/stubsV2/findV2"
	"github.com/integration-system/isp-mdb-lib/stubsV2/sudirV2"
)

type PayloadResponse struct {
	ConvertRequest *structure.ConvertRequestPayload
	Error          *pdpcodes.PdpError
}

func GetResponseWithCode(code pdpcodes.SudirPDPCode, errArgs ...interface{}) *sudirV2.ResponseType {
	if errorDesc, ok := pdpcodes.GetDescByCode(code); ok {
		return &sudirV2.ResponseType{
			Response_Code:        int32(code),
			Response_Description: fmt.Sprintf(errorDesc, errArgs...),
		}
	}

	desc, _ := pdpcodes.GetDescByCode(pdpcodes.InternalError)
	return &sudirV2.ResponseType{
		Response_Code:        int32(pdpcodes.InternalError),
		Response_Description: desc,
	}
}

func GetResponseWithError(err error) *sudirV2.ResponseType {
	var pdpErr *pdpcodes.PdpError
	if ok := errors.As(err, &pdpErr); ok {
		return &sudirV2.ResponseType{
			Response_Code:        int32(pdpErr.GetCode()),
			Response_Description: pdpErr.GetMessage(),
		}
	}

	desc, _ := pdpcodes.GetDescByCode(pdpcodes.InternalError)
	return &sudirV2.ResponseType{
		Response_Code:        int32(pdpcodes.InternalError),
		Response_Description: desc,
	}
}

func GetFindResponseWithCode(code pdpcodes.FindPDPCode, errArgs ...interface{}) *findV2.Response {
	if errorDesc, ok := pdpcodes.GetDescByFindCode(code); ok {
		return &findV2.Response{
			Code:        int32(code),
			Description: fmt.Sprintf(errorDesc, errArgs...),
		}
	}

	desc, _ := pdpcodes.GetDescByFindCode(pdpcodes.FindInternalError)
	return &findV2.Response{
		Code:        int32(pdpcodes.FindInternalError),
		Description: desc,
	}
}
