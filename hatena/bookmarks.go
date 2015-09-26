package hatena

import (
	"errors"
	"fmt"
	"net/url"

	dg "github.com/tj/go-debug"
)

const (
	bookmarkURL = "http://api.b.st-hatena.com/entry.count?url=%s"
)

var (
	bd = dg.Debug("boolmark")
)

type BookmarksService struct {
	client *Client
}

func (s *BookmarksService) Count(urlStr string) (count int, err error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return
	}

	reqURL := fmt.Sprintf(bookmarkURL, u)
	bd("reqURL %v", reqURL)

	req, err := s.client.NewRequest("GET", reqURL, nil)
	if err != nil {
		return
	}
	bd("req %+v", req)

	res, err := s.client.Do(req, &count)
	bd("req %+v", res)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		return count, errors.New("not.ok")
	}

	return
}

func (s *BookmarksService) Add(url, comment string, tags []string) (err error) {
	return
}

func (s *BookmarksService) Delete(url string) (err error) {
	return
}
