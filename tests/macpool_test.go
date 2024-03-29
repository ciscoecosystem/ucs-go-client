package test

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
)

func GetTestClient() *client.Client {
	return client.GetClient("http://10.0.8.224/", "ucspe", client.Insecure(true), client.Password("ucspe"))
}

func TestMacPoolCreate(t *testing.T) {
	c := GetTestClient()

	macPool := models.NewMacpoolPool(fmt.Sprintf("mac-pool-%s", "macpool_with_client"), "org-root", "testMacPool", models.MacpoolPoolAttributes{})
	err := c.Save(macPool)
	if err != nil {
		t.Error(err)
	}

}

func TestMacPoolDelete(t *testing.T) {
	c := GetTestClient()
	macPool := models.NewMacpoolPool(fmt.Sprintf("mac-pool-%s", "macpool_with_client"), "org-root", "testMacPool", models.MacpoolPoolAttributes{})
	_, err := c.Delete(macPool)
	if err != nil {
		t.Error(err)
	}
}

func TestMacPoolRead(t *testing.T) {
	c := GetTestClient()
	doc, err := c.Get("org-root/mac-pool-macpool_with_client")
	if err != nil {
		t.Error(err)
	}

	macPool := models.MacpoolPoolFromDoc(doc, "configResolveDn")

	fmt.Println(macPool)

	t.Error(fmt.Errorf("dummy error"))
}

func TestMacPoolDeleteByDn(t *testing.T) {
	c := GetTestClient()
	err := c.DeleteByDn("org-root/mac-pool-macpool_with_client", "macpoolPool")
	if err != nil {
		t.Error(err)
	}
}
