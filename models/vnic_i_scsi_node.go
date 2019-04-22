package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicIScsiNodeClassName = "vnicIScsiNode"

type VnicIScsiNode struct {
    BaseAttributes
    VnicIScsiNodeAttributes
}

type VnicIScsiNodeAttributes struct{
    Child_action string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Init_name_suffix string `xml:",omitempty"`
    Initiator_name string `xml:",omitempty"`
    Initiator_policy_name string `xml:",omitempty"`
    Iqn_ident_pool_name string `xml:",omitempty"`
    Oper_initiator_policy_name string `xml:",omitempty"`
    Oper_iqn_ident_pool_name string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewVnicIScsiNode(vnicIScsiNodeRn,parentDn,description string, vnicIScsiNodeAttributes VnicIScsiNodeAttributes) *VnicIScsiNode { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicIScsiNodeRn)
	return &VnicIScsiNode{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicIScsiNodeClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicIScsiNodeAttributes: vnicIScsiNodeAttributes,
	}
}

func (vnicIScsiNode *VnicIScsiNode) ToMap() (map[string]string, error) {
	vnicIScsiNodeMap, err := vnicIScsiNode.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicIScsiNodeMap["childAction"] = vnicIScsiNode.Child_action
    vnicIScsiNodeMap["fltAggr"] = vnicIScsiNode.Flt_aggr
    vnicIScsiNodeMap["initNameSuffix"] = vnicIScsiNode.Init_name_suffix
    vnicIScsiNodeMap["initiatorName"] = vnicIScsiNode.Initiator_name
    vnicIScsiNodeMap["initiatorPolicyName"] = vnicIScsiNode.Initiator_policy_name
    vnicIScsiNodeMap["iqnIdentPoolName"] = vnicIScsiNode.Iqn_ident_pool_name
    vnicIScsiNodeMap["operInitiatorPolicyName"] = vnicIScsiNode.Oper_initiator_policy_name
    vnicIScsiNodeMap["operIqnIdentPoolName"] = vnicIScsiNode.Oper_iqn_ident_pool_name
    vnicIScsiNodeMap["owner"] = vnicIScsiNode.Owner
    vnicIScsiNodeMap["propAcl"] = vnicIScsiNode.Prop_acl
    vnicIScsiNodeMap["sacl"] = vnicIScsiNode.Sacl
    
	return vnicIScsiNodeMap, nil

}

func VnicIScsiNodeFromDoc(doc *etree.Document, rootClass string) *VnicIScsiNode {
	element, err := GetMoElement(doc, rootClass, VnicIScsiNodeClassName)
	if err != nil {
		return nil
	}
	return &VnicIScsiNode{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicIScsiNodeClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicIScsiNodeAttributes: VnicIScsiNodeAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Init_name_suffix:  element.SelectAttrValue("initNameSuffix", ""),
			Initiator_name:  element.SelectAttrValue("initiatorName", ""),
			Initiator_policy_name:  element.SelectAttrValue("initiatorPolicyName", ""),
			Iqn_ident_pool_name:  element.SelectAttrValue("iqnIdentPoolName", ""),
			Oper_initiator_policy_name:  element.SelectAttrValue("operInitiatorPolicyName", ""),
			Oper_iqn_ident_pool_name:  element.SelectAttrValue("operIqnIdentPoolName", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}