package discogsclient

import "github.com/fredericomozzato/discogs_client/internal/discogs"

type Client = discogs.Client

func NewClient(url string, timeout int, userAgent string) *Client {
	return discogs.NewClient(url, timeout, userAgent)
}

func NewDefaultClient(userAgent string) *Client {
	return discogs.NewDefaultClient(userAgent)
}
