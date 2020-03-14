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
		Public:      int(public),
		Type:        string(typ),
	}
	return crr
}

type CreateRepositoryRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Public      int    `json:"public"`
	Type        string `json:"type"`
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
