package models

import (
	"fmt"

	"github.com/beevik/etree"
)

func GetMoElement(doc *etree.Document, rootClass, className string) (*etree.Element, error) {
	rootEle := doc.SelectElement(rootClass)
	if rootEle == nil {
		return nil, fmt.Errorf("Unable to load configConfMo root element")
	}
	outConfEle := rootEle.SelectElement("outConfig")
	if outConfEle == nil {
		return nil, fmt.Errorf("Unable to load outConfig Element")
	}
	classEle := outConfEle.SelectElement(className)
	fmt.Printf("\n\n\n ********* class ele %+v", classEle)
	if classEle == nil {
		return nil, fmt.Errorf("Unable to load %s element", className)
	}

	return classEle, nil

}

func GetChildEle(rootClass string, childList []*etree.Element) *etree.Element {

	for _, ele := range childList {
		if ele.Tag == rootClass {
			return ele
		}
	}

	return nil
}
