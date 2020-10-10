package client

import (
	"fmt"
	"testing"
)

func TestClient_CreateGroup(t *testing.T) {
	c := newClient()

	cgp := &CreateGroupParams{
		Name:        "test create group",
		Login:       "89ui8",
		Description: "test",
	}
	us, err := c.CreateGroup(cgp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", us)
}

func TestClient_GetGroupDetail(t *testing.T) {
	c := newClient()

	us, err := c.GetGroupDetail("89ui8")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", us)
}

func TestClient_UpdateGroupDetail(t *testing.T) {
	c := newClient()

	ugd := &UpdateGroupDetailParams{
		Name:        "test update group detail",
		Login:       "",
		Description: "test",
	}
	us, err := c.UpdateGroupDetail("89ui8", ugd)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", us)
}

func TestClient_DeleteGroup(t *testing.T) {
	c := newClient()

	us, err := c.DeleteGroup("89ui8")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", us)
}

func TestClient_ListGroupUsers(t *testing.T) {
	c := newClient()

	guss, err := c.ListGroupUsers("book-academy")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("len: %d\n", len(guss))
	for _, gus := range guss {
		fmt.Printf("%#v\n", gus)
	}
}
