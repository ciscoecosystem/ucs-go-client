package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicFcNodeClassName = "vnicFcNode"

type VnicFcNode struct {
    BaseAttributes
    VnicFcNodeAttributes
}

type VnicFcNodeAttributes struct{
    Addr string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Ident_pool_name string `xml:",omitempty"`
    Max_derivable_wwpn string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    

}

func NewVnicFcNode(vnicFcNodeRn,parentDn,description string, vnicFcNodeAttributes VnicFcNodeAttributes) *VnicFcNode { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicFcNodeRn)
	return &VnicFcNode{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicFcNodeClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicFcNodeAttributes: vnicFcNodeAttributes,
	}
}

func (vnicFcNode *VnicFcNode) ToMap() (map[string]string, error) {
	vnicFcNodeMap, err := vnicFcNode.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicFcNodeMap["addr"] = vnicFcNode.Addr
    vnicFcNodeMap["childAction"] = vnicFcNode.Child_action
    vnicFcNodeMap["fltAggr"] = vnicFcNode.Flt_aggr
    vnicFcNodeMap["identPoolName"] = vnicFcNode.Ident_pool_name
    vnicFcNodeMap["maxDerivableWWPN"] = vnicFcNode.Max_derivable_wwpn
    vnicFcNodeMap["operIdentPoolName"] = vnicFcNode.Oper_ident_pool_name
    vnicFcNodeMap["owner"] = vnicFcNode.Owner
    vnicFcNodeMap["sacl"] = vnicFcNode.Sacl
    
	return vnicFcNodeMap, nil

}

func VnicFcNodeFromDoc(doc *etree.Document, rootClass string) *VnicFcNode {
	element, err := GetMoElement(doc, rootClass, VnicFcNodeClassName)
	if err != nil {
		return nil
	}
	return &VnicFcNode{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicFcNodeClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicFcNodeAttributes: VnicFcNodeAttributes{

        
			Addr:  element.SelectAttrValue("addr", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Ident_pool_name:  element.SelectAttrValue("identPoolName", ""),
			Max_derivable_wwpn:  element.SelectAttrValue("maxDerivableWWPN", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
		},
    
	}
}