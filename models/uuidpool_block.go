package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const UuidpoolBlockClassName = "uuidpoolBlock"

type UuidpoolBlock struct {
    BaseAttributes
    UuidpoolBlockAttributes
}

type UuidpoolBlockAttributes struct{
    Child_action string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewUuidpoolBlock(uuidpoolBlockRn,parentDn,description string, uuidpoolBlockAttributes UuidpoolBlockAttributes) *UuidpoolBlock { 
    dn := fmt.Sprintf("%s/%s", parentDn, uuidpoolBlockRn)
	return &UuidpoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         UuidpoolBlockClassName,
			Status:            "created, modified",
			Description:       description,
		},

		UuidpoolBlockAttributes: uuidpoolBlockAttributes,
	}
}

func (uuidpoolBlock *UuidpoolBlock) ToMap() (map[string]string, error) {
	uuidpoolBlockMap, err := uuidpoolBlock.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    uuidpoolBlockMap["childAction"] = uuidpoolBlock.Child_action
    uuidpoolBlockMap["sacl"] = uuidpoolBlock.Sacl
    
	return uuidpoolBlockMap, nil

}

func UuidpoolBlockFromDoc(doc *etree.Document, rootClass string) *UuidpoolBlock {
	element, err := GetMoElement(doc, rootClass, UuidpoolBlockClassName)
	if err != nil {
		return nil
	}
	return &UuidpoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         UuidpoolBlockClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		UuidpoolBlockAttributes: UuidpoolBlockAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}