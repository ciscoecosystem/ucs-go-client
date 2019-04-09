package client

import (
	"fmt"
	"time"

	"github.com/beevik/etree"
)

type Auth struct {
	Token         string
	RefreshPeriod int64
	offset        int64
	Expiry        time.Time
}

func (client *Client) InjectAuthenticationHeader(body *etree.Document) (*etree.Document, error) {
	if client.Password != "" {
		if client.AuthToken == nil {
			err := client.Authenticate()
			if err != nil {
				return nil, err
			}
		} else if !client.AuthToken.IsValid() {
			client.refreshToken(client.AuthToken.Token)
		}
		//return req, nil
	} else {
		return body, fmt.Errorf("Please provide password")
	}
	childs := body.ChildElements()[0]
	childs.CreateAttr("cookie", client.AuthToken.Token)
	body.ChildElements()[0].CreateAttr("cookie", "testing")
	return body, nil
}

func (au *Auth) IsValid() bool {
	if au.Token != "" && au.Expiry.Unix() > au.estimateExpireTime() {
		return true
	}
	return false
}

func (t *Auth) CalculateExpiry(refreshTime int64) {
	t.Expiry = time.Unix((time.Now().Unix() + refreshTime), 0)
}

func (t *Auth) estimateExpireTime() int64 {
	return time.Now().Unix() + t.RefreshPeriod - t.offset
}

func (c *Client) Authenticate() error {
	method := "POST"
	payLoad := fmt.Sprintf(LoginPayLoad, c.UserName, c.Password)
	doc := etree.NewDocument()
	err := doc.ReadFromString(payLoad)
	if err != nil {
		return err
	}

	req, err := c.MakeRestRequest(method, doc, false)
	if err != nil {
		return err
	}
	doc, _, err = c.Do(req)
	if err != nil {
		return err
	}

	if c.AuthToken == nil {
		c.AuthToken = &Auth{}
	}
	c.AuthToken.offset = 2
	c.AuthToken.Token = doc.SelectAttrValue("outCookie", "")
	c.AuthToken.RefreshPeriod, err = StrtoInt(doc.SelectAttrValue("outRefreshPeriod", "0"), 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing Refresh period.")
	}
	c.AuthToken.CalculateExpiry(c.AuthToken.RefreshPeriod)
	return nil

}
func (c *Client) refreshToken(oldtoken string) error {
	method := "POST"
	payLoad := fmt.Sprintf(RefreshPayLoad, c.UserName, c.Password, c.AuthToken.Token)
	doc := etree.NewDocument()
	err := doc.ReadFromString(payLoad)
	if err != nil {
		return err
	}

	req, err := c.MakeRestRequest(method, doc, false)
	if err != nil {
		return err
	}
	doc, _, err = c.Do(req)
	if err != nil {
		return err
	}
	c.AuthToken.Token = doc.SelectAttrValue("outCookie", "")
	c.AuthToken.RefreshPeriod, err = StrtoInt(doc.SelectAttrValue("outRefreshPeriod", "0"), 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing Refresh period.")
	}
	c.AuthToken.CalculateExpiry(c.AuthToken.RefreshPeriod)
	return nil
}
