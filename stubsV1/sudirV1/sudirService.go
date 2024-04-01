package sudirV1

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type EntryType struct {
	EntryName string `xml:"EntryName"`

	Attribute []*Attribute `xml:"Attribute"`
}

type EntryListType struct {
	EntryItem []*EntryType `xml:"EntryItem"`
}

type Attribute struct {
	Name string `xml:"Name"`

	Value []string `xml:"Value"`
}

type ResponseType struct {
	Response_Code int32 `xml:"Response_Code"`

	Response_Description string `xml:"Response_Description"`
}

type AddEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector AddEntryRequest"`

	EntryName string `xml:"EntryName"`

	Attribute []*Attribute `xml:"Attribute"`
}

type AddEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector AddEntryResponse"`

	Response_Code int32 `xml:"Response_Code"`

	Response_Description string `xml:"Response_Description"`
}

type FindEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector FindEntryRequest"`

	EntryItem *EntryType `xml:"EntryItem"`
}

type FindEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector FindEntryResponse"`

	Response *ResponseType `xml:"Response"`

	EntryItem *EntryType `xml:"EntryItem"`
}

type DeleteEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector DeleteEntryRequest"`

	EntryItem *EntryType `xml:"EntryItem"`
}

type DeleteEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector DeleteEntryResponse"`

	Response *ResponseType `xml:"Response"`

	EntryItem *EntryType `xml:"EntryItem"`
}

type UpdateEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector UpdateEntryRequest"`

	EntryItem *EntryType `xml:"EntryItem"`
}

type UpdateEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector UpdateEntryResponse"`

	Response *ResponseType `xml:"Response"`

	EntryItem *EntryType `xml:"EntryItem"`
}

type SelectEntriesRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector SelectEntriesRequest"`

	EntryTypeName string `xml:"EntryTypeName"`
}

type SelectEntriesResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector SelectEntriesResponse"`

	Response *ResponseType `xml:"Response"`
}

type GetNextEntriesRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector GetNextEntriesRequest"`

	Length int32 `xml:"Length"`
}

type GetNextEntriesResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector GetNextEntriesResponse"`

	Response *ResponseType `xml:"Response"`

	EntryList *EntryListType `xml:"EntryList"`
}
