package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicFcIfClassName = "vnicFcIf"

type VnicFcIf struct {
    BaseAttributes
    VnicFcIfAttributes
}

type VnicFcIfAttributes struct{
    Child_action string `xml:",omitempty"`
    Config_qualifier string `xml:",omitempty"`
    Initiator string `xml:",omitempty"`
    Name string `xml:",omitempty"`
    Oper_primary_vnet_dn string `xml:",omitempty"`
    Oper_primary_vnet_name string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Oper_vnet_dn string `xml:",omitempty"`
    Oper_vnet_name string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Pub_nw_id string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Sharing string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    Vnet string `xml:",omitempty"`
    

}

func NewVnicFcIf(vnicFcIfRn,parentDn,description string, vnicFcIfAttributes VnicFcIfAttributes) *VnicFcIf { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicFcIfRn)
	return &VnicFcIf{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicFcIfClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicFcIfAttributes: vnicFcIfAttributes,
	}
}

func (vnicFcIf *VnicFcIf) ToMap() (map[string]string, error) {
	vnicFcIfMap, err := vnicFcIf.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicFcIfMap["childAction"] = vnicFcIf.Child_action
    vnicFcIfMap["configQualifier"] = vnicFcIf.Config_qualifier
    vnicFcIfMap["initiator"] = vnicFcIf.Initiator
    vnicFcIfMap["name"] = vnicFcIf.Name
    vnicFcIfMap["operPrimaryVnetDn"] = vnicFcIf.Oper_primary_vnet_dn
    vnicFcIfMap["operPrimaryVnetName"] = vnicFcIf.Oper_primary_vnet_name
    vnicFcIfMap["operState"] = vnicFcIf.Oper_state
    vnicFcIfMap["operVnetDn"] = vnicFcIf.Oper_vnet_dn
    vnicFcIfMap["operVnetName"] = vnicFcIf.Oper_vnet_name
    vnicFcIfMap["owner"] = vnicFcIf.Owner
    vnicFcIfMap["pubNwId"] = vnicFcIf.Pub_nw_id
    vnicFcIfMap["sacl"] = vnicFcIf.Sacl
    vnicFcIfMap["sharing"] = vnicFcIf.Sharing
    vnicFcIfMap["switchId"] = vnicFcIf.Switch_id
    vnicFcIfMap["type"] = vnicFcIf.Type
    vnicFcIfMap["vnet"] = vnicFcIf.Vnet
    
	return vnicFcIfMap, nil

}

func VnicFcIfFromDoc(doc *etree.Document, rootClass string) *VnicFcIf {
	element, err := GetMoElement(doc, rootClass, VnicFcIfClassName)
	if err != nil {
		return nil
	}
	return &VnicFcIf{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicFcIfClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicFcIfAttributes: VnicFcIfAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_qualifier:  element.SelectAttrValue("configQualifier", ""),
			Initiator:  element.SelectAttrValue("initiator", ""),
			Name:  element.SelectAttrValue("name", ""),
			Oper_primary_vnet_dn:  element.SelectAttrValue("operPrimaryVnetDn", ""),
			Oper_primary_vnet_name:  element.SelectAttrValue("operPrimaryVnetName", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Oper_vnet_dn:  element.SelectAttrValue("operVnetDn", ""),
			Oper_vnet_name:  element.SelectAttrValue("operVnetName", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Pub_nw_id:  element.SelectAttrValue("pubNwId", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Sharing:  element.SelectAttrValue("sharing", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Type:  element.SelectAttrValue("type", ""),
			Vnet:  element.SelectAttrValue("vnet", ""),
		},
    
	}
}