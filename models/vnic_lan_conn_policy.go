package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicLanConnPolicyClassName = "vnicLanConnPolicy"

type VnicLanConnPolicy struct {
    BaseAttributes
    VnicLanConnPolicyAttributes
}

type VnicLanConnPolicyAttributes struct{
    Child_action string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewVnicLanConnPolicy(vnicLanConnPolicyRn,parentDn,description string, vnicLanConnPolicyAttributes VnicLanConnPolicyAttributes) *VnicLanConnPolicy { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicLanConnPolicyRn)
	return &VnicLanConnPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicLanConnPolicyClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicLanConnPolicyAttributes: vnicLanConnPolicyAttributes,
	}
}

func (vnicLanConnPolicy *VnicLanConnPolicy) ToMap() (map[string]string, error) {
	vnicLanConnPolicyMap, err := vnicLanConnPolicy.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicLanConnPolicyMap["childAction"] = vnicLanConnPolicy.Child_action
    vnicLanConnPolicyMap["fltAggr"] = vnicLanConnPolicy.Flt_aggr
    vnicLanConnPolicyMap["intId"] = vnicLanConnPolicy.Int_id
    vnicLanConnPolicyMap["policyLevel"] = vnicLanConnPolicy.Policy_level
    vnicLanConnPolicyMap["policyOwner"] = vnicLanConnPolicy.Policy_owner
    vnicLanConnPolicyMap["sacl"] = vnicLanConnPolicy.Sacl
    
	return vnicLanConnPolicyMap, nil

}

func VnicLanConnPolicyFromDoc(doc *etree.Document, rootClass string) *VnicLanConnPolicy {
	element, err := GetMoElement(doc, rootClass, VnicLanConnPolicyClassName)
	if err != nil {
		return nil
	}
	return &VnicLanConnPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicLanConnPolicyClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicLanConnPolicyAttributes: VnicLanConnPolicyAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}