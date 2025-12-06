package discogs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// TODO: caller must provide a User-Agent
const UserAgent = "DiscogsClient github.com/fredericomozzato/discogs_client"

type Client struct {
	httpClient *http.Client
	url        string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		url:        "https://api.discogs.com",
	}
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

	req.Header.Add("User-Agent", UserAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("response error: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("unexpected HTTP status: %d\n", res.StatusCode))
	}

	var release Release
	err = json.NewDecoder(res.Body).Decode(&release)
	if err != nil {
		return nil, fmt.Errorf("json unmarshalling: %w", err)
	}

	return &release, nil
}
