package models

import (
	"fmt"

	"github.com/beevik/etree"

)

const MacpoolBlockClassName = "macpoolBlock"

type MacpoolBlock struct {
    BaseAttributes
    MacpoolBlockAttributes
}

type MacpoolBlockAttributes struct{
    Child_action string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewMacpoolBlock(macpoolBlockRn,parentDn,description string, macpoolBlockAttributes MacpoolBlockAttributes) *MacpoolBlock { 
    dn := fmt.Sprintf("%s/%s", parentDn, macpoolBlockRn)
	return &MacpoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         MacpoolBlockClassName,
			Status:            "created, modified",
			Description:       description,
		},

		MacpoolBlockAttributes: macpoolBlockAttributes,
	}
}

func (macpoolBlock *MacpoolBlock) ToMap() (map[string]string, error) {
	macpoolBlockMap, err := macpoolBlock.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    macpoolBlockMap["childAction"] = macpoolBlock.Child_action
    macpoolBlockMap["sacl"] = macpoolBlock.Sacl
    
	return macpoolBlockMap, nil

}

func MacpoolBlockFromDoc(doc *etree.Document, rootClass string) *MacpoolBlock {
	element, err := GetMoElement(doc, rootClass, MacpoolBlockClassName)
	if err != nil {
		return nil
	}
	return &MacpoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         MacpoolBlockClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		MacpoolBlockAttributes: MacpoolBlockAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}