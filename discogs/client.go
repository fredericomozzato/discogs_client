package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	url        string
	userAgent  string
}

func NewClient(url string, timeout int, userAgent string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: time.Duration(timeout) * time.Second},
		url:        url,
		userAgent:  userAgent,
	}
}

func NewDefaultClient(userAgent string) *Client {
	defaultURL := "https://api.discogs.com"
	defaultTimeout := 30
	return NewClient(defaultURL, defaultTimeout, userAgent)
}

func (c *Client) GetRelease(ctx context.Context, id int) (*Release, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/releases/%d", c.url, id),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	req.Header.Add("User-Agent", c.userAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("response error: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %d", res.StatusCode)
	}

	var release Release
	err = json.NewDecoder(res.Body).Decode(&release)
	if err != nil {
		return nil, fmt.Errorf("json unmarshalling: %w", err)
	}

	return &release, nil
}
