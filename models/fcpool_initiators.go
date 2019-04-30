package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const FcpoolInitiatorsClassName = "fcpoolInitiators"

type FcpoolInitiators struct {
    BaseAttributes
    FcpoolInitiatorsAttributes
}

type FcpoolInitiatorsAttributes struct{
    Assigned string `xml:",omitempty"`
    Assignment_order string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Max_ports_per_node string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Purpose string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Size string `xml:",omitempty"`
    

}

func NewFcpoolInitiators(fcpoolInitiatorsRn,parentDn,description string, fcpoolInitiatorsAttributes FcpoolInitiatorsAttributes) *FcpoolInitiators { 
    dn := fmt.Sprintf("%s/%s", parentDn, fcpoolInitiatorsRn)
	return &FcpoolInitiators{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         FcpoolInitiatorsClassName,
			Status:            "created, modified",
			Description:       description,
		},

		FcpoolInitiatorsAttributes: fcpoolInitiatorsAttributes,
	}
}

func (fcpoolInitiators *FcpoolInitiators) ToMap() (map[string]string, error) {
	fcpoolInitiatorsMap, err := fcpoolInitiators.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    fcpoolInitiatorsMap["assigned"] = fcpoolInitiators.Assigned
    fcpoolInitiatorsMap["assignmentOrder"] = fcpoolInitiators.Assignment_order
    fcpoolInitiatorsMap["childAction"] = fcpoolInitiators.Child_action
    fcpoolInitiatorsMap["intId"] = fcpoolInitiators.Int_id
    fcpoolInitiatorsMap["maxPortsPerNode"] = fcpoolInitiators.Max_ports_per_node
    fcpoolInitiatorsMap["policyLevel"] = fcpoolInitiators.Policy_level
    fcpoolInitiatorsMap["policyOwner"] = fcpoolInitiators.Policy_owner
    fcpoolInitiatorsMap["purpose"] = fcpoolInitiators.Purpose
    fcpoolInitiatorsMap["sacl"] = fcpoolInitiators.Sacl
    fcpoolInitiatorsMap["size"] = fcpoolInitiators.Size
    
	return fcpoolInitiatorsMap, nil

}

func FcpoolInitiatorsFromDoc(doc *etree.Document, rootClass string) *FcpoolInitiators {
	element, err := GetMoElement(doc, rootClass, FcpoolInitiatorsClassName)
	if err != nil {
		return nil
	}
	return &FcpoolInitiators{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         FcpoolInitiatorsClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		FcpoolInitiatorsAttributes: FcpoolInitiatorsAttributes{

        
			Assigned:  element.SelectAttrValue("assigned", ""),
			Assignment_order:  element.SelectAttrValue("assignmentOrder", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Max_ports_per_node:  element.SelectAttrValue("maxPortsPerNode", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Purpose:  element.SelectAttrValue("purpose", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Size:  element.SelectAttrValue("size", ""),
		},
    
	}
}