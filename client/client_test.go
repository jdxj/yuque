package client

import (
	"fmt"
	"testing"

	"github.com/jdxj/yuque/config"
)

func newClient() *Client {
	err := config.Init("../config/config.yaml")
	if err != nil {
		panic(err)
	}

	return New(config.Token())
}

func TestClient_ListUserJoinedGroup(t *testing.T) {
	c := newClient()
	res, err := c.ListUserJoinedGroup("barretlee")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, user := range res {
		fmt.Printf("%#v\n", *user)
	}
}

func TestClient_ListPublicGroups(t *testing.T) {
	c := newClient()

	res, err := c.ListPublicGroups()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, user := range res {
		fmt.Printf("%#v\n", *user)
	}
}
