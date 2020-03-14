package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

const (
	APIPath        = "https://www.yuque.com/api/v2"
	APIUser        = "/user"
	APIUserRepos   = "/users/%s/repos"
	APIGroupsRepos = "/groups/%s/repos"
	APIReposDel    = "/repos/%s"
	APIDocs        = "/repos/%s/docs"

	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"

	SlugLength = 6

	RepositoryNamePrefix = "AutoCreate"
)

type Typ string

const (
	Book   Typ = "Book"
	Design     = "Design"
	All        = "All"
)

type Public int

const (
	Private Public = iota
	Intranet
	Open
)

func NewClientToken(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("token can not be empty")
	}

	c := &Client{
		httpClient: &http.Client{},
		token:      token,
	}
	return c, nil
}

type Client struct {
	httpClient *http.Client
	token      string
	user       *modules.UserSerializer
}

func (c *Client) newHTTPRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("X-Auth-Token", c.token)
	return req, nil
}
