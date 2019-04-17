package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const OrgOrgClassName = "orgOrg"

type OrgOrg struct {
    BaseAttributes
    OrgOrgAttributes
}

type OrgOrgAttributes struct{
    Child_action string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Level string `xml:",omitempty"`
    Perm_access string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewOrgOrg(orgOrgRn,parentDn,description string, orgOrgAttributes OrgOrgAttributes) *OrgOrg { 
    dn := fmt.Sprintf("%s/%s", parentDn, orgOrgRn)
	return &OrgOrg{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         OrgOrgClassName,
			Status:            "created, modified",
			Description:       description,
		},

		OrgOrgAttributes: orgOrgAttributes,
	}
}

func (orgOrg *OrgOrg) ToMap() (map[string]string, error) {
	orgOrgMap, err := orgOrg.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    orgOrgMap["childAction"] = orgOrg.Child_action
    orgOrgMap["fltAggr"] = orgOrg.Flt_aggr
    orgOrgMap["level"] = orgOrg.Level
    orgOrgMap["permAccess"] = orgOrg.Perm_access
    orgOrgMap["sacl"] = orgOrg.Sacl
    
	return orgOrgMap, nil

}

func OrgOrgFromDoc(doc *etree.Document, rootClass string) *OrgOrg {
	element, err := GetMoElement(doc, rootClass, OrgOrgClassName)
	if err != nil {
		return nil
	}
	return &OrgOrg{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         OrgOrgClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		OrgOrgAttributes: OrgOrgAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Level:  element.SelectAttrValue("level", ""),
			Perm_access:  element.SelectAttrValue("permAccess", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}