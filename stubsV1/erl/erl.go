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
	Citizen           Citizen                 `xml:"citizen,omitempty"`
	Inn               *StringWithLastModified `xml:"inn,omitempty"`
	Oms               *StringWithLastModified `xml:"oms,omitempty"`
	IdentityDocuments IdentityDocuments       `xml:"identity_documents,omitempty"`
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
	Lastmodified   string `xml:"last_modified,attr"`
}
