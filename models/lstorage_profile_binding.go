package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const LstorageProfileBindingClassName = "lstorageProfileBinding"

type LstorageProfileBinding struct {
    BaseAttributes
    LstorageProfileBindingAttributes
}

type LstorageProfileBindingAttributes struct{
    Assigned_to_dn string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Issues string `xml:",omitempty"`
    Name string `xml:",omitempty"`
    Oper_storage_profile_name string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Storage_profile_name string `xml:",omitempty"`
    

}

func NewLstorageProfileBinding(lstorageProfileBindingRn,parentDn,description string, lstorageProfileBindingAttributes LstorageProfileBindingAttributes) *LstorageProfileBinding { 
    dn := fmt.Sprintf("%s/%s", parentDn, lstorageProfileBindingRn)
	return &LstorageProfileBinding{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LstorageProfileBindingClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LstorageProfileBindingAttributes: lstorageProfileBindingAttributes,
	}
}

func (lstorageProfileBinding *LstorageProfileBinding) ToMap() (map[string]string, error) {
	lstorageProfileBindingMap, err := lstorageProfileBinding.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lstorageProfileBindingMap["assignedToDn"] = lstorageProfileBinding.Assigned_to_dn
    lstorageProfileBindingMap["childAction"] = lstorageProfileBinding.Child_action
    lstorageProfileBindingMap["issues"] = lstorageProfileBinding.Issues
    lstorageProfileBindingMap["name"] = lstorageProfileBinding.Name
    lstorageProfileBindingMap["operStorageProfileName"] = lstorageProfileBinding.Oper_storage_profile_name
    lstorageProfileBindingMap["sacl"] = lstorageProfileBinding.Sacl
    lstorageProfileBindingMap["storageProfileName"] = lstorageProfileBinding.Storage_profile_name
    
	return lstorageProfileBindingMap, nil

}

func LstorageProfileBindingFromDoc(doc *etree.Document, rootClass string) *LstorageProfileBinding {
	element, err := GetMoElement(doc, rootClass, LstorageProfileBindingClassName)
	if err != nil {
		return nil
	}
	return &LstorageProfileBinding{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LstorageProfileBindingClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LstorageProfileBindingAttributes: LstorageProfileBindingAttributes{

        
			Assigned_to_dn:  element.SelectAttrValue("assignedToDn", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Issues:  element.SelectAttrValue("issues", ""),
			Name:  element.SelectAttrValue("name", ""),
			Oper_storage_profile_name:  element.SelectAttrValue("operStorageProfileName", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Storage_profile_name:  element.SelectAttrValue("storageProfileName", ""),
		},
    
	}
}