package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const VnicIpV4MgmtPooledAddrClassName = "vnicIpV4MgmtPooledAddr"

type VnicIpV4MgmtPooledAddr struct {
    BaseAttributes
    VnicIpV4MgmtPooledAddrAttributes
}

type VnicIpV4MgmtPooledAddrAttributes struct{
    Addr string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Def_gw string `xml:",omitempty"`
    Name string `xml:",omitempty"`
    Oper_name string `xml:",omitempty"`
    Prim_dns string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Sec_dns string `xml:",omitempty"`
    Subnet string `xml:",omitempty"`
    

}

func NewVnicIpV4MgmtPooledAddr(vnicIpV4MgmtPooledAddrRn,parentDn,description string, vnicIpV4MgmtPooledAddrAttributes VnicIpV4MgmtPooledAddrAttributes) *VnicIpV4MgmtPooledAddr { 
    dn := fmt.Sprintf("%s/%s", parentDn, vnicIpV4MgmtPooledAddrRn)
	return &VnicIpV4MgmtPooledAddr{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         VnicIpV4MgmtPooledAddrClassName,
			Status:            "created, modified",
			Description:       description,
		},

		VnicIpV4MgmtPooledAddrAttributes: vnicIpV4MgmtPooledAddrAttributes,
	}
}

func (vnicIpV4MgmtPooledAddr *VnicIpV4MgmtPooledAddr) ToMap() (map[string]string, error) {
	vnicIpV4MgmtPooledAddrMap, err := vnicIpV4MgmtPooledAddr.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    vnicIpV4MgmtPooledAddrMap["addr"] = vnicIpV4MgmtPooledAddr.Addr
    vnicIpV4MgmtPooledAddrMap["childAction"] = vnicIpV4MgmtPooledAddr.Child_action
    vnicIpV4MgmtPooledAddrMap["defGw"] = vnicIpV4MgmtPooledAddr.Def_gw
    vnicIpV4MgmtPooledAddrMap["name"] = vnicIpV4MgmtPooledAddr.Name
    vnicIpV4MgmtPooledAddrMap["operName"] = vnicIpV4MgmtPooledAddr.Oper_name
    vnicIpV4MgmtPooledAddrMap["primDns"] = vnicIpV4MgmtPooledAddr.Prim_dns
    vnicIpV4MgmtPooledAddrMap["sacl"] = vnicIpV4MgmtPooledAddr.Sacl
    vnicIpV4MgmtPooledAddrMap["secDns"] = vnicIpV4MgmtPooledAddr.Sec_dns
    vnicIpV4MgmtPooledAddrMap["subnet"] = vnicIpV4MgmtPooledAddr.Subnet
    
	return vnicIpV4MgmtPooledAddrMap, nil

}

func VnicIpV4MgmtPooledAddrFromDoc(doc *etree.Document, rootClass string) *VnicIpV4MgmtPooledAddr {
	element, err := GetMoElement(doc, rootClass, VnicIpV4MgmtPooledAddrClassName)
	if err != nil {
		return nil
	}
	return &VnicIpV4MgmtPooledAddr{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         VnicIpV4MgmtPooledAddrClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		VnicIpV4MgmtPooledAddrAttributes: VnicIpV4MgmtPooledAddrAttributes{

        
			Addr:  element.SelectAttrValue("addr", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Def_gw:  element.SelectAttrValue("defGw", ""),
			Name:  element.SelectAttrValue("name", ""),
			Oper_name:  element.SelectAttrValue("operName", ""),
			Prim_dns:  element.SelectAttrValue("primDns", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Sec_dns:  element.SelectAttrValue("secDns", ""),
			Subnet:  element.SelectAttrValue("subnet", ""),
		},
    
	}
}