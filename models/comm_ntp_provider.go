package models

import (


	"github.com/beevik/etree"

)

const CommNtpProviderClassName = "commNtpProvider"

type CommNtpProvider struct {
    BaseAttributes
    CommNtpProviderAttributes
}

type CommNtpProviderAttributes struct{
    Admin_state string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Hostname string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewCommNtpProvider(commNtpProviderRn,description string, commNtpProviderAttributes CommNtpProviderAttributes) *CommNtpProvider {
    dn := commNtpProviderRn
	return &CommNtpProvider{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         CommNtpProviderClassName,
			Status:            "created, modified",
			Description:       description,
		},

		CommNtpProviderAttributes: commNtpProviderAttributes,
	}
}

func (commNtpProvider *CommNtpProvider) ToMap() (map[string]string, error) {
	commNtpProviderMap, err := commNtpProvider.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    commNtpProviderMap["adminState"] = commNtpProvider.Admin_state
    commNtpProviderMap["childAction"] = commNtpProvider.Child_action
    commNtpProviderMap["hostname"] = commNtpProvider.Hostname
    commNtpProviderMap["sacl"] = commNtpProvider.Sacl
    
	return commNtpProviderMap, nil

}

func CommNtpProviderFromDoc(doc *etree.Document, rootClass string) *CommNtpProvider {
	element, err := GetMoElement(doc, rootClass, CommNtpProviderClassName)
	if err != nil {
		return nil
	}
	return &CommNtpProvider{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         CommNtpProviderClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		CommNtpProviderAttributes: CommNtpProviderAttributes{

        
			Admin_state:  element.SelectAttrValue("adminState", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Hostname:  element.SelectAttrValue("hostname", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}