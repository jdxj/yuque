package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jdxj/yuque/modules"
)

func (c *Client) ListDoc(namespace string) ([]*modules.DocSerializer, error) {
	path := fmt.Sprintf(APIPath+APIDocs, namespace)
	req, err := c.newHTTPRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	reader, err := c.do(req)
	if err != nil {
		return nil, err
	}

	docListed := &DocListed{}
	decoder := json.NewDecoder(reader)
	return docListed.Data, decoder.Decode(docListed)
}

func (c *Client) CreateDoc(namespace string, docReq *CreateDocRequest) (*modules.DocDetailSerializer, error) {
	path := fmt.Sprintf(APIPath+APIDocs, namespace)
	data, err := json.Marshal(docReq)
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

	dc := &DocCreated{}
	decoder := json.NewDecoder(reader)
	return dc.Data, decoder.Decode(dc)

}
