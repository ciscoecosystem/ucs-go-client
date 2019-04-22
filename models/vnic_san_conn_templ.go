package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicSanConnTemplClassName = "vnicSanConnTempl"

type VnicSanConnTempl struct {
    BaseAttributes
    VnicSanConnTemplAttributes
}

type VnicSanConnTemplAttributes struct{
    Child_action string `xml:",omitempty"`
    Ident_pool_name string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Max_data_field_size string `xml:",omitempty"`
    Nw_ctrl_policy_name string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Oper_peer_redundancy_templ_name string `xml:",omitempty"`
    Oper_qos_policy_name string `xml:",omitempty"`
    Oper_stats_policy_name string `xml:",omitempty"`
    Peer_redundancy_templ_name string `xml:",omitempty"`
    Pin_to_group_name string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Qos_policy_name string `xml:",omitempty"`
    Redundancy_pair_type string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Stats_policy_name string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Target string `xml:",omitempty"`
    Templ_type string `xml:",omitempty"`
    

}

func NewVnicSanConnTempl(vnicSanConnTemplRn,parentDn,description string, vnicSanConnTemplAttributes VnicSanConnTemplAttributes) *VnicSanConnTempl { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicSanConnTemplRn)
	return &VnicSanConnTempl{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicSanConnTemplClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicSanConnTemplAttributes: vnicSanConnTemplAttributes,
	}
}

func (vnicSanConnTempl *VnicSanConnTempl) ToMap() (map[string]string, error) {
	vnicSanConnTemplMap, err := vnicSanConnTempl.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicSanConnTemplMap["childAction"] = vnicSanConnTempl.Child_action
    vnicSanConnTemplMap["identPoolName"] = vnicSanConnTempl.Ident_pool_name
    vnicSanConnTemplMap["intId"] = vnicSanConnTempl.Int_id
    vnicSanConnTemplMap["maxDataFieldSize"] = vnicSanConnTempl.Max_data_field_size
    vnicSanConnTemplMap["nwCtrlPolicyName"] = vnicSanConnTempl.Nw_ctrl_policy_name
    vnicSanConnTemplMap["operIdentPoolName"] = vnicSanConnTempl.Oper_ident_pool_name
    vnicSanConnTemplMap["operPeerRedundancyTemplName"] = vnicSanConnTempl.Oper_peer_redundancy_templ_name
    vnicSanConnTemplMap["operQosPolicyName"] = vnicSanConnTempl.Oper_qos_policy_name
    vnicSanConnTemplMap["operStatsPolicyName"] = vnicSanConnTempl.Oper_stats_policy_name
    vnicSanConnTemplMap["peerRedundancyTemplName"] = vnicSanConnTempl.Peer_redundancy_templ_name
    vnicSanConnTemplMap["pinToGroupName"] = vnicSanConnTempl.Pin_to_group_name
    vnicSanConnTemplMap["policyLevel"] = vnicSanConnTempl.Policy_level
    vnicSanConnTemplMap["policyOwner"] = vnicSanConnTempl.Policy_owner
    vnicSanConnTemplMap["qosPolicyName"] = vnicSanConnTempl.Qos_policy_name
    vnicSanConnTemplMap["redundancyPairType"] = vnicSanConnTempl.Redundancy_pair_type
    vnicSanConnTemplMap["sacl"] = vnicSanConnTempl.Sacl
    vnicSanConnTemplMap["statsPolicyName"] = vnicSanConnTempl.Stats_policy_name
    vnicSanConnTemplMap["switchId"] = vnicSanConnTempl.Switch_id
    vnicSanConnTemplMap["target"] = vnicSanConnTempl.Target
    vnicSanConnTemplMap["templType"] = vnicSanConnTempl.Templ_type
    
	return vnicSanConnTemplMap, nil

}

func VnicSanConnTemplFromDoc(doc *etree.Document, rootClass string) *VnicSanConnTempl {
	element, err := GetMoElement(doc, rootClass, VnicSanConnTemplClassName)
	if err != nil {
		return nil
	}
	return &VnicSanConnTempl{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicSanConnTemplClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicSanConnTemplAttributes: VnicSanConnTemplAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Ident_pool_name:  element.SelectAttrValue("identPoolName", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Max_data_field_size:  element.SelectAttrValue("maxDataFieldSize", ""),
			Nw_ctrl_policy_name:  element.SelectAttrValue("nwCtrlPolicyName", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Oper_peer_redundancy_templ_name:  element.SelectAttrValue("operPeerRedundancyTemplName", ""),
			Oper_qos_policy_name:  element.SelectAttrValue("operQosPolicyName", ""),
			Oper_stats_policy_name:  element.SelectAttrValue("operStatsPolicyName", ""),
			Peer_redundancy_templ_name:  element.SelectAttrValue("peerRedundancyTemplName", ""),
			Pin_to_group_name:  element.SelectAttrValue("pinToGroupName", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Qos_policy_name:  element.SelectAttrValue("qosPolicyName", ""),
			Redundancy_pair_type:  element.SelectAttrValue("redundancyPairType", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Stats_policy_name:  element.SelectAttrValue("statsPolicyName", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Target:  element.SelectAttrValue("target", ""),
			Templ_type:  element.SelectAttrValue("templType", ""),
		},
    
	}
}