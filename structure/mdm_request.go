package structure

import "github.com/txix-open/isp-mdb-lib/entity"

type DataRecordByExternalId struct {
	ExternalId string `json:"externalId" valid:"required~Required"`
	TypeDescriptor
}

type DataRecordsByExternalIdList struct {
	ExternalIdList []string `valid:"required~Required"`
	TypeDescriptor
}

type MdmHandleRequest struct {
	TechRecord  bool
	OperationId string
	Record      *entity.DataRecord `valid:"required~Required"`
}

type MdmAttributeRequest struct {
	TypeName string
}
