package client

import "testing"

func TestClient_AuthenticatedUser(t *testing.T) {
	c := newClient()
	user, err := c.AuthenticatedUser()
	if err != nil {
		t.Fatalf("Get user fail: %s\n", err)
	}
	t.Logf("%#v\n", user)
}

func TestClient_IndividualUser(t *testing.T) {
	c := newClient()

	user, err := c.IndividualUser("barretlee")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("use name: %#v\n", user)
}
