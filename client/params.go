package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

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

type CreateGroupParams struct {
	Name        string `json:"name"`
	Login       string `json:"login"` // 用户个人路径
	Description string `json:"description"`
}

func (cgp *CreateGroupParams) Reader() io.Reader {
	data, _ := json.Marshal(cgp)
	return bytes.NewReader(data)
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

type UpdateGroupUsersParams struct {
	Role int `json:"role"` // 0: 管理员, 1: 普通成员
}

func (ugu *UpdateGroupUsersParams) Reader() io.Reader {
	data, _ := json.Marshal(ugu)
	return bytes.NewReader(data)
}

type CreateRepoParams struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Public      int    `json:"public"`
	Type        string `json:"type"`
}

func (crp *CreateRepoParams) Reader() io.Reader {
	data, _ := json.Marshal(crp)
	return bytes.NewReader(data)
}

type ListReposParams struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
}

func (lrp *ListReposParams) String() string {
	values := url.Values{}
	if lrp.Type != "" {
		values.Set("type", lrp.Type)
	}
	if lrp.Offset != 0 {
		values.Set("offset", strconv.Itoa(lrp.Offset))
	}
	return values.Encode()
}

type GetRepoDetailParams struct {
	Type string `json:"type"`
}

func (grd *GetRepoDetailParams) String() string {
	values := url.Values{}
	if grd.Type != "" {
		values.Set("type", grd.Type)
	}
	return values.Encode()
}

type UpdateRepoParams struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Toc         string `json:"toc"`
	Description string `json:"description"`
	Public      int    `json:"public"`
}

func (urp *UpdateRepoParams) Reader() io.Reader {
	data, _ := json.Marshal(urp)
	return bytes.NewReader(data)
}
