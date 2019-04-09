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

	macPool := models.NewMacPoolPool(fmt.Sprintf("mac-pool-%s", "macpool_with_client"), "org-root", "testMacPool", models.MacPoolPoolAttributes{})
	err := c.Save(macPool)
	if err != nil {
		t.Error(err)
	}

}
