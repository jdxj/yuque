package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

func (c *Client) User() (*modules.UserSerializer, error) {
	path := fmt.Sprintf("%s%s", APIPath, APIUser)
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

	ur := &UserResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(ur); err != nil {
		return nil, err
	}

	c.user = ur.Data
	return ur.Data, nil
}
