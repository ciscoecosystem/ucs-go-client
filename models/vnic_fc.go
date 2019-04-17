package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicFcClassName = "vnicFc"

type VnicFc struct {
    BaseAttributes
    VnicFcAttributes
}

type VnicFcAttributes struct{
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
    Is_supported string `xml:",omitempty"`
    Max_data_field_size string `xml:",omitempty"`
    Node_addr string `xml:",omitempty"`
    Nw_templ_name string `xml:",omitempty"`
    Oper_adaptor_profile_name string `xml:",omitempty"`
    Oper_cdn_name string `xml:",omitempty"`
    Oper_host_port string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Oper_nw_templ_name string `xml:",omitempty"`
    Oper_order string `xml:",omitempty"`
    Oper_pin_to_group_name string `xml:",omitempty"`
    Oper_qos_policy_name string `xml:",omitempty"`
    Oper_speed string `xml:",omitempty"`
    Oper_stats_policy_name string `xml:",omitempty"`
    Oper_vcon string `xml:",omitempty"`
    Order string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Pers_bind string `xml:",omitempty"`
    Pers_bind_clear string `xml:",omitempty"`
    Pin_to_group_name string `xml:",omitempty"`
    Qos_policy_name string `xml:",omitempty"`
    Redundancy_pair_type string `xml:",omitempty"`
    Redundancy_peer string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Stats_policy_name string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    

}

func NewVnicFc(vnicFcRn,parentDn,description string, vnicFcAttributes VnicFcAttributes) *VnicFc { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicFcRn)
	return &VnicFc{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicFcClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicFcAttributes: vnicFcAttributes,
	}
}

func (vnicFc *VnicFc) ToMap() (map[string]string, error) {
	vnicFcMap, err := vnicFc.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicFcMap["adaptorProfileName"] = vnicFc.Adaptor_profile_name
    vnicFcMap["addr"] = vnicFc.Addr
    vnicFcMap["adminCdnName"] = vnicFc.Admin_cdn_name
    vnicFcMap["adminHostPort"] = vnicFc.Admin_host_port
    vnicFcMap["adminVcon"] = vnicFc.Admin_vcon
    vnicFcMap["bootDev"] = vnicFc.Boot_dev
    vnicFcMap["cdnPropInSync"] = vnicFc.Cdn_prop_in_sync
    vnicFcMap["cdnSource"] = vnicFc.Cdn_source
    vnicFcMap["childAction"] = vnicFc.Child_action
    vnicFcMap["configQualifier"] = vnicFc.Config_qualifier
    vnicFcMap["configState"] = vnicFc.Config_state
    vnicFcMap["equipmentDn"] = vnicFc.Equipment_dn
    vnicFcMap["fltAggr"] = vnicFc.Flt_aggr
    vnicFcMap["identPoolName"] = vnicFc.Ident_pool_name
    vnicFcMap["instType"] = vnicFc.Inst_type
    vnicFcMap["isSupported"] = vnicFc.Is_supported
    vnicFcMap["maxDataFieldSize"] = vnicFc.Max_data_field_size
    vnicFcMap["nodeAddr"] = vnicFc.Node_addr
    vnicFcMap["nwTemplName"] = vnicFc.Nw_templ_name
    vnicFcMap["operAdaptorProfileName"] = vnicFc.Oper_adaptor_profile_name
    vnicFcMap["operCdnName"] = vnicFc.Oper_cdn_name
    vnicFcMap["operHostPort"] = vnicFc.Oper_host_port
    vnicFcMap["operIdentPoolName"] = vnicFc.Oper_ident_pool_name
    vnicFcMap["operNwTemplName"] = vnicFc.Oper_nw_templ_name
    vnicFcMap["operOrder"] = vnicFc.Oper_order
    vnicFcMap["operPinToGroupName"] = vnicFc.Oper_pin_to_group_name
    vnicFcMap["operQosPolicyName"] = vnicFc.Oper_qos_policy_name
    vnicFcMap["operSpeed"] = vnicFc.Oper_speed
    vnicFcMap["operStatsPolicyName"] = vnicFc.Oper_stats_policy_name
    vnicFcMap["operVcon"] = vnicFc.Oper_vcon
    vnicFcMap["order"] = vnicFc.Order
    vnicFcMap["owner"] = vnicFc.Owner
    vnicFcMap["persBind"] = vnicFc.Pers_bind
    vnicFcMap["persBindClear"] = vnicFc.Pers_bind_clear
    vnicFcMap["pinToGroupName"] = vnicFc.Pin_to_group_name
    vnicFcMap["qosPolicyName"] = vnicFc.Qos_policy_name
    vnicFcMap["redundancyPairType"] = vnicFc.Redundancy_pair_type
    vnicFcMap["redundancyPeer"] = vnicFc.Redundancy_peer
    vnicFcMap["sacl"] = vnicFc.Sacl
    vnicFcMap["statsPolicyName"] = vnicFc.Stats_policy_name
    vnicFcMap["switchId"] = vnicFc.Switch_id
    vnicFcMap["type"] = vnicFc.Type
    
	return vnicFcMap, nil

}

func VnicFcFromDoc(doc *etree.Document, rootClass string) *VnicFc {
	element, err := GetMoElement(doc, rootClass, VnicFcClassName)
	if err != nil {
		return nil
	}
	return &VnicFc{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicFcClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicFcAttributes: VnicFcAttributes{

        
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
			Is_supported:  element.SelectAttrValue("isSupported", ""),
			Max_data_field_size:  element.SelectAttrValue("maxDataFieldSize", ""),
			Node_addr:  element.SelectAttrValue("nodeAddr", ""),
			Nw_templ_name:  element.SelectAttrValue("nwTemplName", ""),
			Oper_adaptor_profile_name:  element.SelectAttrValue("operAdaptorProfileName", ""),
			Oper_cdn_name:  element.SelectAttrValue("operCdnName", ""),
			Oper_host_port:  element.SelectAttrValue("operHostPort", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Oper_nw_templ_name:  element.SelectAttrValue("operNwTemplName", ""),
			Oper_order:  element.SelectAttrValue("operOrder", ""),
			Oper_pin_to_group_name:  element.SelectAttrValue("operPinToGroupName", ""),
			Oper_qos_policy_name:  element.SelectAttrValue("operQosPolicyName", ""),
			Oper_speed:  element.SelectAttrValue("operSpeed", ""),
			Oper_stats_policy_name:  element.SelectAttrValue("operStatsPolicyName", ""),
			Oper_vcon:  element.SelectAttrValue("operVcon", ""),
			Order:  element.SelectAttrValue("order", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Pers_bind:  element.SelectAttrValue("persBind", ""),
			Pers_bind_clear:  element.SelectAttrValue("persBindClear", ""),
			Pin_to_group_name:  element.SelectAttrValue("pinToGroupName", ""),
			Qos_policy_name:  element.SelectAttrValue("qosPolicyName", ""),
			Redundancy_pair_type:  element.SelectAttrValue("redundancyPairType", ""),
			Redundancy_peer:  element.SelectAttrValue("redundancyPeer", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Stats_policy_name:  element.SelectAttrValue("statsPolicyName", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Type:  element.SelectAttrValue("type", ""),
		},
    
	}
}