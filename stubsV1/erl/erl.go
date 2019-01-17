package erl

import (
	"encoding/xml"
)

type PersonsIncoming struct {
	XMLName        xml.Name         `xml:"http://erl.msr.com/schemas/master-system/elc persons_incoming"`
	PersonIncoming []PersonIncoming `xml:"person_incoming,omitempty"`
}

type PersonIncoming struct {
	Citizen           Citizen           `xml:"citizen,omitempty"`
	Inn               string            `xml:"inn,omitempty"`
	IdentityDocuments IdentityDocuments `xml:"identity_documents,omitempty"`
	Oms               string            `xml:"oms,omitempty"`
}

type Citizen struct {
	CitizenPk string      `xml:"citizen_pk,omitempty"`
	Name      CitizenName `xml:"name,omitempty"`
	Birthday  string      `xml:"birthday,omitempty"`
	Sex       string      `xml:"sex,omitempty"`
	Snils     string      `xml:"snils,omitempty"`
}

type CitizenName struct {
	Surname    string `xml:"surname,omitempty"`
	Firstname  string `xml:"firstname,omitempty"`
	Patronymic string `xml:"patronymic,omitempty"`
}

type IdentityDocuments struct {
	Document []Document `xml:"document,omitempty"`
}

type Document struct {
	DoctypePk      string `xml:"doctype_pk,omitempty"`
	Serial         string `xml:"serial,omitempty"`
	Number         string `xml:"number,omitempty"`
	DocIssuedate   string `xml:"doc_issuedate,omitempty"`
	DocIssuePlace  string `xml:"doc_issue_place,omitempty"`
	DocauthorityPk string `xml:"docauthority_pk,omitempty"`
	IsDropped      string `xml:"is_dropped,omitempty"`
}
