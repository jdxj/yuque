package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrNoFound      = errors.New("http resp err: not found")
	ErrInvalidToken = errors.New("invalid token")
)

const (
	bufLimit = 8 * 1 << 10 // 8KB
)

var (
	pool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte, bufLimit)
			return bytes.NewBuffer(buf)
		},
	}
)

func New(token string) *Client {
	jar, _ := cookiejar.New(nil)
	hc := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	c := &Client{
		token:      token,
		httpClient: hc,
	}
	return c
}

type Client struct {
	token      string
	httpClient *http.Client

	xRateLimitLimit     int32
	xRateLimitRemaining int32
}

func (c *Client) newHTTPRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", DefaultUserAgent)
	req.Header.Set("X-Auth-Token", c.token)
	return req, nil
}

func (c *Client) newReqGet(path string) *http.Request {
	req, _ := c.newHTTPRequest(http.MethodGet, path, nil)
	return req
}

func (c *Client) newReqPost(path string, body io.Reader) *http.Request {
	req, _ := c.newHTTPRequest(http.MethodPost, path, body)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (c *Client) newReqPut(path string, body io.Reader) *http.Request {
	req, _ := c.newHTTPRequest(http.MethodPut, path, body)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (c *Client) newReqDelete(path string) *http.Request {
	req, _ := c.newHTTPRequest(http.MethodDelete, path, nil)
	return req
}

func (c *Client) do(req *http.Request) (json.RawMessage, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	c.updateLimit(resp)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("low-level data: %s\n", data)

	if resp.StatusCode == 200 {
		mResp := new(Response)
		return mResp.Data, json.Unmarshal(data, mResp)
	}

	err, ok := ErrMsg[resp.StatusCode]
	if ok {
		return nil, err
	}
	return nil, fmt.Errorf("%s: %d", ErrCodeNotDefine, resp.StatusCode)
}

func (c *Client) XRateLimitLimit() int {
	return int(atomic.LoadInt32(&c.xRateLimitLimit))
}

func (c *Client) XRateLimitRemaining() int {
	return int(atomic.LoadInt32(&c.xRateLimitRemaining))
}

func (c *Client) updateLimit(resp *http.Response) {
	limit, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Limit"))
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))

	atomic.StoreInt32(&c.xRateLimitLimit, int32(limit))
	atomic.StoreInt32(&c.xRateLimitRemaining, int32(remaining))
}
