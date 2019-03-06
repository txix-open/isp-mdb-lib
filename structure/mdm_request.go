package structure

import "github.com/integration-system/isp-mdb-lib/entity"

type DataRecordByExternalId struct {
	ExternalId string `json:"externalId" valid:"required~Required"`
	TechRecord bool
}

type DataRecordsByExternalIdList struct {
	ExternalIdList []string `valid:"required~Required"`
	TechRecord     bool
}

type MdmHandleRequest struct {
	TechRecord  bool
	OperationId string
	Record      *entity.DataRecord `valid:"required~Required"`
}
