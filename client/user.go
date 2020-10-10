package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthenticatedUser 获取认证的用户的个人信息
func (c *Client) AuthenticatedUser() (*UserSerializer, error) {
	req, err := c.newHTTPRequest(http.MethodGet, APIUser, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.do(req)
	if err != nil {
		return nil, err
	}
	us := new(UserSerializer)
	return us, json.Unmarshal(data, us)
}

// IndividualUser 获取单个用户信息
// id: 1: 用户编号 (数字), 2: 用户个人路径 (字符串)
func (c *Client) IndividualUser(id string) (*UserSerializer, error) {
	path := fmt.Sprintf(APIUsers, id)
	req, err := c.newHTTPRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.do(req)
	if err != nil {
		return nil, err
	}
	us := new(UserSerializer)
	return us, json.Unmarshal(data, us)
}
