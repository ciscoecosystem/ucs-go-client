package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const LsRequirementClassName = "lsRequirement"

type LsRequirement struct {
    BaseAttributes
    LsRequirementAttributes
}

type LsRequirementAttributes struct{
    Admin_action string `xml:",omitempty"`
    Admin_action_trigger string `xml:",omitempty"`
    Assigned_to_dn string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Compute_ep_dn string `xml:",omitempty"`
    Issues string `xml:",omitempty"`
    Name string `xml:",omitempty"`
    Oper_name string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Pn_dn string `xml:",omitempty"`
    Pn_pool_dn string `xml:",omitempty"`
    Qualifier string `xml:",omitempty"`
    Restrict_migration string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewLsRequirement(lsRequirementRn,parentDn,description string, lsRequirementAttributes LsRequirementAttributes) *LsRequirement { 
    dn := fmt.Sprintf("%s/%s", parentDn, lsRequirementRn)
	return &LsRequirement{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LsRequirementClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LsRequirementAttributes: lsRequirementAttributes,
	}
}

func (lsRequirement *LsRequirement) ToMap() (map[string]string, error) {
	lsRequirementMap, err := lsRequirement.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lsRequirementMap["adminAction"] = lsRequirement.Admin_action
    lsRequirementMap["adminActionTrigger"] = lsRequirement.Admin_action_trigger
    lsRequirementMap["assignedToDn"] = lsRequirement.Assigned_to_dn
    lsRequirementMap["childAction"] = lsRequirement.Child_action
    lsRequirementMap["computeEpDn"] = lsRequirement.Compute_ep_dn
    lsRequirementMap["issues"] = lsRequirement.Issues
    lsRequirementMap["name"] = lsRequirement.Name
    lsRequirementMap["operName"] = lsRequirement.Oper_name
    lsRequirementMap["operState"] = lsRequirement.Oper_state
    lsRequirementMap["pnDn"] = lsRequirement.Pn_dn
    lsRequirementMap["pnPoolDn"] = lsRequirement.Pn_pool_dn
    lsRequirementMap["qualifier"] = lsRequirement.Qualifier
    lsRequirementMap["restrictMigration"] = lsRequirement.Restrict_migration
    lsRequirementMap["sacl"] = lsRequirement.Sacl
    
	return lsRequirementMap, nil

}

func LsRequirementFromDoc(doc *etree.Document, rootClass string) *LsRequirement {
	element, err := GetMoElement(doc, rootClass, LsRequirementClassName)
	if err != nil {
		return nil
	}
	return &LsRequirement{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LsRequirementClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LsRequirementAttributes: LsRequirementAttributes{

        
			Admin_action:  element.SelectAttrValue("adminAction", ""),
			Admin_action_trigger:  element.SelectAttrValue("adminActionTrigger", ""),
			Assigned_to_dn:  element.SelectAttrValue("assignedToDn", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Compute_ep_dn:  element.SelectAttrValue("computeEpDn", ""),
			Issues:  element.SelectAttrValue("issues", ""),
			Name:  element.SelectAttrValue("name", ""),
			Oper_name:  element.SelectAttrValue("operName", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Pn_dn:  element.SelectAttrValue("pnDn", ""),
			Pn_pool_dn:  element.SelectAttrValue("pnPoolDn", ""),
			Qualifier:  element.SelectAttrValue("qualifier", ""),
			Restrict_migration:  element.SelectAttrValue("restrictMigration", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}