package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

// ListRepositoryDocs 获取一个仓库的文档列表
func (c *Client) ListRepositoryDocs(id string) ([]*DocSerializer, error) {
	path := fmt.Sprintf(APIReposDocs, id)
	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var dss []*DocSerializer
	return dss, json.Unmarshal(data, &dss)
}

type GetRepoDocDetailParams struct {
	Raw int `json:"raw"`
}

func (grd *GetRepoDocDetailParams) String() string {
	values := url.Values{}
	if grd.Raw != 0 {
		values.Set("raw", strconv.Itoa(grd.Raw))
	}
	return values.Encode()
}

// GetRepositoryDocDetail 获取单篇文档的详细信息
func (c *Client) GetRepositoryDocDetail(id, slug string, grd *GetRepoDocDetailParams) (*DocDetailSerializer, error) {
	path := fmt.Sprintf(APIReposDocsDetail, id, slug)
	paramsKV := grd.String()
	if len(paramsKV) != 0 {
		path = fmt.Sprintf("%s?%s", path, paramsKV)
	}
	req := c.newReqGet(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	dds := new(DocDetailSerializer)
	return dds, json.Unmarshal(data, dds)
}

type CreateRepoDocParams struct {
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Public int    `json:"public"`
	Format string `json:"format"`
	Body   string `json:"body"`
}

func (crd *CreateRepoDocParams) Reader() io.Reader {
	data, _ := json.Marshal(crd)
	return bytes.NewReader(data)
}

// CreateRepositoryDoc 创建文档
func (c *Client) CreateRepositoryDoc(id string, crd *CreateRepoDocParams) (*DocDetailSerializer, error) {
	path := fmt.Sprintf(APIReposDocs, id)
	req := c.newReqPost(path, crd.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	dds := new(DocDetailSerializer)
	return dds, json.Unmarshal(data, dds)
}

type UpdateRepoDocParams struct {
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Public int    `json:"public"`
	Body   string `json:"body"`
}

func (urd *UpdateRepoDocParams) Reader() io.Reader {
	data, _ := json.Marshal(urd)
	return bytes.NewReader(data)
}

// UpdateRepositoryDoc 更新文档
func (c *Client) UpdateRepositoryDoc(namespace string, docID int, urd *UpdateRepoDocParams) (*DocDetailSerializer, error) {
	path := fmt.Sprintf(APIReposDocsDetail, namespace, strconv.Itoa(docID))
	req := c.newReqPut(path, urd.Reader())
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	dds := new(DocDetailSerializer)
	return dds, json.Unmarshal(data, dds)
}

// DeleteRepositoryDoc 删除文档
func (c *Client) DeleteRepositoryDoc(namespace string, docID int) (*DocDetailSerializer, error) {
	path := fmt.Sprintf(APIReposDocsDetail, namespace, strconv.Itoa(docID))
	req := c.newReqDelete(path)
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	dds := new(DocDetailSerializer)
	return dds, json.Unmarshal(data, dds)
}
