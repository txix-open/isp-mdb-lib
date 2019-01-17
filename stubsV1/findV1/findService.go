package findV1

import (
	"encoding/xml"
	"github.com/integration-system/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type LogicOperation string

const (
	LogicOperationAND LogicOperation = "AND"

	LogicOperationOR LogicOperation = "OR"
)

type FindObjects struct {
	XMLName xml.Name `xml:"http://find2.ws.elk.itb.ru/ findObjects"`

	Select *Select `xml:"select,omitempty"`

	SystemInfo *SystemInfo `xml:"systemInfo,omitempty"`
}

type Select struct {
	ObjectType string `xml:"objectType,omitempty"`

	Field []string `xml:"field,omitempty"`

	Filter *Filter `xml:"filter,omitempty"`

	Offset int32 `xml:"offset,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	AttachObjTypes []string `xml:"attachObjTypes,omitempty"`

	AttachRefTypes []string `xml:"attachRefTypes,omitempty"`

	NoOtherValues bool `xml:"noOtherValues,omitempty"`
}

type Filter struct {
	LogicOper *LogicOper `xml:"logicOper,omitempty"`

	Relation []string `xml:"relation,omitempty"`
}

type LogicOper struct {
	Oper *LogicOperation `xml:"oper,omitempty"`

	Cond []*Condition `xml:"cond,omitempty"`
}

type Condition struct {
	BinOper *BinOper `xml:"binOper,omitempty"`

	LogicOper *LogicOper `xml:"logicOper,omitempty"`
}

type BinOper struct {
	Field string `xml:"field,omitempty"`

	Oper string `xml:"oper,omitempty"`

	Value string `xml:"value,omitempty"`
}

type SystemInfo struct {
	From string `xml:"from,omitempty"`

	To string `xml:"to,omitempty"`

	MessageId string `xml:"messageId,omitempty"`

	SrcMessageId string `xml:"srcMessageId,omitempty"`

	SentDateTime time.Time `xml:"sentDateTime,omitempty"`

	Priority int32 `xml:"priority,omitempty"`
}

type FindObjectsResponse struct {
	XMLName xml.Name `xml:"http://find2.ws.elk.itb.ru/ findObjectsResponse"`

	Found *FindObjectsResult `xml:"found,omitempty"`
}

type FindObjectsResult struct {
	SystemInfo *SystemInfo `xml:"systemInfo,omitempty"`

	Response *Response `xml:"response,omitempty"`

	Objects *FoundObjects `xml:"objects,omitempty"`
}

type Response struct {
	Code int32 `xml:"code"`

	Description string `xml:"description,omitempty"`
}

type FoundObjects struct {
	Object []*MdmObject `xml:"object,omitempty"`
}

type MdmObject struct {
	ObjectType string `xml:"objectType,omitempty"`

	RelationType string `xml:"relationType,omitempty"`

	Attribute []*Attribute `xml:"attribute,omitempty"`

	Relations *Relations `xml:"relations,omitempty"`
}

type Attribute struct {
	Name string `xml:"name,omitempty"`

	Value []string `xml:"value,omitempty"`
}

type Relations struct {
	Object []*MdmObject `xml:"object,omitempty"`

	Ref []*Reference `xml:"ref,omitempty"`
}

type Reference struct {
	ObjectType string `xml:"objectType,omitempty"`

	RelationType string `xml:"relationType,omitempty"`

	ObjectId string `xml:"objectId,omitempty"`
}

type Find interface {
	FindObjects(request *FindObjects) (*FindObjectsResponse, error)
}

type find struct {
	client *soap.Client
}

func NewFind(client *soap.Client) Find {
	return &find{
		client: client,
	}
}

func (service *find) FindObjects(request *FindObjects) (*FindObjectsResponse, error) {
	response := new(FindObjectsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
