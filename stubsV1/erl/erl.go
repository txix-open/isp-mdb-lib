package erl

import (
	"encoding/xml"
)

type StringWithLastModified struct {
	Value        string `xml:",chardata"`
	Lastmodified string `xml:"last_modified,attr"`
}

type PersonsIncoming struct {
	XMLName        xml.Name         `xml:"http://erl.msr.com/schemas/master-system/elc persons_incoming"`
	PersonIncoming []PersonIncoming `xml:"person_incoming,omitempty"`
}

type PersonIncoming struct {
	ProviderIdentifier string                  `xml:"provider_identifier,omitempty"`
	Citizen            Citizen                 `xml:"citizen,omitempty"`
	IdentityDocuments  *IdentityDocuments      `xml:"identity_documents,omitempty"`
	Inn                *StringWithLastModified `xml:"inn,omitempty"`
	Oms                *StringWithLastModified `xml:"oms,omitempty"`
	Death              *Death                  `xml:"death,omitempty"`
	//Divorce            *DivorceData            `xml:"divorce,omitempty"`
	CitizenAddresses *CitizenAddresses `xml:"citizen_addresses,omitempty"`
}

type CitizenAddresses struct {
	CitizenAddress []CitizenAddress `xml:"citizen_address,omitempty"`
}

type CitizenAddress struct {
	AddressType      string                  `xml:"address_type,omitempty"`
	Address          *Address                `xml:"address,omitempty"`
	RegistrationDate *StringWithLastModified `xml:"registration_date,omitempty"`
	DepartureDate    *StringWithLastModified `xml:"departure_date,omitempty"`
	Lastmodified     string                  `xml:"last_modified,attr"`
}

type Address struct {
	BtiIdentity    string `xml:"bti_identity,omitempty"`
	PostIndex      string `xml:"post_index,omitempty"`
	SettlementCode string `xml:"settlement_code,omitempty"`
	StreetCode     string `xml:"street_code,omitempty"`
	HouseCode      string `xml:"house_code,omitempty"`
	House          string `xml:"house,omitempty"`
	Corpus         string `xml:"corpus,omitempty"`
	Building       string `xml:"building,omitempty"`
	Flat           string `xml:"flat,omitempty"`
}

type DivorceData struct {
	Document        *SimpleDocument         `xml:"document,omitempty"`
	DivorceDate     *StringWithLastModified `xml:"divorce_date,omitempty"`
	SourceValueCode string                  `xml:"source_value_code,omitempty"`
}

type Death struct {
	Document        *SimpleDocument         `xml:"document,omitempty"`
	DeathDate       *StringWithLastModified `xml:"death_date,omitempty"`
	SourceValueCode string                  `xml:"source_value_code,omitempty"`
}

type SimpleDocument struct {
	Serial           string `xml:"serial,omitempty"`
	Number           string `xml:"number,omitempty"`
	DocIssuedate     string `xml:"doc_issuedate,omitempty"`
	DocauthorityName string `xml:"docauthority_name,omitempty"`
}

type Citizen struct {
	CitizenPk string                  `xml:"citizen_pk,omitempty"`
	Name      CitizenName             `xml:"name,omitempty"`
	Birthday  *StringWithLastModified `xml:"birthday,omitempty"`
	Sex       *StringWithLastModified `xml:"sex,omitempty"`
	Snils     *StringWithLastModified `xml:"snils,omitempty"`
}

type CitizenName struct {
	Surname      string `xml:"surname,omitempty"`
	Firstname    string `xml:"firstname,omitempty"`
	Patronymic   string `xml:"patronymic,omitempty"`
	Lastmodified string `xml:"last_modified,attr"`
}

type IdentityDocuments struct {
	Document []CitizenDocument `xml:"document,omitempty"`
}

type CitizenDocument struct {
	DoctypePk      string `xml:"doctype_pk,omitempty"`
	Serial         string `xml:"serial,omitempty"`
	Number         string `xml:"number,omitempty"`
	DocIssuedate   string `xml:"doc_issuedate,omitempty"`
	DocIssuePlace  string `xml:"doc_issue_place,omitempty"`
	DocauthorityPk string `xml:"docauthority_pk,omitempty"`
	IsDropped      string `xml:"is_dropped,omitempty"`
	Lastmodified   string `xml:"last_modified,attr"`
}
