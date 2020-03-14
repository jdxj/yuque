package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

func (c *Client) User() (*modules.UserSerializer, error) {
	if c.user != nil {
		return c.user, nil
	}

	path := fmt.Sprintf("%s%s", APIPath, APIUser)
	req, err := c.newHTTPRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	reader, err := c.do(req)
	if err != nil {
		return nil, err
	}

	ur := &UserResponse{}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(ur); err != nil {
		return nil, err
	}

	c.user = ur.Data
	return ur.Data, nil
}
