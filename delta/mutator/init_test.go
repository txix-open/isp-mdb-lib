package mutator_test

import (
	"encoding/json"
	"testing"

	"github.com/integration-system/isp-mdb-lib/delta/mutator"
	"github.com/stretchr/testify/suite"
)

func TestServiceApply_Suite(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestServiceApply{})
}

type TestServiceApply struct {
	suite.Suite
	service mutator.Service
}

func (t *TestServiceApply) SetupSuite() {
	t.service = mutator.NewService()
}

func (t *TestServiceApply) makeData(msg []byte) map[string]any {
	data := make(map[string]any)
	err := json.Unmarshal(msg, &data)
	t.Require().NoError(err)
	return data
}
