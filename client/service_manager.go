package client

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/ciscoecosystem/ucs-go-client/models"
)

func (c *Client) Save(obj models.Model) error {

	doc, _, err := c.PrepareModel(obj, "configConfMo", "inConfig", "", "")
	if err != nil {
		return err
	}
	req, err := c.MakeRestRequest("POST", doc, true)
	if err != nil {
		return err
	}

	respDoc, _, err := c.Do(req)

	strDoc, err := respDoc.WriteToString()
	fmt.Println("\n\n\n*********** response %s **********\n\n\n", strDoc)
	return CheckForErrors(respDoc, "configConfMo")

}

func (c *Client) Get(dn string) (*etree.Document, error) {
	doc, _, err := c.PrepareModel(nil, "configResolveDn", "", dn, "")
	if err != nil {
		return nil, err
	}

	req, err := c.MakeRestRequest("POST", doc, true)
	if err != nil {
		return nil, err
	}

	respDoc, _, err := c.Do(req)
	return respDoc, CheckForErrors(respDoc, "configResolveDn")
}

func (c *Client) Delete(obj models.Model) (*etree.Document, error) {

	doc, _, err := c.PrepareModel(obj, "configConfMo", "inConfig", "", "deleted")
	if err != nil {
		return nil, err
	}
	req, err := c.MakeRestRequest("POST", doc, true)
	if err != nil {
		return nil, err
	}

	respDoc, _, err := c.Do(req)

	return nil, CheckForErrors(respDoc, "configConfMo")

}

func (c *Client) PrepareModel(obj models.Model, rootClass, subClass, confDn, status string) (*etree.Document, string, error) {
	var moEle *etree.Element
	var classname string
	if obj != nil {
		mapObject, err := obj.ToMap()
		if err != nil {
			return nil, "", err
		}

		if status != "" {
			mapObject["status"] = status
		}

		classname := mapObject["classname"]
		moEle = etree.NewElement(classname)
		delete(mapObject, "classname")

		for key, value := range mapObject {
			if value != "" {
				moEle.CreateAttr(key, value)
			}
		}
	}
	fmt.Println("*****here****\n\n\n")
	doc := CreateHierarchy(rootClass, subClass, confDn, moEle)
	strDoc, _ := doc.WriteToString()
	fmt.Println(strDoc)
	return doc, classname, nil

}

func CreateHierarchy(rootClass, subClass, confDn string, moEle *etree.Element) *etree.Document {
	doc := etree.NewDocument()

	rootEle := doc.CreateElement(rootClass)
	rootEle.CreateAttr("dn", confDn)
	if subClass != "" {
		subEle := rootEle.CreateElement(subClass)
		subEle.AddChild(moEle)
	}
	return doc
}

func CheckForErrors(doc *etree.Document, rootClass string) error {
	rootEle := doc.SelectElement(rootClass)
	if rootEle != nil {
		return fmt.Errorf("Unable to load %s element", rootClass)
	}

	errCodeAttr := rootEle.SelectAttr("errorCode")

	if errCodeAttr != nil {
		return fmt.Errorf("Error occured %s Reason %s", errCodeAttr.Key, errCodeAttr.Value)
	}

	return nil
}
