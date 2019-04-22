package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const MgmtVnetClassName = "mgmtVnet"

type MgmtVnet struct {
    BaseAttributes
    MgmtVnetAttributes
}

type MgmtVnetAttributes struct{
    Child_action string `xml:",omitempty"`
    Config_state string `xml:",omitempty"`
    Mgmt_vnet_id string `xml:",omitempty"`
    Name string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewMgmtVnet(mgmtVnetRn,parentDn,description string, mgmtVnetAttributes MgmtVnetAttributes) *MgmtVnet { 
    dn := fmt.Sprintf("%s/%s", parentDn, mgmtVnetRn)
	return &MgmtVnet{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         MgmtVnetClassName,
			Status:            "created, modified",
			Description:       description,
		},

		MgmtVnetAttributes: mgmtVnetAttributes,
	}
}

func (mgmtVnet *MgmtVnet) ToMap() (map[string]string, error) {
	mgmtVnetMap, err := mgmtVnet.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    mgmtVnetMap["childAction"] = mgmtVnet.Child_action
    mgmtVnetMap["configState"] = mgmtVnet.Config_state
    mgmtVnetMap["id"] = mgmtVnet.Mgmt_vnet_id
    mgmtVnetMap["name"] = mgmtVnet.Name
    mgmtVnetMap["sacl"] = mgmtVnet.Sacl
    
	return mgmtVnetMap, nil

}

func MgmtVnetFromDoc(doc *etree.Document, rootClass string) *MgmtVnet {
	element, err := GetMoElement(doc, rootClass, MgmtVnetClassName)
	if err != nil {
		return nil
	}
	return &MgmtVnet{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         MgmtVnetClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		MgmtVnetAttributes: MgmtVnetAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_state:  element.SelectAttrValue("configState", ""),
			Mgmt_vnet_id:  element.SelectAttrValue("id", ""),
			Name:  element.SelectAttrValue("name", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}