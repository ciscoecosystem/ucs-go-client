package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicIScsiLCPClassName = "vnicIScsiLCP"

type VnicIScsiLCP struct {
    BaseAttributes
    VnicIScsiLCPAttributes
}

type VnicIScsiLCPAttributes struct{
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
    Equipment_dn string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Ident_pool_name string `xml:",omitempty"`
    Inst_type string `xml:",omitempty"`
    Nw_templ_name string `xml:",omitempty"`
    Oper_adaptor_profile_name string `xml:",omitempty"`
    Oper_cdn_name string `xml:",omitempty"`
    Oper_host_port string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Oper_order string `xml:",omitempty"`
    Oper_speed string `xml:",omitempty"`
    Oper_stats_policy_name string `xml:",omitempty"`
    Oper_vcon string `xml:",omitempty"`
    Order string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Pin_to_group_name string `xml:",omitempty"`
    Qos_policy_name string `xml:",omitempty"`
    Redundancy_pair_type string `xml:",omitempty"`
    Redundancy_peer string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Stats_policy_name string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    Vnic_name string `xml:",omitempty"`
    

}

func NewVnicIScsiLCP(vnicIScsiLCPRn,parentDn,description string, vnicIScsiLCPAttributes VnicIScsiLCPAttributes) *VnicIScsiLCP { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicIScsiLCPRn)
	return &VnicIScsiLCP{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicIScsiLCPClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicIScsiLCPAttributes: vnicIScsiLCPAttributes,
	}
}

func (vnicIScsiLCP *VnicIScsiLCP) ToMap() (map[string]string, error) {
	vnicIScsiLCPMap, err := vnicIScsiLCP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicIScsiLCPMap["adaptorProfileName"] = vnicIScsiLCP.Adaptor_profile_name
    vnicIScsiLCPMap["addr"] = vnicIScsiLCP.Addr
    vnicIScsiLCPMap["adminCdnName"] = vnicIScsiLCP.Admin_cdn_name
    vnicIScsiLCPMap["adminHostPort"] = vnicIScsiLCP.Admin_host_port
    vnicIScsiLCPMap["adminVcon"] = vnicIScsiLCP.Admin_vcon
    vnicIScsiLCPMap["bootDev"] = vnicIScsiLCP.Boot_dev
    vnicIScsiLCPMap["cdnPropInSync"] = vnicIScsiLCP.Cdn_prop_in_sync
    vnicIScsiLCPMap["cdnSource"] = vnicIScsiLCP.Cdn_source
    vnicIScsiLCPMap["childAction"] = vnicIScsiLCP.Child_action
    vnicIScsiLCPMap["configQualifier"] = vnicIScsiLCP.Config_qualifier
    vnicIScsiLCPMap["configState"] = vnicIScsiLCP.Config_state
    vnicIScsiLCPMap["equipmentDn"] = vnicIScsiLCP.Equipment_dn
    vnicIScsiLCPMap["fltAggr"] = vnicIScsiLCP.Flt_aggr
    vnicIScsiLCPMap["identPoolName"] = vnicIScsiLCP.Ident_pool_name
    vnicIScsiLCPMap["instType"] = vnicIScsiLCP.Inst_type
    vnicIScsiLCPMap["nwTemplName"] = vnicIScsiLCP.Nw_templ_name
    vnicIScsiLCPMap["operAdaptorProfileName"] = vnicIScsiLCP.Oper_adaptor_profile_name
    vnicIScsiLCPMap["operCdnName"] = vnicIScsiLCP.Oper_cdn_name
    vnicIScsiLCPMap["operHostPort"] = vnicIScsiLCP.Oper_host_port
    vnicIScsiLCPMap["operIdentPoolName"] = vnicIScsiLCP.Oper_ident_pool_name
    vnicIScsiLCPMap["operOrder"] = vnicIScsiLCP.Oper_order
    vnicIScsiLCPMap["operSpeed"] = vnicIScsiLCP.Oper_speed
    vnicIScsiLCPMap["operStatsPolicyName"] = vnicIScsiLCP.Oper_stats_policy_name
    vnicIScsiLCPMap["operVcon"] = vnicIScsiLCP.Oper_vcon
    vnicIScsiLCPMap["order"] = vnicIScsiLCP.Order
    vnicIScsiLCPMap["owner"] = vnicIScsiLCP.Owner
    vnicIScsiLCPMap["pinToGroupName"] = vnicIScsiLCP.Pin_to_group_name
    vnicIScsiLCPMap["qosPolicyName"] = vnicIScsiLCP.Qos_policy_name
    vnicIScsiLCPMap["redundancyPairType"] = vnicIScsiLCP.Redundancy_pair_type
    vnicIScsiLCPMap["redundancyPeer"] = vnicIScsiLCP.Redundancy_peer
    vnicIScsiLCPMap["sacl"] = vnicIScsiLCP.Sacl
    vnicIScsiLCPMap["statsPolicyName"] = vnicIScsiLCP.Stats_policy_name
    vnicIScsiLCPMap["switchId"] = vnicIScsiLCP.Switch_id
    vnicIScsiLCPMap["type"] = vnicIScsiLCP.Type
    vnicIScsiLCPMap["vnicName"] = vnicIScsiLCP.Vnic_name
    
	return vnicIScsiLCPMap, nil

}

func VnicIScsiLCPFromDoc(doc *etree.Document, rootClass string) *VnicIScsiLCP {
	element, err := GetMoElement(doc, rootClass, VnicIScsiLCPClassName)
	if err != nil {
		return nil
	}
	return &VnicIScsiLCP{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicIScsiLCPClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicIScsiLCPAttributes: VnicIScsiLCPAttributes{

        
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
			Equipment_dn:  element.SelectAttrValue("equipmentDn", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Ident_pool_name:  element.SelectAttrValue("identPoolName", ""),
			Inst_type:  element.SelectAttrValue("instType", ""),
			Nw_templ_name:  element.SelectAttrValue("nwTemplName", ""),
			Oper_adaptor_profile_name:  element.SelectAttrValue("operAdaptorProfileName", ""),
			Oper_cdn_name:  element.SelectAttrValue("operCdnName", ""),
			Oper_host_port:  element.SelectAttrValue("operHostPort", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Oper_order:  element.SelectAttrValue("operOrder", ""),
			Oper_speed:  element.SelectAttrValue("operSpeed", ""),
			Oper_stats_policy_name:  element.SelectAttrValue("operStatsPolicyName", ""),
			Oper_vcon:  element.SelectAttrValue("operVcon", ""),
			Order:  element.SelectAttrValue("order", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Pin_to_group_name:  element.SelectAttrValue("pinToGroupName", ""),
			Qos_policy_name:  element.SelectAttrValue("qosPolicyName", ""),
			Redundancy_pair_type:  element.SelectAttrValue("redundancyPairType", ""),
			Redundancy_peer:  element.SelectAttrValue("redundancyPeer", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Stats_policy_name:  element.SelectAttrValue("statsPolicyName", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Type:  element.SelectAttrValue("type", ""),
			Vnic_name:  element.SelectAttrValue("vnicName", ""),
		},
    
	}
}