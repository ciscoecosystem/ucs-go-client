package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const UuidpoolPoolClassName = "uuidpoolPool"

type UuidpoolPool struct {
    BaseAttributes
    UuidpoolPoolAttributes
}

type UuidpoolPoolAttributes struct{
    Assigned string `xml:",omitempty"`
    Assignment_order string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Prefix string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Size string `xml:",omitempty"`
    

}

func NewUuidpoolPool(uuidpoolPoolRn,parentDn,description string, uuidpoolPoolAttributes UuidpoolPoolAttributes) *UuidpoolPool { 
    dn := fmt.Sprintf("%s/%s", parentDn, uuidpoolPoolRn)
	return &UuidpoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         UuidpoolPoolClassName,
			Status:            "created, modified",
			Description:       description,
		},

		UuidpoolPoolAttributes: uuidpoolPoolAttributes,
	}
}

func (uuidpoolPool *UuidpoolPool) ToMap() (map[string]string, error) {
	uuidpoolPoolMap, err := uuidpoolPool.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    uuidpoolPoolMap["assigned"] = uuidpoolPool.Assigned
    uuidpoolPoolMap["assignmentOrder"] = uuidpoolPool.Assignment_order
    uuidpoolPoolMap["childAction"] = uuidpoolPool.Child_action
    uuidpoolPoolMap["intId"] = uuidpoolPool.Int_id
    uuidpoolPoolMap["policyLevel"] = uuidpoolPool.Policy_level
    uuidpoolPoolMap["policyOwner"] = uuidpoolPool.Policy_owner
    uuidpoolPoolMap["prefix"] = uuidpoolPool.Prefix
    uuidpoolPoolMap["sacl"] = uuidpoolPool.Sacl
    uuidpoolPoolMap["size"] = uuidpoolPool.Size
    
	return uuidpoolPoolMap, nil

}

func UuidpoolPoolFromDoc(doc *etree.Document, rootClass string) *UuidpoolPool {
	element, err := GetMoElement(doc, rootClass, UuidpoolPoolClassName)
	if err != nil {
		return nil
	}
	return &UuidpoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         UuidpoolPoolClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		UuidpoolPoolAttributes: UuidpoolPoolAttributes{

        
			Assigned:  element.SelectAttrValue("assigned", ""),
			Assignment_order:  element.SelectAttrValue("assignmentOrder", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Prefix:  element.SelectAttrValue("prefix", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Size:  element.SelectAttrValue("size", ""),
		},
    
	}
}