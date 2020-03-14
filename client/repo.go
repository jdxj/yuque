package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

func (c *Client) CreateUserRepository(repoReq *CreateRepositoryRequest) (*modules.BookDetailSerializer, error) {
	if c.user == nil {
		if _, err := c.User(); err != nil {
			return nil, err
		}
	}

	path := fmt.Sprintf(APIPath+APIUserRepos, c.user.Login)
	return c.createRepository(path, repoReq)
}

func (c *Client) CreateGroupRepository(group string, repoReq *CreateRepositoryRequest) (*modules.BookDetailSerializer, error) {
	path := fmt.Sprintf(APIPath+APIGroupsRepos, group)
	return c.createRepository(path, repoReq)
}

func (c *Client) createRepository(path string, repoReq *CreateRepositoryRequest) (*modules.BookDetailSerializer, error) {
	data, err := json.Marshal(repoReq)
	if err != nil {
		return nil, err
	}

	req, err := c.newHTTPRequest(http.MethodPost, path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	reader, err := c.do(req)
	if err != nil {
		return nil, err
	}

	rr := &RepositoryCreated{}
	decoder := json.NewDecoder(reader)
	return rr.Data, decoder.Decode(rr)
}

func (c *Client) DeleteUserRepository(namespace string) (*modules.BookDetailSerializer, error) {
	path := fmt.Sprintf(APIPath+APIReposDel, namespace)
	req, err := c.newHTTPRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	reader, err := c.do(req)
	if err != nil {
		return nil, err
	}

	dur := &RepositoryDeleted{}
	decoder := json.NewDecoder(reader)
	return dur.Data, decoder.Decode(dur)
}

func (c *Client) ListOwnUserRepositories() ([]*modules.BookSerializer, error) {
	if c.user == nil {
		if _, err := c.User(); err != nil {
			return nil, err
		}
	}

	return c.ListUserRepositories(c.user.Login)
}

func (c *Client) ListUserRepositories(login string) ([]*modules.BookSerializer, error) {
	path := fmt.Sprintf(APIPath+APIUserRepos, login)
	return c.listRepositories(path)
}

func (c *Client) ListGroupRepositories(group string) ([]*modules.BookSerializer, error) {
	path := fmt.Sprintf(APIPath+APIGroupsRepos, group)
	return c.listRepositories(path)
}

func (c *Client) listRepositories(path string) ([]*modules.BookSerializer, error) {
	req, err := c.newHTTPRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	reader, err := c.do(req)
	if err != nil {
		return nil, err
	}

	reposList := &RepositoriesListed{}
	decoder := json.NewDecoder(reader)
	return reposList.Data, decoder.Decode(reposList)
}
