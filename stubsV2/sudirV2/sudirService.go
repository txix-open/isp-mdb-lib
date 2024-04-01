package sudirV2

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SystemInfo struct {
	From string `xml:"From,omitempty"`

	To string `xml:"To,omitempty"`

	MessageId string `xml:"MessageId,omitempty"`

	SrcMessageId string `xml:"SrcMessageId,omitempty"`

	SentDateTime time.Time `xml:"SentDateTime,omitempty"`

	Priority int32 `xml:"Priority,omitempty"`

	ReqSeq int64 `xml:"ReqSeq,omitempty"`

	ExchKey string `xml:"ExchKey,omitempty"`

	SessKey string `xml:"SessKey,omitempty"`

	Digest string `xml:"Digest,omitempty"`
}

type EntryType struct {
	EntryName string `xml:"EntryName"`

	Seq int64 `xml:"Seq,omitempty"`

	Attribute []*Attribute `xml:"Attribute"`

	Object []*Object `xml:"Object"`
}

type EntryListType struct {
	EntryItem []*EntryType `xml:"EntryItem,omitempty"`
}

type Attribute struct {
	Name  string   `xml:"Name"`
	Value []string `xml:"Value"`
}

type Object struct {
	Name      string       `xml:"Name"`
	Attribute []*Attribute `xml:"Attribute"`
}

type ResponseType struct {
	Response_Code int32 `xml:"Response_Code"`

	Response_Description string `xml:"Response_Description,omitempty"`

	Response_ErrorList []string `xml:"Response_ErrorList,omitempty"`
}

type AddEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector AddEntryResponse"`

	*ResponseType `xml:"Response,omitempty"`
}

type FindEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector FindEntryRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	GetPending string `xml:"GetPending,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`

	AttachObjTypes []string `xml:"AttachObjTypes,omitempty"`
}

type FindEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector FindEntryResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type DeleteEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector DeleteEntryRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type DeleteEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector DeleteEntryResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type UpdateEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector UpdateEntryRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	ComplexOperId string `xml:"ComplexOperId,omitempty"`

	EntryItem []*EntryType `xml:"EntryItem,omitempty"`
}

type UpdateEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector UpdateEntryResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem []*EntryType `xml:"EntryItem,omitempty"`
}

type SelectEntriesRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector SelectEntriesRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	EntryTypeName string `xml:"EntryTypeName,omitempty"`
}

type SelectEntriesResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector SelectEntriesResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`
}

type GetNextEntriesRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector GetNextEntriesRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Length int32 `xml:"Length,omitempty"`
}

type GetNextEntriesResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector GetNextEntriesResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryList *EntryListType `xml:"EntryList,omitempty"`
}

type GetSessKeyRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector GetSessKeyRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Sample string `xml:"Sample,omitempty"`
}

type GetSessKeyResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector GetSessKeyResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`

	SessKey string `xml:"SessKey,omitempty"`
}

type AddEntryExRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector AddEntryExRequest"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	ComplexOperId string `xml:"ComplexOperId,omitempty"`

	EntryItem []*EntryType `xml:"EntryItem,omitempty"`
}

type AddEntryExResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/itb/connector AddEntryExResponse"`

	SystemInfo *SystemInfo `xml:"SystemInfo,omitempty"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem []*EntryType `xml:"EntryItem,omitempty"`
}

type SudirItbPortType interface {
	AddEntry(request *EntryType) (*ResponseType, error)

	FindEntry(request *FindEntryRequestType) (*FindEntryResponseType, error)

	DeleteEntry(request *DeleteEntryRequestType) (*DeleteEntryResponseType, error)

	UpdateEntry(request *UpdateEntryRequestType) (*UpdateEntryResponseType, error)

	UpdateEntryWithInterceptor(
		request *UpdateEntryRequestType,
		interceptor func(request string, response string),
	) (*UpdateEntryResponseType, error)

	SelectEntries(request *SelectEntriesRequestType) (*SelectEntriesResponseType, error)

	GetNextEntries(request *GetNextEntriesRequestType) (*GetNextEntriesResponseType, error)

	GetSessKey(request *GetSessKeyRequestType) (*GetSessKeyResponseType, error)

	AddEntryEx(request *AddEntryExRequestType) (*AddEntryExResponseType, error)
}
