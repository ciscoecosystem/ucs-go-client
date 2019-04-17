package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const IppoolBlockClassName = "ippoolBlock"

type IppoolBlock struct {
    BaseAttributes
    IppoolBlockAttributes
}

type IppoolBlockAttributes struct{
    Child_action string `xml:",omitempty"`
    Def_gw string `xml:",omitempty"`
    Prim_dns string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Sec_dns string `xml:",omitempty"`
    Subnet string `xml:",omitempty"`
    

}

func NewIppoolBlock(ippoolBlockRn,parentDn,description string, ippoolBlockAttributes IppoolBlockAttributes) *IppoolBlock { 
    dn := fmt.Sprintf("%s/%s", parentDn, ippoolBlockRn)
	return &IppoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         IppoolBlockClassName,
			Status:            "created, modified",
			Description:       description,
		},

		IppoolBlockAttributes: ippoolBlockAttributes,
	}
}

func (ippoolBlock *IppoolBlock) ToMap() (map[string]string, error) {
	ippoolBlockMap, err := ippoolBlock.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    ippoolBlockMap["childAction"] = ippoolBlock.Child_action
    ippoolBlockMap["defGw"] = ippoolBlock.Def_gw
    ippoolBlockMap["primDns"] = ippoolBlock.Prim_dns
    ippoolBlockMap["sacl"] = ippoolBlock.Sacl
    ippoolBlockMap["secDns"] = ippoolBlock.Sec_dns
    ippoolBlockMap["subnet"] = ippoolBlock.Subnet
    
	return ippoolBlockMap, nil

}

func IppoolBlockFromDoc(doc *etree.Document, rootClass string) *IppoolBlock {
	element, err := GetMoElement(doc, rootClass, IppoolBlockClassName)
	if err != nil {
		return nil
	}
	return &IppoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         IppoolBlockClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		IppoolBlockAttributes: IppoolBlockAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Def_gw:  element.SelectAttrValue("defGw", ""),
			Prim_dns:  element.SelectAttrValue("primDns", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Sec_dns:  element.SelectAttrValue("secDns", ""),
			Subnet:  element.SelectAttrValue("subnet", ""),
		},
    
	}
}