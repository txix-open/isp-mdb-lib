package structure

import (
	"errors"
	"fmt"

	"github.com/integration-system/isp-mdb-lib/pdp/codes"
	"github.com/integration-system/isp-mdb-lib/structure"
	"github.com/integration-system/isp-mdb-lib/stubsV2/findV2"
	"github.com/integration-system/isp-mdb-lib/stubsV2/sudirV2"
)

type PayloadResponse struct {
	ConvertRequest *structure.ConvertRequestPayload
	Error          *codes.PdpError
}

func GetResponseWithCode(code codes.SudirPDPCode, errArgs ...interface{}) *sudirV2.ResponseType {
	if errorDesc, ok := codes.GetDescByCode(code); ok {
		return &sudirV2.ResponseType{
			Response_Code:        int32(code),
			Response_Description: fmt.Sprintf(errorDesc, errArgs...),
		}
	}

	desc, _ := codes.GetDescByCode(codes.InternalError)
	return &sudirV2.ResponseType{
		Response_Code:        int32(codes.InternalError),
		Response_Description: desc,
	}
}

func GetResponseWithError(err error) *sudirV2.ResponseType {
	var pdpErr *codes.PdpError
	if ok := errors.As(err, &pdpErr); ok {
		return &sudirV2.ResponseType{
			Response_Code:        int32(pdpErr.GetCode()),
			Response_Description: pdpErr.GetMessage(),
		}
	}

	desc, _ := codes.GetDescByCode(codes.InternalError)
	return &sudirV2.ResponseType{
		Response_Code:        int32(codes.InternalError),
		Response_Description: desc,
	}
}

func GetFindResponseWithCode(code codes.FindPDPCode, errArgs ...interface{}) *findV2.Response {
	if errorDesc, ok := codes.GetDescByFindCode(code); ok {
		return &findV2.Response{
			Code:        int32(code),
			Description: fmt.Sprintf(errorDesc, errArgs...),
		}
	}

	desc, _ := codes.GetDescByFindCode(codes.FindInternalError)
	return &findV2.Response{
		Code:        int32(codes.FindInternalError),
		Description: desc,
	}
}
