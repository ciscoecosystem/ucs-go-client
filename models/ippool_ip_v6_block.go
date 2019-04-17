package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const IppoolIpV6BlockClassName = "ippoolIpV6Block"

type IppoolIpV6Block struct {
    BaseAttributes
    IppoolIpV6BlockAttributes
}

type IppoolIpV6BlockAttributes struct{
    Child_action string `xml:",omitempty"`
    Def_gw string `xml:",omitempty"`
    Prefix string `xml:",omitempty"`
    Prim_dns string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Sec_dns string `xml:",omitempty"`
    

}

func NewIppoolIpV6Block(ippoolIpV6BlockRn,parentDn,description string, ippoolIpV6BlockAttributes IppoolIpV6BlockAttributes) *IppoolIpV6Block { 
    dn := fmt.Sprintf("%s/%s", parentDn, ippoolIpV6BlockRn)
	return &IppoolIpV6Block{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         IppoolIpV6BlockClassName,
			Status:            "created, modified",
			Description:       description,
		},

		IppoolIpV6BlockAttributes: ippoolIpV6BlockAttributes,
	}
}

func (ippoolIpV6Block *IppoolIpV6Block) ToMap() (map[string]string, error) {
	ippoolIpV6BlockMap, err := ippoolIpV6Block.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    ippoolIpV6BlockMap["childAction"] = ippoolIpV6Block.Child_action
    ippoolIpV6BlockMap["defGw"] = ippoolIpV6Block.Def_gw
    ippoolIpV6BlockMap["prefix"] = ippoolIpV6Block.Prefix
    ippoolIpV6BlockMap["primDns"] = ippoolIpV6Block.Prim_dns
    ippoolIpV6BlockMap["sacl"] = ippoolIpV6Block.Sacl
    ippoolIpV6BlockMap["secDns"] = ippoolIpV6Block.Sec_dns
    
	return ippoolIpV6BlockMap, nil

}

func IppoolIpV6BlockFromDoc(doc *etree.Document, rootClass string) *IppoolIpV6Block {
	element, err := GetMoElement(doc, rootClass, IppoolIpV6BlockClassName)
	if err != nil {
		return nil
	}
	return &IppoolIpV6Block{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         IppoolIpV6BlockClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		IppoolIpV6BlockAttributes: IppoolIpV6BlockAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Def_gw:  element.SelectAttrValue("defGw", ""),
			Prefix:  element.SelectAttrValue("prefix", ""),
			Prim_dns:  element.SelectAttrValue("primDns", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Sec_dns:  element.SelectAttrValue("secDns", ""),
		},
    
	}
}