package client

import (
	"fmt"
	"testing"
)

//func newClient() *Client {
//	c, err := NewClientToken("your token")
//	if err != nil {
//		panic(fmt.Sprintf("Create client fail: %s\n", err))
//	}
//	return c
//}

func TestClient_User(t *testing.T) {
	c := newClient()
	user, err := c.User()
	if err != nil {
		t.Fatalf("Get user fail: %s\n", err)
	}
	fmt.Printf("%#v\n", user)
}

func TestClient_ListOwnUserRepositories(t *testing.T) {
	c := newClient()
	repos, err := c.ListOwnUserRepositories()
	if err != nil {
		t.Fatalf("Get repositories faild: %s\n", err)
	}

	for _, repo := range repos {
		fmt.Printf("%#v\n", repo)
	}
}

func TestClient_CreateRepository(t *testing.T) {
	c := newClient()

	req := NewCreateRepositoryRequestSlug(RepositoryNamePrefix, "", Book, Open)
	bds, err := c.CreateRepository(req)
	if err != nil {
		t.Fatalf("create repository faild: %s\n", err)
	}

	fmt.Printf("%#v\n", bds)
}

func TestClient_DeleteRepository(t *testing.T) {
	c := newClient()

	if bds, err := c.DeleteUserRepository("jdxj/xkcfoq"); err != nil {
		t.Fatalf("delete repository faild: %s\n", err)
	} else {
		fmt.Printf("%#v\n", bds)
	}
}

func TestClient_ListGroupRepositories(t *testing.T) {
	c := newClient()

	repos, err := c.ListGroupRepositories("yuque")
	if err != nil {
		t.Fatalf("Group repositories faild: %s\n", err)
	}

	for _, repo := range repos {
		fmt.Printf("%#v\n", repo)
	}
}

func TestClient_ListDoc(t *testing.T) {
	c := newClient()

	docs, err := c.ListDoc("yuque/help")
	if err != nil {
		t.Fatalf("List docs faild: %s\n", err)
	}

	for _, doc := range docs {
		fmt.Printf("%#v\n", doc.LastEditor)
	}
}
