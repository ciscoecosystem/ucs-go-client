package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicEtherIfClassName = "vnicEtherIf"

type VnicEtherIf struct {
    BaseAttributes
    VnicEtherIfAttributes
}

type VnicEtherIfAttributes struct{
    Addr string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Config_qualifier string `xml:",omitempty"`
    Default_net string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Oper_primary_vnet_dn string `xml:",omitempty"`
    Oper_primary_vnet_name string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Oper_vnet_dn string `xml:",omitempty"`
    Oper_vnet_name string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Pub_nw_id string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Sharing string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    Vnet string `xml:",omitempty"`
    

}

func NewVnicEtherIf(vnicEtherIfRn,parentDn,description string, vnicEtherIfAttributes VnicEtherIfAttributes) *VnicEtherIf { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicEtherIfRn)
	return &VnicEtherIf{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicEtherIfClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicEtherIfAttributes: vnicEtherIfAttributes,
	}
}

func (vnicEtherIf *VnicEtherIf) ToMap() (map[string]string, error) {
	vnicEtherIfMap, err := vnicEtherIf.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicEtherIfMap["addr"] = vnicEtherIf.Addr
    vnicEtherIfMap["childAction"] = vnicEtherIf.Child_action
    vnicEtherIfMap["configQualifier"] = vnicEtherIf.Config_qualifier
    vnicEtherIfMap["defaultNet"] = vnicEtherIf.Default_net
    vnicEtherIfMap["fltAggr"] = vnicEtherIf.Flt_aggr
    vnicEtherIfMap["operPrimaryVnetDn"] = vnicEtherIf.Oper_primary_vnet_dn
    vnicEtherIfMap["operPrimaryVnetName"] = vnicEtherIf.Oper_primary_vnet_name
    vnicEtherIfMap["operState"] = vnicEtherIf.Oper_state
    vnicEtherIfMap["operVnetDn"] = vnicEtherIf.Oper_vnet_dn
    vnicEtherIfMap["operVnetName"] = vnicEtherIf.Oper_vnet_name
    vnicEtherIfMap["owner"] = vnicEtherIf.Owner
    vnicEtherIfMap["propAcl"] = vnicEtherIf.Prop_acl
    vnicEtherIfMap["pubNwId"] = vnicEtherIf.Pub_nw_id
    vnicEtherIfMap["sacl"] = vnicEtherIf.Sacl
    vnicEtherIfMap["sharing"] = vnicEtherIf.Sharing
    vnicEtherIfMap["switchId"] = vnicEtherIf.Switch_id
    vnicEtherIfMap["type"] = vnicEtherIf.Type
    vnicEtherIfMap["vnet"] = vnicEtherIf.Vnet
    
	return vnicEtherIfMap, nil

}

func VnicEtherIfFromDoc(doc *etree.Document, rootClass string) *VnicEtherIf {
	element, err := GetMoElement(doc, rootClass, VnicEtherIfClassName)
	if err != nil {
		return nil
	}
	return &VnicEtherIf{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicEtherIfClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicEtherIfAttributes: VnicEtherIfAttributes{

        
			Addr:  element.SelectAttrValue("addr", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_qualifier:  element.SelectAttrValue("configQualifier", ""),
			Default_net:  element.SelectAttrValue("defaultNet", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Oper_primary_vnet_dn:  element.SelectAttrValue("operPrimaryVnetDn", ""),
			Oper_primary_vnet_name:  element.SelectAttrValue("operPrimaryVnetName", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Oper_vnet_dn:  element.SelectAttrValue("operVnetDn", ""),
			Oper_vnet_name:  element.SelectAttrValue("operVnetName", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Pub_nw_id:  element.SelectAttrValue("pubNwId", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Sharing:  element.SelectAttrValue("sharing", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Type:  element.SelectAttrValue("type", ""),
			Vnet:  element.SelectAttrValue("vnet", ""),
		},
    
	}
}