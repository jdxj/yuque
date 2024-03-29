package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/jdxj/yuque/utils"
)

func TestClient_ListUserRepositories(t *testing.T) {
	c := newClient()

	lrp := &ListReposParams{
		Type:   "Design",
		Offset: 0,
	}

	repos, err := c.ListUserRepositories("jdxj", lrp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("len: %d\n", len(repos))
	for _, repo := range repos {
		fmt.Printf("%#v\n", *repo)
		fmt.Printf("\t%#v\n", repo.User)
	}
}

func TestStringsJoin(t *testing.T) {
	s := []string{}
	r := strings.Join(s, "&")
	fmt.Printf("1: %s\n", r)

	s = []string{"1"}
	r = strings.Join(s, "&")
	fmt.Printf("2: %s\n", r)

	s = []string{"1", "2"}
	r = strings.Join(s, "&")
	fmt.Printf("3: %s\n", r)
}

func TestClient_ListGroupRepositories(t *testing.T) {
	c := newClient()

	repos, err := c.ListGroupRepositories("book-academy", nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("len: %d\n", len(repos))
	for _, repo := range repos {
		fmt.Printf("%#v\n", *repo)
		fmt.Printf("\t%#v\n", repo.User)
	}
}

func TestURLValues(t *testing.T) {
	v := url.Values{
		"type":   []string{"1"},
		"offset": []string{"8"},
	}
	fmt.Printf("1: %s\n", v.Encode())

	v = url.Values{
		"offset": []string{},
	}
	fmt.Printf("2: %s\n", v.Encode())

	v = url.Values{}
	v.Set("offset", "")
	fmt.Printf("3: %s\n", v.Encode())

	v = url.Values{
		"offset": []string{""},
	}
	fmt.Printf("2: %s\n", v.Encode())

	var v2 url.Values

	v2.Add("abc", "def")
	fmt.Printf("v2: %s\n", v2.Encode())
}

func TestClient_CreateUserRepository(t *testing.T) {
	c := newClient()
	maxI, maxJ := 500, 10

	for i := 0; i < maxI; i++ {
		crp := &CreateRepoParams{
			Name:        "autoCreate",
			Slug:        utils.GenRandString(8),
			Description: "test",
			Public:      0,
			Type:        Book,
		}

		bds, err := c.CreateUserRepository("jdxj", crp)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		fmt.Printf("%#v\n", bds)

		for j := 0; j < maxJ; j++ {
			crd := &CreateRepoDocParams{
				Title:  "abc",
				Slug:   utils.GenRandString(8),
				Public: 1,
				Format: Markdown,
				Body:   "- abc",
			}
			dds, err := c.CreateRepositoryDoc(strconv.Itoa(bds.ID), crd)
			if err != nil {
				t.Fatalf("CreateRepositoryDoc: %s\n", err)
			}
			fmt.Printf("%#v\n", dds)
		}
	}

}

func TestClient_GetRepositoryDetail(t *testing.T) {
	c := newClient()

	grd := &GetRepoDetailParams{
		Type: Book,
	}

	bds, err := c.GetRepositoryDetail("jdxj/azdm6s", grd)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%#v\n", bds.BookSerializer)
}

func TestUpdateRepoParams_Reader(t *testing.T) {
	c := newClient()

	urp := &UpdateRepoParams{
		Name:        "ialli",
		Slug:        "ii87ny",
		Toc:         "",
		Description: "haha",
		Public:      Open,
	}

	bds, err := c.UpdateRepository("jdxj/ii8a22", urp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", bds.BookSerializer)
}

func TestClient_DeleteRepository(t *testing.T) {
	c := newClient()

	bds, err := c.DeleteRepository("jdxj/ii87ny")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", bds.BookSerializer)
}

func TestClient_GetRepositoryToc(t *testing.T) {
	c := newClient()

	tocs, err := c.GetRepositoryToc("jdxj/azdm6s")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, toc := range tocs {
		fmt.Printf("%#v\n", *toc)
		fmt.Printf("%v\n", toc.UnmarshalDocID())
	}
}
