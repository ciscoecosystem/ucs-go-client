package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const IppoolPoolClassName = "ippoolPool"

type IppoolPool struct {
    BaseAttributes
    IppoolPoolAttributes
}

type IppoolPoolAttributes struct{
    Assigned string `xml:",omitempty"`
    Assignment_order string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Ext_managed string `xml:",omitempty"`
    Guid string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Is_net_bios_enabled string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Size string `xml:",omitempty"`
    Supports_dhcp string `xml:",omitempty"`
    V4_assigned string `xml:",omitempty"`
    V4_size string `xml:",omitempty"`
    V6_assigned string `xml:",omitempty"`
    V6_size string `xml:",omitempty"`
    

}

func NewIppoolPool(ippoolPoolRn,parentDn,description string, ippoolPoolAttributes IppoolPoolAttributes) *IppoolPool { 
    dn := fmt.Sprintf("%s/%s", parentDn, ippoolPoolRn)
	return &IppoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         IppoolPoolClassName,
			Status:            "created, modified",
			Description:       description,
		},

		IppoolPoolAttributes: ippoolPoolAttributes,
	}
}

func (ippoolPool *IppoolPool) ToMap() (map[string]string, error) {
	ippoolPoolMap, err := ippoolPool.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    ippoolPoolMap["assigned"] = ippoolPool.Assigned
    ippoolPoolMap["assignmentOrder"] = ippoolPool.Assignment_order
    ippoolPoolMap["childAction"] = ippoolPool.Child_action
    ippoolPoolMap["extManaged"] = ippoolPool.Ext_managed
    ippoolPoolMap["guid"] = ippoolPool.Guid
    ippoolPoolMap["intId"] = ippoolPool.Int_id
    ippoolPoolMap["isNetBIOSEnabled"] = ippoolPool.Is_net_bios_enabled
    ippoolPoolMap["policyLevel"] = ippoolPool.Policy_level
    ippoolPoolMap["policyOwner"] = ippoolPool.Policy_owner
    ippoolPoolMap["propAcl"] = ippoolPool.Prop_acl
    ippoolPoolMap["sacl"] = ippoolPool.Sacl
    ippoolPoolMap["size"] = ippoolPool.Size
    ippoolPoolMap["supportsDHCP"] = ippoolPool.Supports_dhcp
    ippoolPoolMap["v4Assigned"] = ippoolPool.V4_assigned
    ippoolPoolMap["v4Size"] = ippoolPool.V4_size
    ippoolPoolMap["v6Assigned"] = ippoolPool.V6_assigned
    ippoolPoolMap["v6Size"] = ippoolPool.V6_size
    
	return ippoolPoolMap, nil

}

func IppoolPoolFromDoc(doc *etree.Document, rootClass string) *IppoolPool {
	element, err := GetMoElement(doc, rootClass, IppoolPoolClassName)
	if err != nil {
		return nil
	}
	return &IppoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         IppoolPoolClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		IppoolPoolAttributes: IppoolPoolAttributes{

        
			Assigned:  element.SelectAttrValue("assigned", ""),
			Assignment_order:  element.SelectAttrValue("assignmentOrder", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Ext_managed:  element.SelectAttrValue("extManaged", ""),
			Guid:  element.SelectAttrValue("guid", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Is_net_bios_enabled:  element.SelectAttrValue("isNetBIOSEnabled", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Size:  element.SelectAttrValue("size", ""),
			Supports_dhcp:  element.SelectAttrValue("supportsDHCP", ""),
			V4_assigned:  element.SelectAttrValue("v4Assigned", ""),
			V4_size:  element.SelectAttrValue("v4Size", ""),
			V6_assigned:  element.SelectAttrValue("v6Assigned", ""),
			V6_size:  element.SelectAttrValue("v6Size", ""),
		},
    
	}
}