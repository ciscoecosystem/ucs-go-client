package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicVlanClassName = "vnicVlan"

type VnicVlan struct {
    BaseAttributes
    VnicVlanAttributes
}

type VnicVlanAttributes struct{
    Child_action string `xml:",omitempty"`
    Config_qualifier string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
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
    Vlan_name string `xml:",omitempty"`
    Vnet string `xml:",omitempty"`
    

}

func NewVnicVlan(vnicVlanRn,parentDn,description string, vnicVlanAttributes VnicVlanAttributes) *VnicVlan { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicVlanRn)
	return &VnicVlan{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicVlanClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicVlanAttributes: vnicVlanAttributes,
	}
}

func (vnicVlan *VnicVlan) ToMap() (map[string]string, error) {
	vnicVlanMap, err := vnicVlan.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicVlanMap["childAction"] = vnicVlan.Child_action
    vnicVlanMap["configQualifier"] = vnicVlan.Config_qualifier
    vnicVlanMap["fltAggr"] = vnicVlan.Flt_aggr
    vnicVlanMap["name"] = vnicVlan.Name
    vnicVlanMap["operPrimaryVnetDn"] = vnicVlan.Oper_primary_vnet_dn
    vnicVlanMap["operPrimaryVnetName"] = vnicVlan.Oper_primary_vnet_name
    vnicVlanMap["operState"] = vnicVlan.Oper_state
    vnicVlanMap["operVnetDn"] = vnicVlan.Oper_vnet_dn
    vnicVlanMap["operVnetName"] = vnicVlan.Oper_vnet_name
    vnicVlanMap["owner"] = vnicVlan.Owner
    vnicVlanMap["pubNwId"] = vnicVlan.Pub_nw_id
    vnicVlanMap["sacl"] = vnicVlan.Sacl
    vnicVlanMap["sharing"] = vnicVlan.Sharing
    vnicVlanMap["switchId"] = vnicVlan.Switch_id
    vnicVlanMap["type"] = vnicVlan.Type
    vnicVlanMap["vlanName"] = vnicVlan.Vlan_name
    vnicVlanMap["vnet"] = vnicVlan.Vnet
    
	return vnicVlanMap, nil

}

func VnicVlanFromDoc(doc *etree.Document, rootClass string) *VnicVlan {
	element, err := GetMoElement(doc, rootClass, VnicVlanClassName)
	if err != nil {
		return nil
	}
	return &VnicVlan{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicVlanClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicVlanAttributes: VnicVlanAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_qualifier:  element.SelectAttrValue("configQualifier", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
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
			Vlan_name:  element.SelectAttrValue("vlanName", ""),
			Vnet:  element.SelectAttrValue("vnet", ""),
		},
    
	}
}