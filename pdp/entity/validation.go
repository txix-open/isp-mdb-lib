package entity

import (
	"database/sql/driver"
	"time"

	"github.com/txix-open/isp-kit/json"

	"github.com/txix-open/isp-mdb-lib/delta"
	"github.com/txix-open/isp-mdb-lib/pdp/consts"
)

type Validation struct {
	Id                  string
	ExternalId          string
	DebugRequestId      string
	SsoId               string
	ItemId              string
	SourceApplicationId int
	BlockName           string
	Changelogs          Changelogs
	GeneratedChangelogs Changelogs // сгенерированные из record ченжлоги, необходимые для валидации
	ValidationStatus    consts.ValidationStatus
	ValidationSystem    consts.ValidationSystem
	ValidationCode      string
	Request             json.RawMessage  `swaggertype:"object"`
	Response            *json.RawMessage `swaggertype:"object"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Metadata            *json.RawMessage `swaggertype:"object"`
}

type Changelogs []delta.Changelog

func (c *Changelogs) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}

func (c Changelogs) Value() (driver.Value, error) {
	bytes, err := json.Marshal(c)
	return driver.Value(bytes), err
}
