package client

import (
	"fmt"
	"testing"
)

func TestClient_ListRepositoryDocs(t *testing.T) {
	c := newClient()

	dss, err := c.ListRepositoryDocs("jdxj/azdm6s")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, ds := range dss {
		fmt.Printf("%#v\n", *ds)
	}
}

func TestClient_GetRepositoryDocsDetail(t *testing.T) {
	c := newClient()

	grd := &GetRepoDocDetailParams{
		Raw: 0,
	}
	dds, err := c.GetRepositoryDocDetail("angel-vwrye/gv2pb8", "gi9u90", grd)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%#v\n", *dds)
}

func TestClient_CreateRepositoryDoc(t *testing.T) {
	c := newClient()

	crd := &CreateRepoDocParams{
		Title:  "abc",
		Slug:   "88u90",
		Public: 1,
		Format: Markdown,
		Body:   "- abc",
	}
	dds, err := c.CreateRepositoryDoc("jdxj/qkfjzq", crd)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *dds)
}

func TestClient_UpdateRepositoryDoc(t *testing.T) {
	c := newClient()

	urd := &UpdateRepoDocParams{
		Title:  "iuaa",
		Slug:   "00p1i",
		Public: 1,
		Body:   "- 456",
	}
	dds, err := c.UpdateRepositoryDoc("jdxj/qkfjzq", 14168983, urd)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *dds)
}

func TestClient_DeleteRepositoryDoc(t *testing.T) {
	c := newClient()

	dds, err := c.DeleteRepositoryDoc("jdxj/qkfjzq", 14168983)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *dds)
}
