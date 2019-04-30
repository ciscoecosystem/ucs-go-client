package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicLanConnTemplClassName = "vnicLanConnTempl"

type VnicLanConnTempl struct {
    BaseAttributes
    VnicLanConnTemplAttributes
}

type VnicLanConnTemplAttributes struct{
    Admin_cdn_name string `xml:",omitempty"`
    Cdn_source string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Ident_pool_name string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Mtu string `xml:",omitempty"`
    Nw_ctrl_policy_name string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Oper_nw_ctrl_policy_name string `xml:",omitempty"`
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

func NewVnicLanConnTempl(vnicLanConnTemplRn,parentDn,description string, vnicLanConnTemplAttributes VnicLanConnTemplAttributes) *VnicLanConnTempl { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicLanConnTemplRn)
	return &VnicLanConnTempl{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicLanConnTemplClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicLanConnTemplAttributes: vnicLanConnTemplAttributes,
	}
}

func (vnicLanConnTempl *VnicLanConnTempl) ToMap() (map[string]string, error) {
	vnicLanConnTemplMap, err := vnicLanConnTempl.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicLanConnTemplMap["adminCdnName"] = vnicLanConnTempl.Admin_cdn_name
    vnicLanConnTemplMap["cdnSource"] = vnicLanConnTempl.Cdn_source
    vnicLanConnTemplMap["childAction"] = vnicLanConnTempl.Child_action
    vnicLanConnTemplMap["identPoolName"] = vnicLanConnTempl.Ident_pool_name
    vnicLanConnTemplMap["intId"] = vnicLanConnTempl.Int_id
    vnicLanConnTemplMap["mtu"] = vnicLanConnTempl.Mtu
    vnicLanConnTemplMap["nwCtrlPolicyName"] = vnicLanConnTempl.Nw_ctrl_policy_name
    vnicLanConnTemplMap["operIdentPoolName"] = vnicLanConnTempl.Oper_ident_pool_name
    vnicLanConnTemplMap["operNwCtrlPolicyName"] = vnicLanConnTempl.Oper_nw_ctrl_policy_name
    vnicLanConnTemplMap["operPeerRedundancyTemplName"] = vnicLanConnTempl.Oper_peer_redundancy_templ_name
    vnicLanConnTemplMap["operQosPolicyName"] = vnicLanConnTempl.Oper_qos_policy_name
    vnicLanConnTemplMap["operStatsPolicyName"] = vnicLanConnTempl.Oper_stats_policy_name
    vnicLanConnTemplMap["peerRedundancyTemplName"] = vnicLanConnTempl.Peer_redundancy_templ_name
    vnicLanConnTemplMap["pinToGroupName"] = vnicLanConnTempl.Pin_to_group_name
    vnicLanConnTemplMap["policyLevel"] = vnicLanConnTempl.Policy_level
    vnicLanConnTemplMap["policyOwner"] = vnicLanConnTempl.Policy_owner
    vnicLanConnTemplMap["qosPolicyName"] = vnicLanConnTempl.Qos_policy_name
    vnicLanConnTemplMap["redundancyPairType"] = vnicLanConnTempl.Redundancy_pair_type
    vnicLanConnTemplMap["sacl"] = vnicLanConnTempl.Sacl
    vnicLanConnTemplMap["statsPolicyName"] = vnicLanConnTempl.Stats_policy_name
    vnicLanConnTemplMap["switchId"] = vnicLanConnTempl.Switch_id
    vnicLanConnTemplMap["target"] = vnicLanConnTempl.Target
    vnicLanConnTemplMap["templType"] = vnicLanConnTempl.Templ_type
    
	return vnicLanConnTemplMap, nil

}

func VnicLanConnTemplFromDoc(doc *etree.Document, rootClass string) *VnicLanConnTempl {
	element, err := GetMoElement(doc, rootClass, VnicLanConnTemplClassName)
	if err != nil {
		return nil
	}
	return &VnicLanConnTempl{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicLanConnTemplClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicLanConnTemplAttributes: VnicLanConnTemplAttributes{

        
			Admin_cdn_name:  element.SelectAttrValue("adminCdnName", ""),
			Cdn_source:  element.SelectAttrValue("cdnSource", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Ident_pool_name:  element.SelectAttrValue("identPoolName", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Mtu:  element.SelectAttrValue("mtu", ""),
			Nw_ctrl_policy_name:  element.SelectAttrValue("nwCtrlPolicyName", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Oper_nw_ctrl_policy_name:  element.SelectAttrValue("operNwCtrlPolicyName", ""),
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