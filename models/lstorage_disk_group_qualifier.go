package models

import (
	"fmt"

	"github.com/beevik/etree"

)

const LstorageDiskGroupQualifierClassName = "lstorageDiskGroupQualifier"

type LstorageDiskGroupQualifier struct {
    BaseAttributes
    LstorageDiskGroupQualifierAttributes
}

type LstorageDiskGroupQualifierAttributes struct{
    Child_action string `xml:",omitempty"`
    Drive_type string `xml:",omitempty"`
    Min_drive_size string `xml:",omitempty"`
    Num_ded_hot_spares string `xml:",omitempty"`
    Num_drives string `xml:",omitempty"`
    Num_glob_hot_spares string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Use_jbod_disks string `xml:",omitempty"`
    Use_remaining_disks string `xml:",omitempty"`
    

}

func NewLstorageDiskGroupQualifier(lstorageDiskGroupQualifierRn,parentDn,description string, lstorageDiskGroupQualifierAttributes LstorageDiskGroupQualifierAttributes) *LstorageDiskGroupQualifier { 
    dn := fmt.Sprintf("%s/%s", parentDn, lstorageDiskGroupQualifierRn)
	return &LstorageDiskGroupQualifier{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LstorageDiskGroupQualifierClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LstorageDiskGroupQualifierAttributes: lstorageDiskGroupQualifierAttributes,
	}
}

func (lstorageDiskGroupQualifier *LstorageDiskGroupQualifier) ToMap() (map[string]string, error) {
	lstorageDiskGroupQualifierMap, err := lstorageDiskGroupQualifier.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lstorageDiskGroupQualifierMap["childAction"] = lstorageDiskGroupQualifier.Child_action
    lstorageDiskGroupQualifierMap["driveType"] = lstorageDiskGroupQualifier.Drive_type
    lstorageDiskGroupQualifierMap["minDriveSize"] = lstorageDiskGroupQualifier.Min_drive_size
    lstorageDiskGroupQualifierMap["numDedHotSpares"] = lstorageDiskGroupQualifier.Num_ded_hot_spares
    lstorageDiskGroupQualifierMap["numDrives"] = lstorageDiskGroupQualifier.Num_drives
    lstorageDiskGroupQualifierMap["numGlobHotSpares"] = lstorageDiskGroupQualifier.Num_glob_hot_spares
    lstorageDiskGroupQualifierMap["sacl"] = lstorageDiskGroupQualifier.Sacl
    lstorageDiskGroupQualifierMap["useJbodDisks"] = lstorageDiskGroupQualifier.Use_jbod_disks
    lstorageDiskGroupQualifierMap["useRemainingDisks"] = lstorageDiskGroupQualifier.Use_remaining_disks
    
	return lstorageDiskGroupQualifierMap, nil

}

func LstorageDiskGroupQualifierFromDoc(doc *etree.Document, rootClass string) *LstorageDiskGroupQualifier {
	element, err := GetMoElement(doc, rootClass, LstorageDiskGroupQualifierClassName)
	if err != nil {
		return nil
	}
	return &LstorageDiskGroupQualifier{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LstorageDiskGroupQualifierClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LstorageDiskGroupQualifierAttributes: LstorageDiskGroupQualifierAttributes{

        
			Child_action:  element.SelectAttrValue("childAction", ""),
			Drive_type:  element.SelectAttrValue("driveType", ""),
			Min_drive_size:  element.SelectAttrValue("minDriveSize", ""),
			Num_ded_hot_spares:  element.SelectAttrValue("numDedHotSpares", ""),
			Num_drives:  element.SelectAttrValue("numDrives", ""),
			Num_glob_hot_spares:  element.SelectAttrValue("numGlobHotSpares", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Use_jbod_disks:  element.SelectAttrValue("useJbodDisks", ""),
			Use_remaining_disks:  element.SelectAttrValue("useRemainingDisks", ""),
		},
    
	}
}