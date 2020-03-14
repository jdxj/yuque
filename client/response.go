package client

import (
	"strings"

	"github.com/jdxj/yuque/modules"
	"github.com/jdxj/yuque/utils"
)

func NewCreateRepositoryRequestSlug(name, description string, typ Typ, public Public) *CreateRepositoryRequest {
	slug := utils.GenRandString(SlugLength)
	slug = strings.ToLower(slug)

	crr := &CreateRepositoryRequest{
		Name:        name,
		Slug:        slug,
		Description: description,
		Public:      public,
		Type:        typ,
	}
	return crr
}

type CreateRepositoryRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Public      Public `json:"public"`
	Type        Typ    `json:"type"`
}

type UserResponse struct {
	Data *modules.UserSerializer `json:"data"`
}

type RepositoriesListed struct {
	Data []*modules.BookSerializer `json:"data"`
}

type RepositoryCreated struct {
	Data *modules.BookDetailSerializer `json:"data"`
}

type RepositoryDeleted struct {
	Data *modules.BookDetailSerializer `json:"data"`
}

type DocListed struct {
	Data []*modules.DocSerializer `json:"data"`
}

func NewCreateDocRequestSlug(title, body string, public Public, format Format) *CreateDocRequest {
	slug := utils.GenRandString(SlugLength)
	slug = strings.ToLower(slug)

	cdr := &CreateDocRequest{
		Title:  title,
		Slug:   slug,
		Public: public,
		Format: format,
		Body:   body,
	}
	return cdr
}

type CreateDocRequest struct {
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Public Public `json:"public"`
	Format Format `json:"format"`
	Body   string `json:"body"`
}

type DocCreated struct {
	Data *modules.DocDetailSerializer `json:"data"`
}
