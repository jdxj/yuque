package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// ListUserJoinedGroup 获取某个用户的加入的组织列表
func (c *Client) ListUserJoinedGroup(login string) ([]*UserSerializer, error) {
	path := fmt.Sprintf(APIUsersGroups, login)
	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	uss := new([]*UserSerializer)
	return *uss, json.Unmarshal(data, uss)
}

// ListPublicGroups 获取公开组织列表
func (c *Client) ListPublicGroups() ([]*UserSerializer, error) {
	req := c.newReqGet(APIGroups)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	uss := new([]*UserSerializer)
	return *uss, json.Unmarshal(data, uss)
}

type CreateGroupParams struct {
	Name        string `json:"name"`
	Login       string `json:"login"` // 用户个人路径
	Description string `json:"description"`
}

func (cgp *CreateGroupParams) Reader() io.Reader {
	data, _ := json.Marshal(cgp)
	return bytes.NewReader(data)
}

// CreateGroup 创建 Group
func (c *Client) CreateGroup(cgp *CreateGroupParams) (*UserSerializer, error) {
	req := c.newReqPost(APIGroups, cgp.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	us := new(UserSerializer)
	return us, json.Unmarshal(data, us)
}

// GetGroupDetail 获取单个组织的详细信息
func (c *Client) GetGroupDetail(id string) (*UserSerializer, error) {
	path := fmt.Sprintf(APIGroupsDetail, id)
	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	us := new(UserSerializer)
	return us, json.Unmarshal(data, us)
}

type UpdateGroupDetailParams struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Description string `json:"description"`
}

func (ugd *UpdateGroupDetailParams) Reader() io.Reader {
	data, _ := json.Marshal(ugd)
	return bytes.NewReader(data)
}

// UpdateGroupDetail 更新单个组织的详细信息
func (c *Client) UpdateGroupDetail(id string, ugd *UpdateGroupDetailParams) (*UserSerializer, error) {
	path := fmt.Sprintf(APIGroupsDetail, id)
	req := c.newReqPut(path, ugd.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	us := new(UserSerializer)
	return us, json.Unmarshal(data, us)
}

// DeleteGroup 删除组织
func (c *Client) DeleteGroup(id string) (*UserSerializer, error) {
	path := fmt.Sprintf(APIGroupsDetail, id)
	req := c.newReqDelete(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	us := new(UserSerializer)
	return us, json.Unmarshal(data, us)
}

// ListGroupUsers 获取组织成员信息
func (c *Client) ListGroupUsers(id string) ([]*GroupUserSerializer, error) {
	path := fmt.Sprintf(APIGroupsUsers, id)
	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var guss []*GroupUserSerializer
	return guss, json.Unmarshal(data, &guss)
}

type UpdateGroupUsersParams struct {
	Role int `json:"role"` // 0: 管理员, 1: 普通成员
}

func (ugu *UpdateGroupUsersParams) Reader() io.Reader {
	data, _ := json.Marshal(ugu)
	return bytes.NewReader(data)
}

// UpdateGroupUsers 增加或更新组织成员
// notes: 未测试
func (c *Client) UpdateGroupUser(groupID, userID string, ugu *UpdateGroupUsersParams) (*GroupUserSerializer, error) {
	path := fmt.Sprintf(APIGroupsUsersUpdate, groupID, userID)
	req := c.newReqPut(path, ugu.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	gus := new(GroupUserSerializer)
	return gus, json.Unmarshal(data, gus)
}

// DeleteGroupUser 删除组织成员
// notes: 未测试
func (c *Client) DeleteGroupUser(groupID, userID string) (*GroupUserSerializer, error) {
	path := fmt.Sprintf(APIGroupsUsersUpdate, groupID, userID)
	req := c.newReqDelete(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	gus := new(GroupUserSerializer)
	return gus, json.Unmarshal(data, gus)
}
