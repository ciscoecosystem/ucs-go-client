package models

import (
	"fmt"

	"github.com/beevik/etree"

)

const LstorageDiskGroupConfigPolicyClassName = "lstorageDiskGroupConfigPolicy"

type LstorageDiskGroupConfigPolicy struct {
    BaseAttributes
    LstorageDiskGroupConfigPolicyAttributes
}

type LstorageDiskGroupConfigPolicyAttributes struct{
    Child_action string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Raid_level string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewLstorageDiskGroupConfigPolicy(lstorageDiskGroupConfigPolicyRn,parentDn,description string, lstorageDiskGroupConfigPolicyAttributes LstorageDiskGroupConfigPolicyAttributes) *LstorageDiskGroupConfigPolicy { 
    dn := fmt.Sprintf("%s/%s", parentDn, lstorageDiskGroupConfigPolicyRn)
	return &LstorageDiskGroupConfigPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LstorageDiskGroupConfigPolicyClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LstorageDiskGroupConfigPolicyAttributes: lstorageDiskGroupConfigPolicyAttributes,
	}
}

func (lstorageDiskGroupConfigPolicy *LstorageDiskGroupConfigPolicy) ToMap() (map[string]string, error) {
	lstorageDiskGroupConfigPolicyMap, err := lstorageDiskGroupConfigPolicy.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lstorageDiskGroupConfigPolicyMap["childAction"] = lstorageDiskGroupConfigPolicy.Child_action
    lstorageDiskGroupConfigPolicyMap["intId"] = lstorageDiskGroupConfigPolicy.Int_id
    lstorageDiskGroupConfigPolicyMap["policyLevel"] = lstorageDiskGroupConfigPolicy.Policy_level
    lstorageDiskGroupConfigPolicyMap["policyOwner"] = lstorageDiskGroupConfigPolicy.Policy_owner
    lstorageDiskGroupConfigPolicyMap["raidLevel"] = lstorageDiskGroupConfigPolicy.Raid_level
    lstorageDiskGroupConfigPolicyMap["sacl"] = lstorageDiskGroupConfigPolicy.Sacl
    
	return lstorageDiskGroupConfigPolicyMap, nil

}

func LstorageDiskGroupConfigPolicyFromDoc(doc *etree.Document, rootClass string) *LstorageDiskGroupConfigPolicy {
	element, err := GetMoElement(doc, rootClass, LstorageDiskGroupConfigPolicyClassName)
	if err != nil {
		return nil
	}
	return &LstorageDiskGroupConfigPolicy{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LstorageDiskGroupConfigPolicyClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LstorageDiskGroupConfigPolicyAttributes: LstorageDiskGroupConfigPolicyAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Raid_level:  element.SelectAttrValue("raidLevel", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}