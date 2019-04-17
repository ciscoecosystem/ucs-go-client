package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicSanConnPolicyClassName = "vnicSanConnPolicy"

type VnicSanConnPolicy struct {
    BaseAttributes
    VnicSanConnPolicyAttributes
}

type VnicSanConnPolicyAttributes struct{
    Child_action string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewVnicSanConnPolicy(vnicSanConnPolicyRn,parentDn,description string, vnicSanConnPolicyAttributes VnicSanConnPolicyAttributes) *VnicSanConnPolicy { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicSanConnPolicyRn)
	return &VnicSanConnPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicSanConnPolicyClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicSanConnPolicyAttributes: vnicSanConnPolicyAttributes,
	}
}

func (vnicSanConnPolicy *VnicSanConnPolicy) ToMap() (map[string]string, error) {
	vnicSanConnPolicyMap, err := vnicSanConnPolicy.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicSanConnPolicyMap["childAction"] = vnicSanConnPolicy.Child_action
    vnicSanConnPolicyMap["fltAggr"] = vnicSanConnPolicy.Flt_aggr
    vnicSanConnPolicyMap["intId"] = vnicSanConnPolicy.Int_id
    vnicSanConnPolicyMap["policyLevel"] = vnicSanConnPolicy.Policy_level
    vnicSanConnPolicyMap["policyOwner"] = vnicSanConnPolicy.Policy_owner
    vnicSanConnPolicyMap["sacl"] = vnicSanConnPolicy.Sacl
    
	return vnicSanConnPolicyMap, nil

}

func VnicSanConnPolicyFromDoc(doc *etree.Document, rootClass string) *VnicSanConnPolicy {
	element, err := GetMoElement(doc, rootClass, VnicSanConnPolicyClassName)
	if err != nil {
		return nil
	}
	return &VnicSanConnPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicSanConnPolicyClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicSanConnPolicyAttributes: VnicSanConnPolicyAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}