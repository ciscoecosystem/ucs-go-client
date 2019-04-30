package models

import (


	"github.com/beevik/etree"

)

const FabricVsanClassName = "fabricVsan"

type FabricVsan struct {
    BaseAttributes
    FabricVsanAttributes
}

type FabricVsanAttributes struct{
    Child_action string `xml:",omitempty"`
    Config_overlap string `xml:",omitempty"`
    Default_zoning string `xml:",omitempty"`
    Ep_dn string `xml:",omitempty"`
    Fc_zone_sharing_mode string `xml:",omitempty"`
    Fcoe_vlan string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    R_global string `xml:",omitempty"`
    Fabric_vsan_id string `xml:",omitempty"`
    If_role string `xml:",omitempty"`
    If_type string `xml:",omitempty"`
    Local string `xml:",omitempty"`
    Locale string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Peer_dn string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Switch_id string `xml:",omitempty"`
    Transport string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    Zoning_state string `xml:",omitempty"`
    

}

func NewFabricVsan(fabricVsanRn,description string, fabricVsanAttributes FabricVsanAttributes) *FabricVsan {
    dn := fabricVsanRn
	return &FabricVsan{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         FabricVsanClassName,
			Status:            "created, modified",
			Description:       description,
		},

		FabricVsanAttributes: fabricVsanAttributes,
	}
}

func (fabricVsan *FabricVsan) ToMap() (map[string]string, error) {
	fabricVsanMap, err := fabricVsan.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    fabricVsanMap["childAction"] = fabricVsan.Child_action
    fabricVsanMap["configOverlap"] = fabricVsan.Config_overlap
    fabricVsanMap["defaultZoning"] = fabricVsan.Default_zoning
    fabricVsanMap["epDn"] = fabricVsan.Ep_dn
    fabricVsanMap["fcZoneSharingMode"] = fabricVsan.Fc_zone_sharing_mode
    fabricVsanMap["fcoeVlan"] = fabricVsan.Fcoe_vlan
    fabricVsanMap["fltAggr"] = fabricVsan.Flt_aggr
    fabricVsanMap["global"] = fabricVsan.R_global
    fabricVsanMap["id"] = fabricVsan.Fabric_vsan_id
    fabricVsanMap["ifRole"] = fabricVsan.If_role
    fabricVsanMap["ifType"] = fabricVsan.If_type
    fabricVsanMap["local"] = fabricVsan.Local
    fabricVsanMap["locale"] = fabricVsan.Locale
    fabricVsanMap["operState"] = fabricVsan.Oper_state
    fabricVsanMap["peerDn"] = fabricVsan.Peer_dn
    fabricVsanMap["policyOwner"] = fabricVsan.Policy_owner
    fabricVsanMap["sacl"] = fabricVsan.Sacl
    fabricVsanMap["switchId"] = fabricVsan.Switch_id
    fabricVsanMap["transport"] = fabricVsan.Transport
    fabricVsanMap["type"] = fabricVsan.Type
    fabricVsanMap["zoningState"] = fabricVsan.Zoning_state
    
	return fabricVsanMap, nil

}

func FabricVsanFromDoc(doc *etree.Document, rootClass string) *FabricVsan {
	element, err := GetMoElement(doc, rootClass, FabricVsanClassName)
	if err != nil {
		return nil
	}
	return &FabricVsan{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         FabricVsanClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		FabricVsanAttributes: FabricVsanAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_overlap:  element.SelectAttrValue("configOverlap", ""),
			Default_zoning:  element.SelectAttrValue("defaultZoning", ""),
			Ep_dn:  element.SelectAttrValue("epDn", ""),
			Fc_zone_sharing_mode:  element.SelectAttrValue("fcZoneSharingMode", ""),
			Fcoe_vlan:  element.SelectAttrValue("fcoeVlan", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			R_global:  element.SelectAttrValue("global", ""),
			Fabric_vsan_id:  element.SelectAttrValue("id", ""),
			If_role:  element.SelectAttrValue("ifRole", ""),
			If_type:  element.SelectAttrValue("ifType", ""),
			Local:  element.SelectAttrValue("local", ""),
			Locale:  element.SelectAttrValue("locale", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Peer_dn:  element.SelectAttrValue("peerDn", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Switch_id:  element.SelectAttrValue("switchId", ""),
			Transport:  element.SelectAttrValue("transport", ""),
			Type:  element.SelectAttrValue("type", ""),
			Zoning_state:  element.SelectAttrValue("zoningState", ""),
		},
    
	}
}