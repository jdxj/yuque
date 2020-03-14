package client

import (
	"encoding/json"
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

type Tmp struct {
	Name   string `json:"name"`
	Format Format `json:"format"`
	Typ    Typ    `json:"typ"`
	Public Public `json:"public"`
}

func TestJsonType(t *testing.T) {
	tmp := &Tmp{
		Name:   "fff",
		Format: Markdown,
		Typ:    Book,
		Public: Open,
	}

	data, err := json.Marshal(tmp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%s", data)
}

func TestClient_CreateGroupRepository(t *testing.T) {
	c := newClient()

	repoReq := NewCreateRepositoryRequestSlug("lefaiii", "", Book, Open)
	bds, err := c.CreateGroupRepository("itsi1d", repoReq)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%#v\n", bds)
}

func TestClient_CreateDoc(t *testing.T) {
	c := newClient()

	docReq := NewCreateDocRequestSlug("haha2", "this is a body", Intranet, Markdown)
	dds, err := c.CreateDoc("jdxj/mlakfd", docReq)
	if err != nil {
		t.Fatalf("%s", err)
	}

	fmt.Printf("%#v\n", dds)
}
