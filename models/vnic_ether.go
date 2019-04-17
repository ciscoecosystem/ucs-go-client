package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicEtherClassName = "vnicEther"

type VnicEther struct {
    BaseAttributes
    VnicEtherAttributes
}

type VnicEtherAttributes struct{
    Adaptor_profile_name string `xml:",omitempty"`
    Addr string `xml:",omitempty"`
    Admin_cdn_name string `xml:",omitempty"`
    Admin_host_port string `xml:",omitempty"`
    Admin_vcon string `xml:",omitempty"`
    Boot_dev string `xml:",omitempty"`
    Cdn_prop_in_sync string `xml:",omitempty"`
    Cdn_source string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Config_qualifier string `xml:",omitempty"`
    Config_state string `xml:",omitempty"`
    Dynamic_id string `xml:",omitempty"`
    Equipment_dn string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Ident_pool_name string `xml:",omitempty"`
    Inst_type string `xml:",omitempty"`
    Mtu string `xml:",omitempty"`
    Nw_ctrl_policy_name string `xml:",omitempty"`
    Nw_templ_name string `xml:",omitempty"`
    Oper_adaptor_profile_name string `xml:",omitempty"`
    Oper_cdn_name string `xml:",omitempty"`
    Oper_host_port string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Oper_nw_ctrl_policy_name string `xml:",omitempty"`
    Oper_nw_templ_name string `xml:",omitempty"`
    Oper_order string `xml:",omitempty"`
    Oper_pin_to_group_name string `xml:",omitempty"`
    Oper_qos_policy_name string `xml:",omitempty"`
    Oper_speed string `xml:",omitempty"`
    Oper_stats_policy_name string `xml:",omitempty"`
    Oper_vcon string `xml:",omitempty"`
    Order string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Pf_dn string `xml:",omitempty"`
    Pin_to_group_name string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Purpose string `xml:",omitempty"`
    Qos_policy_name string `xml:",omitempty"`
    Redundancy_pair_type string `xml:",omitempty"`
    Redundancy_peer string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Stats_policy_name string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    Virtualization_preference string `xml:",omitempty"`
    

}

func NewVnicEther(vnicEtherRn,parentDn,description string, vnicEtherAttributes VnicEtherAttributes) *VnicEther { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicEtherRn)
	return &VnicEther{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicEtherClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicEtherAttributes: vnicEtherAttributes,
	}
}

func (vnicEther *VnicEther) ToMap() (map[string]string, error) {
	vnicEtherMap, err := vnicEther.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicEtherMap["adaptorProfileName"] = vnicEther.Adaptor_profile_name
    vnicEtherMap["addr"] = vnicEther.Addr
    vnicEtherMap["adminCdnName"] = vnicEther.Admin_cdn_name
    vnicEtherMap["adminHostPort"] = vnicEther.Admin_host_port
    vnicEtherMap["adminVcon"] = vnicEther.Admin_vcon
    vnicEtherMap["bootDev"] = vnicEther.Boot_dev
    vnicEtherMap["cdnPropInSync"] = vnicEther.Cdn_prop_in_sync
    vnicEtherMap["cdnSource"] = vnicEther.Cdn_source
    vnicEtherMap["childAction"] = vnicEther.Child_action
    vnicEtherMap["configQualifier"] = vnicEther.Config_qualifier
    vnicEtherMap["configState"] = vnicEther.Config_state
    vnicEtherMap["dynamicId"] = vnicEther.Dynamic_id
    vnicEtherMap["equipmentDn"] = vnicEther.Equipment_dn
    vnicEtherMap["fltAggr"] = vnicEther.Flt_aggr
    vnicEtherMap["identPoolName"] = vnicEther.Ident_pool_name
    vnicEtherMap["instType"] = vnicEther.Inst_type
    vnicEtherMap["mtu"] = vnicEther.Mtu
    vnicEtherMap["nwCtrlPolicyName"] = vnicEther.Nw_ctrl_policy_name
    vnicEtherMap["nwTemplName"] = vnicEther.Nw_templ_name
    vnicEtherMap["operAdaptorProfileName"] = vnicEther.Oper_adaptor_profile_name
    vnicEtherMap["operCdnName"] = vnicEther.Oper_cdn_name
    vnicEtherMap["operHostPort"] = vnicEther.Oper_host_port
    vnicEtherMap["operIdentPoolName"] = vnicEther.Oper_ident_pool_name
    vnicEtherMap["operNwCtrlPolicyName"] = vnicEther.Oper_nw_ctrl_policy_name
    vnicEtherMap["operNwTemplName"] = vnicEther.Oper_nw_templ_name
    vnicEtherMap["operOrder"] = vnicEther.Oper_order
    vnicEtherMap["operPinToGroupName"] = vnicEther.Oper_pin_to_group_name
    vnicEtherMap["operQosPolicyName"] = vnicEther.Oper_qos_policy_name
    vnicEtherMap["operSpeed"] = vnicEther.Oper_speed
    vnicEtherMap["operStatsPolicyName"] = vnicEther.Oper_stats_policy_name
    vnicEtherMap["operVcon"] = vnicEther.Oper_vcon
    vnicEtherMap["order"] = vnicEther.Order
    vnicEtherMap["owner"] = vnicEther.Owner
    vnicEtherMap["pfDn"] = vnicEther.Pf_dn
    vnicEtherMap["pinToGroupName"] = vnicEther.Pin_to_group_name
    vnicEtherMap["propAcl"] = vnicEther.Prop_acl
    vnicEtherMap["purpose"] = vnicEther.Purpose
    vnicEtherMap["qosPolicyName"] = vnicEther.Qos_policy_name
    vnicEtherMap["redundancyPairType"] = vnicEther.Redundancy_pair_type
    vnicEtherMap["redundancyPeer"] = vnicEther.Redundancy_peer
    vnicEtherMap["sacl"] = vnicEther.Sacl
    vnicEtherMap["statsPolicyName"] = vnicEther.Stats_policy_name
    vnicEtherMap["switchId"] = vnicEther.Switch_id
    vnicEtherMap["type"] = vnicEther.Type
    vnicEtherMap["virtualizationPreference"] = vnicEther.Virtualization_preference
    
	return vnicEtherMap, nil

}

func VnicEtherFromDoc(doc *etree.Document, rootClass string) *VnicEther {
	element, err := GetMoElement(doc, rootClass, VnicEtherClassName)
	if err != nil {
		return nil
	}
	return &VnicEther{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicEtherClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicEtherAttributes: VnicEtherAttributes{

        
			Adaptor_profile_name:  element.SelectAttrValue("adaptorProfileName", ""),
			Addr:  element.SelectAttrValue("addr", ""),
			Admin_cdn_name:  element.SelectAttrValue("adminCdnName", ""),
			Admin_host_port:  element.SelectAttrValue("adminHostPort", ""),
			Admin_vcon:  element.SelectAttrValue("adminVcon", ""),
			Boot_dev:  element.SelectAttrValue("bootDev", ""),
			Cdn_prop_in_sync:  element.SelectAttrValue("cdnPropInSync", ""),
			Cdn_source:  element.SelectAttrValue("cdnSource", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_qualifier:  element.SelectAttrValue("configQualifier", ""),
			Config_state:  element.SelectAttrValue("configState", ""),
			Dynamic_id:  element.SelectAttrValue("dynamicId", ""),
			Equipment_dn:  element.SelectAttrValue("equipmentDn", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Ident_pool_name:  element.SelectAttrValue("identPoolName", ""),
			Inst_type:  element.SelectAttrValue("instType", ""),
			Mtu:  element.SelectAttrValue("mtu", ""),
			Nw_ctrl_policy_name:  element.SelectAttrValue("nwCtrlPolicyName", ""),
			Nw_templ_name:  element.SelectAttrValue("nwTemplName", ""),
			Oper_adaptor_profile_name:  element.SelectAttrValue("operAdaptorProfileName", ""),
			Oper_cdn_name:  element.SelectAttrValue("operCdnName", ""),
			Oper_host_port:  element.SelectAttrValue("operHostPort", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Oper_nw_ctrl_policy_name:  element.SelectAttrValue("operNwCtrlPolicyName", ""),
			Oper_nw_templ_name:  element.SelectAttrValue("operNwTemplName", ""),
			Oper_order:  element.SelectAttrValue("operOrder", ""),
			Oper_pin_to_group_name:  element.SelectAttrValue("operPinToGroupName", ""),
			Oper_qos_policy_name:  element.SelectAttrValue("operQosPolicyName", ""),
			Oper_speed:  element.SelectAttrValue("operSpeed", ""),
			Oper_stats_policy_name:  element.SelectAttrValue("operStatsPolicyName", ""),
			Oper_vcon:  element.SelectAttrValue("operVcon", ""),
			Order:  element.SelectAttrValue("order", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Pf_dn:  element.SelectAttrValue("pfDn", ""),
			Pin_to_group_name:  element.SelectAttrValue("pinToGroupName", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Purpose:  element.SelectAttrValue("purpose", ""),
			Qos_policy_name:  element.SelectAttrValue("qosPolicyName", ""),
			Redundancy_pair_type:  element.SelectAttrValue("redundancyPairType", ""),
			Redundancy_peer:  element.SelectAttrValue("redundancyPeer", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Stats_policy_name:  element.SelectAttrValue("statsPolicyName", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Type:  element.SelectAttrValue("type", ""),
			Virtualization_preference:  element.SelectAttrValue("virtualizationPreference", ""),
		},
    
	}
}