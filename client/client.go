package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

var ErrNoFound = fmt.Errorf("http resp err: not found")

const (
	APIPath        = "https://www.yuque.com/api/v2"
	APIUsers       = "/users/%s"
	APIUser        = "/user"
	APIUserRepos   = "/users/%s/repos"
	APIGroupsRepos = "/groups/%s/repos"
	APIRepos       = "/repos/%s"
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
	Open
	SpaceMember
	SpaceOpen
	RepositoryMember
)

type Format string

const (
	Markdown Format = "markdown"
	Lake            = "lake"
)

func NewClientToken(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("token can not be empty")
	}

	c := &Client{
		httpClient:    &http.Client{},
		token:         token,
		namespaceTask: make(chan string),
	}
	return c, nil
}

type Client struct {
	httpClient *http.Client
	token      string
	user       *modules.UserSerializer

	// status
	// todo: 是否考虑并发?
	xRateLimitRemaining string

	// play
	namespaceTask chan string
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

func (c *Client) do(req *http.Request) (io.Reader, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	c.xRateLimitRemaining = resp.Header.Get("X-RateLimit-Remaining")

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return nil, ErrNoFound
		}
		return nil, fmt.Errorf("%s", data)
	}
	return bytes.NewBuffer(data), nil
}

func (c *Client) XRateLimitRemaining() string {
	return c.xRateLimitRemaining
}
