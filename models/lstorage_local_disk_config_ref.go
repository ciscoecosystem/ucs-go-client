package models

import (
	"fmt"

	"github.com/beevik/etree"

)

const LstorageLocalDiskConfigRefClassName = "lstorageLocalDiskConfigRef"

type LstorageLocalDiskConfigRef struct {
    BaseAttributes
    LstorageLocalDiskConfigRefAttributes
}

type LstorageLocalDiskConfigRefAttributes struct{
    Child_action string `xml:",omitempty"`
    Role string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Span_id string `xml:",omitempty"`
    

}

func NewLstorageLocalDiskConfigRef(lstorageLocalDiskConfigRefRn,parentDn,description string, lstorageLocalDiskConfigRefAttributes LstorageLocalDiskConfigRefAttributes) *LstorageLocalDiskConfigRef { 
    dn := fmt.Sprintf("%s/%s", parentDn, lstorageLocalDiskConfigRefRn)
	return &LstorageLocalDiskConfigRef{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LstorageLocalDiskConfigRefClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LstorageLocalDiskConfigRefAttributes: lstorageLocalDiskConfigRefAttributes,
	}
}

func (lstorageLocalDiskConfigRef *LstorageLocalDiskConfigRef) ToMap() (map[string]string, error) {
	lstorageLocalDiskConfigRefMap, err := lstorageLocalDiskConfigRef.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lstorageLocalDiskConfigRefMap["childAction"] = lstorageLocalDiskConfigRef.Child_action
    lstorageLocalDiskConfigRefMap["role"] = lstorageLocalDiskConfigRef.Role
    lstorageLocalDiskConfigRefMap["sacl"] = lstorageLocalDiskConfigRef.Sacl
    lstorageLocalDiskConfigRefMap["spanId"] = lstorageLocalDiskConfigRef.Span_id
    
	return lstorageLocalDiskConfigRefMap, nil

}

func LstorageLocalDiskConfigRefFromDoc(doc *etree.Document, rootClass string) *LstorageLocalDiskConfigRef {
	element, err := GetMoElement(doc, rootClass, LstorageLocalDiskConfigRefClassName)
	if err != nil {
		return nil
	}
	return &LstorageLocalDiskConfigRef{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LstorageLocalDiskConfigRefClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LstorageLocalDiskConfigRefAttributes: LstorageLocalDiskConfigRefAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Role:  element.SelectAttrValue("role", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Span_id:  element.SelectAttrValue("spanId", ""),
		},
    
	}
}