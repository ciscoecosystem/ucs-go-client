package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const LstorageDasScsiLunClassName = "lstorageDasScsiLun"

type LstorageDasScsiLun struct {
    BaseAttributes
    LstorageDasScsiLunAttributes
}

type LstorageDasScsiLunAttributes struct{
    Admin_state string `xml:",omitempty"`
    Auto_deploy string `xml:",omitempty"`
    Boot_dev string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Config_qualifier string `xml:",omitempty"`
    Config_state string `xml:",omitempty"`
    Deferred_naming string `xml:",omitempty"`
    Expand_to_avail string `xml:",omitempty"`
    Fractional_size string `xml:",omitempty"`
    Local_disk_policy_name string `xml:",omitempty"`
    Lun_dn string `xml:",omitempty"`
    Lun_map_type string `xml:",omitempty"`
    Oper_local_disk_policy_name string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Order string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Size string `xml:",omitempty"`
    Storage_class string `xml:",omitempty"`
    

}

func NewLstorageDasScsiLun(lstorageDasScsiLunRn,parentDn,description string, lstorageDasScsiLunAttributes LstorageDasScsiLunAttributes) *LstorageDasScsiLun { 
    dn := fmt.Sprintf("%s/%s", parentDn, lstorageDasScsiLunRn)
	return &LstorageDasScsiLun{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LstorageDasScsiLunClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LstorageDasScsiLunAttributes: lstorageDasScsiLunAttributes,
	}
}

func (lstorageDasScsiLun *LstorageDasScsiLun) ToMap() (map[string]string, error) {
	lstorageDasScsiLunMap, err := lstorageDasScsiLun.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lstorageDasScsiLunMap["adminState"] = lstorageDasScsiLun.Admin_state
    lstorageDasScsiLunMap["autoDeploy"] = lstorageDasScsiLun.Auto_deploy
    lstorageDasScsiLunMap["bootDev"] = lstorageDasScsiLun.Boot_dev
    lstorageDasScsiLunMap["childAction"] = lstorageDasScsiLun.Child_action
    lstorageDasScsiLunMap["configQualifier"] = lstorageDasScsiLun.Config_qualifier
    lstorageDasScsiLunMap["configState"] = lstorageDasScsiLun.Config_state
    lstorageDasScsiLunMap["deferredNaming"] = lstorageDasScsiLun.Deferred_naming
    lstorageDasScsiLunMap["expandToAvail"] = lstorageDasScsiLun.Expand_to_avail
    lstorageDasScsiLunMap["fractionalSize"] = lstorageDasScsiLun.Fractional_size
    lstorageDasScsiLunMap["localDiskPolicyName"] = lstorageDasScsiLun.Local_disk_policy_name
    lstorageDasScsiLunMap["lunDn"] = lstorageDasScsiLun.Lun_dn
    lstorageDasScsiLunMap["lunMapType"] = lstorageDasScsiLun.Lun_map_type
    lstorageDasScsiLunMap["operLocalDiskPolicyName"] = lstorageDasScsiLun.Oper_local_disk_policy_name
    lstorageDasScsiLunMap["operState"] = lstorageDasScsiLun.Oper_state
    lstorageDasScsiLunMap["order"] = lstorageDasScsiLun.Order
    lstorageDasScsiLunMap["sacl"] = lstorageDasScsiLun.Sacl
    lstorageDasScsiLunMap["size"] = lstorageDasScsiLun.Size
    lstorageDasScsiLunMap["storageClass"] = lstorageDasScsiLun.Storage_class
    
	return lstorageDasScsiLunMap, nil

}

func LstorageDasScsiLunFromDoc(doc *etree.Document, rootClass string) *LstorageDasScsiLun {
	element, err := GetMoElement(doc, rootClass, LstorageDasScsiLunClassName)
	if err != nil {
		return nil
	}
	return &LstorageDasScsiLun{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LstorageDasScsiLunClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LstorageDasScsiLunAttributes: LstorageDasScsiLunAttributes{

        
			Admin_state:  element.SelectAttrValue("adminState", ""),
			Auto_deploy:  element.SelectAttrValue("autoDeploy", ""),
			Boot_dev:  element.SelectAttrValue("bootDev", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_qualifier:  element.SelectAttrValue("configQualifier", ""),
			Config_state:  element.SelectAttrValue("configState", ""),
			Deferred_naming:  element.SelectAttrValue("deferredNaming", ""),
			Expand_to_avail:  element.SelectAttrValue("expandToAvail", ""),
			Fractional_size:  element.SelectAttrValue("fractionalSize", ""),
			Local_disk_policy_name:  element.SelectAttrValue("localDiskPolicyName", ""),
			Lun_dn:  element.SelectAttrValue("lunDn", ""),
			Lun_map_type:  element.SelectAttrValue("lunMapType", ""),
			Oper_local_disk_policy_name:  element.SelectAttrValue("operLocalDiskPolicyName", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Order:  element.SelectAttrValue("order", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Size:  element.SelectAttrValue("size", ""),
			Storage_class:  element.SelectAttrValue("storageClass", ""),
		},
    
	}
}