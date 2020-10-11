package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateUserRepository 往自己下面创建知识库
func (c *Client) CreateUserRepository(id string, crp *CreateRepoParams) (*BookDetailSerializer, error) {
	path := fmt.Sprintf(APIUsersRepos, id)
	return c.createRepository(path, crp)
}

// CreateGroupRepository 往团队创建知识库
// notes: 未测试
func (c *Client) CreateGroupRepository(id string, crp *CreateRepoParams) (*BookDetailSerializer, error) {
	path := fmt.Sprintf(APIGroupsRepos, id)
	return c.createRepository(path, crp)
}

func (c *Client) createRepository(path string, crp *CreateRepoParams) (*BookDetailSerializer, error) {
	req := c.newReqPost(path, crp.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	bds := new(BookDetailSerializer)
	return bds, json.Unmarshal(data, bds)
}

// DeleteRepository 删除知识库
func (c *Client) DeleteRepository(id string) (*BookDetailSerializer, error) {
	path := fmt.Sprintf(APIRepos, id)
	req := c.newReqDelete(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	bds := new(BookDetailSerializer)
	return bds, json.Unmarshal(data, bds)
}

// ListUserRepositories 获取某个用户的知识库列表
func (c *Client) ListUserRepositories(id string, lrp *ListReposParams) ([]*BookSerializer, error) {
	path := fmt.Sprintf(APIUsersRepos, id)
	return c.listRepositories(path, lrp)
}

// ListGroupRepositories 获取某个团队的知识库列表
func (c *Client) ListGroupRepositories(id string, lrp *ListReposParams) ([]*BookSerializer, error) {
	path := fmt.Sprintf(APIGroupsRepos, id)
	return c.listRepositories(path, lrp)
}

func (c *Client) listRepositories(path string, lrp *ListReposParams) ([]*BookSerializer, error) {
	paramsKV := lrp.String()
	if len(paramsKV) != 0 {
		path = fmt.Sprintf("%s?%s", path, paramsKV)
	}

	req, err := c.newHTTPRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var repos []*BookSerializer
	return repos, json.Unmarshal(data, &repos)
}

// GetRepositoryDetail 获取知识库详情
func (c *Client) GetRepositoryDetail(id string, grd *GetRepoDetailParams) (*BookDetailSerializer, error) {
	path := fmt.Sprintf(APIRepos, id)
	paramsKV := grd.String()
	if len(paramsKV) != 0 {
		path = fmt.Sprintf("%s?%s", path, paramsKV)
	}

	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	bds := new(BookDetailSerializer)
	return bds, json.Unmarshal(data, bds)
}

// UpdateRepository 更新知识库信息
func (c *Client) UpdateRepository(id string, urp *UpdateRepoParams) (*BookDetailSerializer, error) {
	path := fmt.Sprintf(APIRepos, id)
	req := c.newReqPut(path, urp.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	bds := new(BookDetailSerializer)
	return bds, json.Unmarshal(data, bds)
}

// GetRepositoryToc 获取一个知识库的目录结构
func (c *Client) GetRepositoryToc(id string) ([]*Toc, error) {
	path := fmt.Sprintf(APIReposToc, id)
	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var tocs []*Toc
	return tocs, json.Unmarshal(data, &tocs)
}
