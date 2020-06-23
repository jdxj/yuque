package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func newClient() *Client {
	c, err := NewClientToken("")
	if err != nil {
		panic(fmt.Sprintf("Create client fail: %s\n", err))
	}
	return c
}

func TestClient_User(t *testing.T) {
	c := newClient()
	user, err := c.User()
	if err != nil {
		t.Fatalf("Get user fail: %s\n", err)
	}
	t.Logf("%#v\n", user)
}

func TestClient_ListOwnUserRepositories(t *testing.T) {
	c := newClient()
	repos, err := c.ListOwnUserRepositories()
	if err != nil {
		t.Fatalf("Get repositories faild: %s\n", err)
	}

	for _, repo := range repos {
		t.Logf("%#v\n", repo)
	}
}

func TestClient_DeleteRepository(t *testing.T) {
	c := newClient()

	if bds, err := c.DeleteRepository("jdxj/xkcfoq"); err != nil {
		t.Fatalf("delete repository faild: %s\n", err)
	} else {
		t.Logf("%#v\n", bds)
	}
}

func TestClient_ListGroupRepositories(t *testing.T) {
	c := newClient()

	repos, err := c.ListGroupRepositories("yuque")
	if err != nil {
		t.Fatalf("Group repositories faild: %s\n", err)
	}

	for _, repo := range repos {
		t.Logf("%#v\n", repo)
	}
}

func TestClient_ListDoc(t *testing.T) {
	c := newClient()

	docs, err := c.ListDoc("yuque/help")
	if err != nil {
		t.Fatalf("List docs faild: %s\n", err)
	}

	for _, doc := range docs {
		t.Logf("%#v\n", doc.LastEditor)
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

	t.Logf("%s", data)
}

func TestClient_CreateGroupRepository(t *testing.T) {
	c := newClient()

	repoReq := NewCreateRepositoryRequestSlug("lefaiii", "", Book, Open)
	bds, err := c.CreateGroupRepository("itsi1d", repoReq)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	t.Logf("%#v\n", bds)
}

func TestClient_CreateDoc(t *testing.T) {
	c := newClient()

	docReq := NewCreateDocRequestSlug("haha2", "this is a body", Open, Markdown)
	dds, err := c.CreateDoc("jdxj/mlakfd", docReq)
	if err != nil {
		t.Fatalf("%s", err)
	}

	t.Logf("%#v\n", dds)
}

func TestClient_CreateDocAmount(t *testing.T) {
	c := newClient()

	c.CreateDocAmount(10)
}

func TestClient_DeleteAutoCreate(t *testing.T) {
	c := newClient()

	c.DeleteAutoCreate()
}

func TestClient_CreateRepoDoc(t *testing.T) {
	c := newClient()

	c.CreateRepoDoc(50)
}

func TestClient_Users(t *testing.T) {
	c := newClient()

	user, err := c.Users("yubo555")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("use name: %#v\n", user)
}
