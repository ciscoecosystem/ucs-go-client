package models

import (
	"fmt"

	"github.com/beevik/etree"
)

const MacPoolPoolClassName = "macpoolPool"

type MacPoolPool struct {
	BaseAttributes
	MacPoolPoolAttributes
}

type MacPoolPoolAttributes struct {
	Assigned        string `xml:",omitempty"`
	AssignmentOrder string `xml:",omitempty"`
	IntID           string `xml:",omitempty"`
	Name            string `xml:",omitempty"`
	PolicyLevel     string `xml:",omitempty"`
	PolicyOwner     string `xml:",omitempty"`
	Size            string `xml:",omitempty"`
}

func NewMacPoolPool(macPoolPoolRn, parentDn, description string, macPoolPoolAttributes MacPoolPoolAttributes) *MacPoolPool {
	dn := fmt.Sprintf("%s/%s", parentDn, macPoolPoolRn)
	return &MacPoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         MacPoolPoolClassName,
			Status:            "created, modified",
			Description:       description,
		},

		MacPoolPoolAttributes: macPoolPoolAttributes,
	}
}

func (macPoolPool *MacPoolPool) ToMap() (map[string]string, error) {
	macPoolPoolMap, err := macPoolPool.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
	macPoolPoolMap["name"] = macPoolPool.Name
	macPoolPoolMap["assigned"] = macPoolPool.Assigned
	macPoolPoolMap["assignmentOrder"] = macPoolPool.AssignmentOrder
	macPoolPoolMap["intId"] = macPoolPool.IntID
	macPoolPoolMap["policyLevel"] = macPoolPool.PolicyLevel
	macPoolPoolMap["policyOwner"] = macPoolPool.PolicyOwner
	macPoolPoolMap["size"] = macPoolPool.Size

	return macPoolPoolMap, nil

}

func MacPoolPoolFromDoc(doc *etree.Document) *MacPoolPool {
	element, err := GetMoElement(doc, MacPoolPoolClassName)
	if err != nil {
		return nil
	}
	return &MacPoolPool{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         MacPoolPoolClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},

		MacPoolPoolAttributes: MacPoolPoolAttributes{
			Name:            element.SelectAttrValue("name", ""),
			Assigned:        element.SelectAttrValue("assigned", ""),
			AssignmentOrder: element.SelectAttrValue("assignmentOrder", ""),
			IntID:           element.SelectAttrValue("intId", ""),
			Size:            element.SelectAttrValue("size", ""),
			PolicyLevel:     element.SelectAttrValue("policyLevel", ""),
			PolicyOwner:     element.SelectAttrValue("policyOwner", ""),
		},
	}
}
