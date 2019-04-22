package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const LsPowerClassName = "lsPower"

type LsPower struct {
    BaseAttributes
    LsPowerAttributes
}

type LsPowerAttributes struct{
    Child_action string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Soft_shutdown_timer string `xml:",omitempty"`
    State string `xml:",omitempty"`
    

}

func NewLsPower(lsPowerRn,parentDn,description string, lsPowerAttributes LsPowerAttributes) *LsPower { 
    dn := fmt.Sprintf("%s/%s", parentDn, lsPowerRn)
	return &LsPower{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LsPowerClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LsPowerAttributes: lsPowerAttributes,
	}
}

func (lsPower *LsPower) ToMap() (map[string]string, error) {
	lsPowerMap, err := lsPower.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lsPowerMap["childAction"] = lsPower.Child_action
    lsPowerMap["propAcl"] = lsPower.Prop_acl
    lsPowerMap["sacl"] = lsPower.Sacl
    lsPowerMap["softShutdownTimer"] = lsPower.Soft_shutdown_timer
    lsPowerMap["state"] = lsPower.State
    
	return lsPowerMap, nil

}

func LsPowerFromDoc(doc *etree.Document, rootClass string) *LsPower {
	element, err := GetMoElement(doc, rootClass, LsPowerClassName)
	if err != nil {
		return nil
	}
	return &LsPower{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LsPowerClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LsPowerAttributes: LsPowerAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Soft_shutdown_timer:  element.SelectAttrValue("softShutdownTimer", ""),
			State:  element.SelectAttrValue("state", ""),
		},
    
	}
}