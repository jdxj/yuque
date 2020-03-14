package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

func (c *Client) CreateRepository(repoReq *CreateRepositoryRequest) (*modules.BookDetailSerializer, error) {
	if repoReq == nil {
		return nil, fmt.Errorf("invalid repository request param")
	}

	if c.user == nil {
		if _, err := c.User(); err != nil {
			return nil, err
		}
	}

	data, err := json.Marshal(repoReq)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf(APIPath+APIUserRepos, c.user.Login)

	req, err := c.newHTTPRequest(http.MethodPost, path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s", data)
	}

	rr := &RepositoryCreated{}
	decoder := json.NewDecoder(resp.Body)
	return rr.Data, decoder.Decode(rr)
}

func (c *Client) DeleteUserRepository(namespace string) (*modules.BookDetailSerializer, error) {
	path := fmt.Sprintf(APIPath+APIReposDel, namespace)
	req, err := c.newHTTPRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s", data)
	}

	dur := &RepositoryDeleted{}
	decoder := json.NewDecoder(resp.Body)
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

func (c *Client) ListGroupRepositories(login string) ([]*modules.BookSerializer, error) {
	path := fmt.Sprintf(APIPath+APIGroupsRepos, login)
	return c.listRepositories(path)
}

func (c *Client) listRepositories(path string) ([]*modules.BookSerializer, error) {
	req, err := c.newHTTPRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s", data)
	}

	reposList := &RepositoriesListed{}
	decoder := json.NewDecoder(resp.Body)
	return reposList.Data, decoder.Decode(reposList)
}
