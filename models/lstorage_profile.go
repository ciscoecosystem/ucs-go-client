package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const LstorageProfileClassName = "lstorageProfile"

type LstorageProfile struct {
    BaseAttributes
    LstorageProfileAttributes
}

type LstorageProfileAttributes struct{
    Assigned_to_dn string `xml:",omitempty"`
    Availability string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewLstorageProfile(lstorageProfileRn,parentDn,description string, lstorageProfileAttributes LstorageProfileAttributes) *LstorageProfile { 
    dn := fmt.Sprintf("%s/%s", parentDn, lstorageProfileRn)
	return &LstorageProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LstorageProfileClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LstorageProfileAttributes: lstorageProfileAttributes,
	}
}

func (lstorageProfile *LstorageProfile) ToMap() (map[string]string, error) {
	lstorageProfileMap, err := lstorageProfile.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lstorageProfileMap["assignedToDn"] = lstorageProfile.Assigned_to_dn
    lstorageProfileMap["availability"] = lstorageProfile.Availability
    lstorageProfileMap["childAction"] = lstorageProfile.Child_action
    lstorageProfileMap["intId"] = lstorageProfile.Int_id
    lstorageProfileMap["policyLevel"] = lstorageProfile.Policy_level
    lstorageProfileMap["policyOwner"] = lstorageProfile.Policy_owner
    lstorageProfileMap["sacl"] = lstorageProfile.Sacl
    
	return lstorageProfileMap, nil

}

func LstorageProfileFromDoc(doc *etree.Document, rootClass string) *LstorageProfile {
	element, err := GetMoElement(doc, rootClass, LstorageProfileClassName)
	if err != nil {
		return nil
	}
	return &LstorageProfile{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LstorageProfileClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LstorageProfileAttributes: LstorageProfileAttributes{

        
			Assigned_to_dn:  element.SelectAttrValue("assignedToDn", ""),
			Availability:  element.SelectAttrValue("availability", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}