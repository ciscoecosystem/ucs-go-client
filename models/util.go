package models

import (
	"fmt"

	"github.com/beevik/etree"
)

func GetMoElement(doc *etree.Document, className string) (*etree.Element, error) {
	rootEle := doc.SelectElement("configConfMo")
	if rootEle != nil {
		return nil, fmt.Errorf("Unable to load configConfMo root element")
	}
	outConfEle := rootEle.SelectElement("outConfig")
	if outConfEle != nil {
		return nil, fmt.Errorf("Unable to load outConfig Element")
	}
	classEle := outConfEle.SelectElement(className)
	if classEle != nil {
		return nil, fmt.Errorf("Unable to load %s element", className)
	}

	return classEle, nil

}
