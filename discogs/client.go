package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const UserAgent = "DiscogsClient github.com/fredericomozzato/discogs_client"

type Client struct {
	httpClient *http.Client
	url        string
}

func NewClient() Client {
	return Client{
		&http.Client{Timeout: 30 * time.Second},
		"https://api.discogs.com",
	}
}

// TODO: return err instead of panic
func (c Client) GetRelease(id int) Release {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/releases/%d", c.url, id),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", UserAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("unexpected HTTP status: %d\n", res.StatusCode)
	}

	var release Release
	err = json.NewDecoder(res.Body).Decode(&release)
	if err != nil {
		log.Fatal(err)
	}

	return release
}
