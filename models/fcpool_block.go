package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const FcpoolBlockClassName = "fcpoolBlock"

type FcpoolBlock struct {
    BaseAttributes
    FcpoolBlockAttributes
}

type FcpoolBlockAttributes struct{
    Child_action string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewFcpoolBlock(fcpoolBlockRn,parentDn,description string, fcpoolBlockAttributes FcpoolBlockAttributes) *FcpoolBlock { 
    dn := fmt.Sprintf("%s/%s", parentDn, fcpoolBlockRn)
	return &FcpoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         FcpoolBlockClassName,
			Status:            "created, modified",
			Description:       description,
		},

		FcpoolBlockAttributes: fcpoolBlockAttributes,
	}
}

func (fcpoolBlock *FcpoolBlock) ToMap() (map[string]string, error) {
	fcpoolBlockMap, err := fcpoolBlock.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    fcpoolBlockMap["childAction"] = fcpoolBlock.Child_action
    fcpoolBlockMap["sacl"] = fcpoolBlock.Sacl
    
	return fcpoolBlockMap, nil

}

func FcpoolBlockFromDoc(doc *etree.Document, rootClass string) *FcpoolBlock {
	element, err := GetMoElement(doc, rootClass, FcpoolBlockClassName)
	if err != nil {
		return nil
	}
	return &FcpoolBlock{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         FcpoolBlockClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		FcpoolBlockAttributes: FcpoolBlockAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}