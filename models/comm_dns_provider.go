package models

import (


	"github.com/beevik/etree"

)

const CommDnsProviderClassName = "commDnsProvider"

type CommDnsProvider struct {
    BaseAttributes
    CommDnsProviderAttributes
}

type CommDnsProviderAttributes struct{
    Admin_state string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Hostname string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewCommDnsProvider(commDnsProviderRn,description string, commDnsProviderAttributes CommDnsProviderAttributes) *CommDnsProvider {
    dn := commDnsProviderRn
	return &CommDnsProvider{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         CommDnsProviderClassName,
			Status:            "created, modified",
			Description:       description,
		},

		CommDnsProviderAttributes: commDnsProviderAttributes,
	}
}

func (commDnsProvider *CommDnsProvider) ToMap() (map[string]string, error) {
	commDnsProviderMap, err := commDnsProvider.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    commDnsProviderMap["adminState"] = commDnsProvider.Admin_state
    commDnsProviderMap["childAction"] = commDnsProvider.Child_action
    commDnsProviderMap["hostname"] = commDnsProvider.Hostname
    commDnsProviderMap["sacl"] = commDnsProvider.Sacl
    
	return commDnsProviderMap, nil

}

func CommDnsProviderFromDoc(doc *etree.Document, rootClass string) *CommDnsProvider {
	element, err := GetMoElement(doc, rootClass, CommDnsProviderClassName)
	if err != nil {
		return nil
	}
	return &CommDnsProvider{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         CommDnsProviderClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		CommDnsProviderAttributes: CommDnsProviderAttributes{

        
			Admin_state:  element.SelectAttrValue("adminState", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Hostname:  element.SelectAttrValue("hostname", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}