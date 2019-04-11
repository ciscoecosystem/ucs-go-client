package models

import (
	"fmt"

	"github.com/beevik/etree"

)

const MacpoolPoolClassName = "macpoolPool"

type MacpoolPool struct {
    BaseAttributes
    MacpoolPoolAttributes
}

type MacpoolPoolAttributes struct{
    Assignment_order string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Name string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    

}

func NewMacpoolPool(macpoolPoolRn, parentDn, description string, macpoolPoolAttributes MacpoolPoolAttributes) *MacpoolPool {
	dn := fmt.Sprintf("%s/%s", parentDn, macpoolPoolRn)
	return &MacpoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         MacpoolPoolClassName,
			Status:            "created, modified",
			Description:       description,
		},

		MacpoolPoolAttributes: macpoolPoolAttributes,
	}
}

func (macpoolPool *MacpoolPool) ToMap() (map[string]string, error) {
	macpoolPoolMap, err := macpoolPool.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    macpoolPoolMap["assignmentOrder"] = macpoolPool.Assignment_order
    macpoolPoolMap["childAction"] = macpoolPool.Child_action
    macpoolPoolMap["intId"] = macpoolPool.Int_id
    macpoolPoolMap["name"] = macpoolPool.Name
    macpoolPoolMap["policyOwner"] = macpoolPool.Policy_owner
    
	return macpoolPoolMap, nil

}

func MacpoolPoolFromDoc(doc *etree.Document, rootClass string) *MacpoolPool {
	element, err := GetMoElement(doc, rootClass, MacpoolPoolClassName)
	if err != nil {
		return nil
	}
	return &MacpoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         MacpoolPoolClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		MacpoolPoolAttributes: MacpoolPoolAttributes{

        
			Assignment_order:  element.SelectAttrValue("assignmentOrder", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Name:  element.SelectAttrValue("name", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
		},
    
	}
}