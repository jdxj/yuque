package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

func (c *Client) ListDoc(namespace string) ([]*modules.DocSerializer, error) {
	path := fmt.Sprintf(APIPath+APIDocs, namespace)
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

	docListed := &DocListed{}
	decoder := json.NewDecoder(resp.Body)
	return docListed.Data, decoder.Decode(docListed)
}
