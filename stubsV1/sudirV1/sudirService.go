package sudirV1

import (
	"encoding/xml"
	"github.com/integration-system/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type EntryType struct {
	EntryName string `xml:"EntryName,omitempty"`

	Attribute []*Attribute `xml:"Attribute,omitempty"`
}

type EntryListType struct {
	EntryItem []*EntryType `xml:"EntryItem,omitempty"`
}

type Attribute struct {
	Name string `xml:"Name"`

	Value []string `xml:"Value"`
}

type ResponseType struct {
	Response_Code int32 `xml:"Response_Code"`

	Response_Description string `xml:"Response_Description,omitempty"`
}

type FindEntryRequestType struct {
	XMLName   xml.Name   `xml:"http://xmlns.dit.mos.ru/sudir/connector FindEntryRequest"`
	EntryName string     `xml:"EntryName,omitempty"`
	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type FindEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector FindEntryResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type DeleteEntryRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector DeleteEntryRequest"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type DeleteEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector DeleteEntryResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type AddEntryRequestType struct {
	XMLName   xml.Name     `xml:"http://xmlns.dit.mos.ru/sudir/connector AddEntryRequest"`
	EntryName string       `xml:"EntryName,omitempty"`
	Attribute []*Attribute `xml:"Attribute,omitempty"`
}

type UpdateEntryRequestType struct {
	XMLName   xml.Name     `xml:"http://xmlns.dit.mos.ru/sudir/connector UpdateEntryRequest"`
	EntryName string       `xml:"EntryName,omitempty"`
	Attribute []*Attribute `xml:"Attribute,omitempty"`
}

type UpdateEntryResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector UpdateEntryResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryItem *EntryType `xml:"EntryItem,omitempty"`
}

type SelectEntriesRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector SelectEntriesRequest"`

	EntryTypeName string `xml:"EntryTypeName,omitempty"`
}

type SelectEntriesResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector SelectEntriesResponse"`

	Response *ResponseType `xml:"Response,omitempty"`
}

type GetNextEntriesRequestType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector GetNextEntriesRequest"`

	Length int32 `xml:"Length,omitempty"`
}

type GetNextEntriesResponseType struct {
	XMLName xml.Name `xml:"http://xmlns.dit.mos.ru/sudir/connector GetNextEntriesResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	EntryList *EntryListType `xml:"EntryList,omitempty"`
}

type SudirPortType interface {
	AddEntry(request *AddEntryRequestType) (*ResponseType, error)

	AddEntryWithInterceptor(
		request *AddEntryRequestType,
		interceptor func(request string, response string),
	) (*ResponseType, error)

	FindEntry(request *FindEntryRequestType) (*FindEntryResponseType, error)

	FindEntryWithInterceptor(
		request *FindEntryRequestType,
		interceptor func(request string, response string),
	) (*FindEntryResponseType, error)

	DeleteEntry(request *DeleteEntryRequestType) (*DeleteEntryResponseType, error)

	UpdateEntry(request *UpdateEntryRequestType) (*UpdateEntryResponseType, error)

	UpdateEntryWithInterceptor(
		request *UpdateEntryRequestType,
		interceptor func(request string, response string),
	) (*UpdateEntryResponseType, error)

	SelectEntries(request *SelectEntriesRequestType) (*SelectEntriesResponseType, error)

	GetNextEntries(request *GetNextEntriesRequestType) (*GetNextEntriesResponseType, error)
}

type sudirPortType struct {
	client *soap.Client
}

func NewSudirPortType(client *soap.Client) SudirPortType {
	return &sudirPortType{
		client: client,
	}
}

func (service *sudirPortType) AddEntry(request *AddEntryRequestType) (*ResponseType, error) {
	response := new(ResponseType)
	err := service.client.Call("urn::#AddEntry", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) AddEntryWithInterceptor(
	request *AddEntryRequestType,
	interceptor func(request string, response string),
) (*ResponseType, error) {
	response := new(ResponseType)
	err := service.client.CallWithInterceptor("urn::#AddEntry", request, response, interceptor)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) FindEntry(request *FindEntryRequestType) (*FindEntryResponseType, error) {
	response := new(FindEntryResponseType)
	err := service.client.Call("urn::#FindEntry", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) FindEntryWithInterceptor(
	request *FindEntryRequestType,
	interceptor func(request string, response string),
) (*FindEntryResponseType, error) {
	response := new(FindEntryResponseType)
	err := service.client.CallWithInterceptor("urn::#FindEntry", request, response, interceptor)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) DeleteEntry(request *DeleteEntryRequestType) (*DeleteEntryResponseType, error) {
	response := new(DeleteEntryResponseType)
	err := service.client.Call("urn::#DeleteEntry", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) UpdateEntry(request *UpdateEntryRequestType) (*UpdateEntryResponseType, error) {
	response := new(UpdateEntryResponseType)
	err := service.client.Call("urn::#UpdateEntry", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) UpdateEntryWithInterceptor(
	request *UpdateEntryRequestType,
	interceptor func(request string, response string),
) (*UpdateEntryResponseType, error) {
	response := new(UpdateEntryResponseType)
	err := service.client.CallWithInterceptor("urn::#UpdateEntry", request, response, interceptor)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) SelectEntries(request *SelectEntriesRequestType) (*SelectEntriesResponseType, error) {
	response := new(SelectEntriesResponseType)
	err := service.client.Call("urn::#SelectEntries", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sudirPortType) GetNextEntries(request *GetNextEntriesRequestType) (*GetNextEntriesResponseType, error) {
	response := new(GetNextEntriesResponseType)
	err := service.client.Call("urn::#GetNextEntries", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
