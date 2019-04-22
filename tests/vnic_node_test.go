package test

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/ucs-go-client/models"
)

func TestVnicNodeCreate(t *testing.T) {
	c := GetTestClient()
	node := models.NewVnicIScsiNode(fmt.Sprintf("iscsi-node"), "org-root/ls-tf_ls_server", "", models.VnicIScsiNodeAttributes{})
	err := c.Save(node)
	if err != nil {
		t.Error(err)
	}
}

func TestVnicNodeRead(t *testing.T) {
	c := GetTestClient()
	doc, err := c.Get("org-root/ls-tf_ls_server/iscsi-node")
	if err != nil {
		t.Error(err)
	}
	node := models.VnicIScsiNodeFromDoc(doc, "configResolveDn")

	fmt.Println(node.DistinguishedName)
	t.Error(err)

}
