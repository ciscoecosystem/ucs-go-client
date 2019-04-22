package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicConnDefClassName = "vnicConnDef"

type VnicConnDef struct {
    BaseAttributes
    VnicConnDefAttributes
}

type VnicConnDefAttributes struct{
    Child_action string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Lan_conn_policy_name string `xml:",omitempty"`
    Oper_lan_conn_policy_name string `xml:",omitempty"`
    Oper_san_conn_policy_name string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    San_conn_policy_name string `xml:",omitempty"`
    

}

func NewVnicConnDef(vnicConnDefRn,parentDn,description string, vnicConnDefAttributes VnicConnDefAttributes) *VnicConnDef { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicConnDefRn)
	return &VnicConnDef{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicConnDefClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicConnDefAttributes: vnicConnDefAttributes,
	}
}

func (vnicConnDef *VnicConnDef) ToMap() (map[string]string, error) {
	vnicConnDefMap, err := vnicConnDef.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicConnDefMap["childAction"] = vnicConnDef.Child_action
    vnicConnDefMap["fltAggr"] = vnicConnDef.Flt_aggr
    vnicConnDefMap["lanConnPolicyName"] = vnicConnDef.Lan_conn_policy_name
    vnicConnDefMap["operLanConnPolicyName"] = vnicConnDef.Oper_lan_conn_policy_name
    vnicConnDefMap["operSanConnPolicyName"] = vnicConnDef.Oper_san_conn_policy_name
    vnicConnDefMap["propAcl"] = vnicConnDef.Prop_acl
    vnicConnDefMap["sacl"] = vnicConnDef.Sacl
    vnicConnDefMap["sanConnPolicyName"] = vnicConnDef.San_conn_policy_name
    
	return vnicConnDefMap, nil

}

func VnicConnDefFromDoc(doc *etree.Document, rootClass string) *VnicConnDef {
	element, err := GetMoElement(doc, rootClass, VnicConnDefClassName)
	if err != nil {
		return nil
	}
	return &VnicConnDef{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicConnDefClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicConnDefAttributes: VnicConnDefAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Lan_conn_policy_name:  element.SelectAttrValue("lanConnPolicyName", ""),
			Oper_lan_conn_policy_name:  element.SelectAttrValue("operLanConnPolicyName", ""),
			Oper_san_conn_policy_name:  element.SelectAttrValue("operSanConnPolicyName", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			San_conn_policy_name:  element.SelectAttrValue("sanConnPolicyName", ""),
		},
    
	}
}