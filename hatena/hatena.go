package hatena

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	dg "github.com/tj/go-debug"
)

const (
	acceptHeader = "application/json"
	version      = 1
)

var (
	hd             = dg.Debug("hatena")
	defaultBaseURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d", version)
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	BookMarks *BookmarksService
	Entries   *EntriesService
	Tags      *TagsService
	Users     *UsersService
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	hd("rel %v", rel)

	u := c.BaseURL.ResolveReference(rel)
	hd("u %v", u)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", acceptHeader)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	hd("req %v", req)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// NewHatena is initialize package
func NewHatena(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	cli := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	cli.Entries = &EntriesService{client: cli}
	cli.BookMarks = &BookmarksService{client: cli}
	cli.Tags = &TagsService{client: cli}
	cli.Users = &UsersService{client: cli}

	return cli
}
