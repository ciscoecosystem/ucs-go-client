package models

import (


	"github.com/beevik/etree"

)

const FabricVlanClassName = "fabricVlan"

type FabricVlan struct {
    BaseAttributes
    FabricVlanAttributes
}

type FabricVlanAttributes struct{
    Assoc_primary_vlan_state string `xml:",omitempty"`
    Assoc_primary_vlan_switch_id string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Cloud string `xml:",omitempty"`
    Compression_type string `xml:",omitempty"`
    Config_issues string `xml:",omitempty"`
    Config_overlap string `xml:",omitempty"`
    Default_net string `xml:",omitempty"`
    Ep_dn string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    R_global string `xml:",omitempty"`
    Fabric_vlan_id string `xml:",omitempty"`
    If_role string `xml:",omitempty"`
    If_type string `xml:",omitempty"`
    Local string `xml:",omitempty"`
    Locale string `xml:",omitempty"`
    Mcast_policy_name string `xml:",omitempty"`
    Oper_mcast_policy_name string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Overlap_state_for_a string `xml:",omitempty"`
    Overlap_state_for_b string `xml:",omitempty"`
    Peer_dn string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Pub_nw_dn string `xml:",omitempty"`
    Pub_nw_id string `xml:",omitempty"`
    Pub_nw_name string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Sharing string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Transport string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    

}

func NewFabricVlan(fabricVlanRn,description string, fabricVlanAttributes FabricVlanAttributes) *FabricVlan {
    dn := fabricVlanRn
	return &FabricVlan{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         FabricVlanClassName,
			Status:            "created, modified",
			Description:       description,
		},

		FabricVlanAttributes: fabricVlanAttributes,
	}
}

func (fabricVlan *FabricVlan) ToMap() (map[string]string, error) {
	fabricVlanMap, err := fabricVlan.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    fabricVlanMap["assocPrimaryVlanState"] = fabricVlan.Assoc_primary_vlan_state
    fabricVlanMap["assocPrimaryVlanSwitchId"] = fabricVlan.Assoc_primary_vlan_switch_id
    fabricVlanMap["childAction"] = fabricVlan.Child_action
    fabricVlanMap["cloud"] = fabricVlan.Cloud
    fabricVlanMap["compressionType"] = fabricVlan.Compression_type
    fabricVlanMap["configIssues"] = fabricVlan.Config_issues
    fabricVlanMap["configOverlap"] = fabricVlan.Config_overlap
    fabricVlanMap["defaultNet"] = fabricVlan.Default_net
    fabricVlanMap["epDn"] = fabricVlan.Ep_dn
    fabricVlanMap["fltAggr"] = fabricVlan.Flt_aggr
    fabricVlanMap["global"] = fabricVlan.R_global
    fabricVlanMap["id"] = fabricVlan.Fabric_vlan_id
    fabricVlanMap["ifRole"] = fabricVlan.If_role
    fabricVlanMap["ifType"] = fabricVlan.If_type
    fabricVlanMap["local"] = fabricVlan.Local
    fabricVlanMap["locale"] = fabricVlan.Locale
    fabricVlanMap["mcastPolicyName"] = fabricVlan.Mcast_policy_name
    fabricVlanMap["operMcastPolicyName"] = fabricVlan.Oper_mcast_policy_name
    fabricVlanMap["operState"] = fabricVlan.Oper_state
    fabricVlanMap["overlapStateForA"] = fabricVlan.Overlap_state_for_a
    fabricVlanMap["overlapStateForB"] = fabricVlan.Overlap_state_for_b
    fabricVlanMap["peerDn"] = fabricVlan.Peer_dn
    fabricVlanMap["policyOwner"] = fabricVlan.Policy_owner
    fabricVlanMap["pubNwDn"] = fabricVlan.Pub_nw_dn
    fabricVlanMap["pubNwId"] = fabricVlan.Pub_nw_id
    fabricVlanMap["pubNwName"] = fabricVlan.Pub_nw_name
    fabricVlanMap["sacl"] = fabricVlan.Sacl
    fabricVlanMap["sharing"] = fabricVlan.Sharing
    fabricVlanMap["switchId"] = fabricVlan.Switch_id
    fabricVlanMap["transport"] = fabricVlan.Transport
    fabricVlanMap["type"] = fabricVlan.Type
    
	return fabricVlanMap, nil

}

func FabricVlanFromDoc(doc *etree.Document, rootClass string) *FabricVlan {
	element, err := GetMoElement(doc, rootClass, FabricVlanClassName)
	if err != nil {
		return nil
	}
	return &FabricVlan{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         FabricVlanClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		FabricVlanAttributes: FabricVlanAttributes{

        
			Assoc_primary_vlan_state:  element.SelectAttrValue("assocPrimaryVlanState", ""),
			Assoc_primary_vlan_switch_id:  element.SelectAttrValue("assocPrimaryVlanSwitchId", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Cloud:  element.SelectAttrValue("cloud", ""),
			Compression_type:  element.SelectAttrValue("compressionType", ""),
			Config_issues:  element.SelectAttrValue("configIssues", ""),
			Config_overlap:  element.SelectAttrValue("configOverlap", ""),
			Default_net:  element.SelectAttrValue("defaultNet", ""),
			Ep_dn:  element.SelectAttrValue("epDn", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			R_global:  element.SelectAttrValue("global", ""),
			Fabric_vlan_id:  element.SelectAttrValue("id", ""),
			If_role:  element.SelectAttrValue("ifRole", ""),
			If_type:  element.SelectAttrValue("ifType", ""),
			Local:  element.SelectAttrValue("local", ""),
			Locale:  element.SelectAttrValue("locale", ""),
			Mcast_policy_name:  element.SelectAttrValue("mcastPolicyName", ""),
			Oper_mcast_policy_name:  element.SelectAttrValue("operMcastPolicyName", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Overlap_state_for_a:  element.SelectAttrValue("overlapStateForA", ""),
			Overlap_state_for_b:  element.SelectAttrValue("overlapStateForB", ""),
			Peer_dn:  element.SelectAttrValue("peerDn", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Pub_nw_dn:  element.SelectAttrValue("pubNwDn", ""),
			Pub_nw_id:  element.SelectAttrValue("pubNwId", ""),
			Pub_nw_name:  element.SelectAttrValue("pubNwName", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Sharing:  element.SelectAttrValue("sharing", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Transport:  element.SelectAttrValue("transport", ""),
			Type:  element.SelectAttrValue("type", ""),
		},
    
	}
}