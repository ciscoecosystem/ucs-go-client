package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const MgmtInterfaceClassName = "mgmtInterface"

type MgmtInterface struct {
    BaseAttributes
    MgmtInterfaceAttributes
}

type MgmtInterfaceAttributes struct{
    Child_action string `xml:",omitempty"`
    Config_message string `xml:",omitempty"`
    Config_state string `xml:",omitempty"`
    Ip_v4_state string `xml:",omitempty"`
    Ip_v6_state string `xml:",omitempty"`
    Is_default_derived string `xml:",omitempty"`
    Monitor_interval string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewMgmtInterface(mgmtInterfaceRn,parentDn,description string, mgmtInterfaceAttributes MgmtInterfaceAttributes) *MgmtInterface { 
    dn := fmt.Sprintf("%s/%s", parentDn, mgmtInterfaceRn)
	return &MgmtInterface{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         MgmtInterfaceClassName,
			Status:            "created, modified",
			Description:       description,
		},

		MgmtInterfaceAttributes: mgmtInterfaceAttributes,
	}
}

func (mgmtInterface *MgmtInterface) ToMap() (map[string]string, error) {
	mgmtInterfaceMap, err := mgmtInterface.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    mgmtInterfaceMap["childAction"] = mgmtInterface.Child_action
    mgmtInterfaceMap["configMessage"] = mgmtInterface.Config_message
    mgmtInterfaceMap["configState"] = mgmtInterface.Config_state
    mgmtInterfaceMap["ipV4State"] = mgmtInterface.Ip_v4_state
    mgmtInterfaceMap["ipV6State"] = mgmtInterface.Ip_v6_state
    mgmtInterfaceMap["isDefaultDerived"] = mgmtInterface.Is_default_derived
    mgmtInterfaceMap["monitorInterval"] = mgmtInterface.Monitor_interval
    mgmtInterfaceMap["operState"] = mgmtInterface.Oper_state
    mgmtInterfaceMap["sacl"] = mgmtInterface.Sacl
    
	return mgmtInterfaceMap, nil

}

func MgmtInterfaceFromDoc(doc *etree.Document, rootClass string) *MgmtInterface {
	element, err := GetMoElement(doc, rootClass, MgmtInterfaceClassName)
	if err != nil {
		return nil
	}
	return &MgmtInterface{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         MgmtInterfaceClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		MgmtInterfaceAttributes: MgmtInterfaceAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_message:  element.SelectAttrValue("configMessage", ""),
			Config_state:  element.SelectAttrValue("configState", ""),
			Ip_v4_state:  element.SelectAttrValue("ipV4State", ""),
			Ip_v6_state:  element.SelectAttrValue("ipV6State", ""),
			Is_default_derived:  element.SelectAttrValue("isDefaultDerived", ""),
			Monitor_interval:  element.SelectAttrValue("monitorInterval", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}