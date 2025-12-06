package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fredericomozzato/discogs_client/discogs"
)

func main() {
	c := discogs.NewClient()
	r, err := c.GetRelease(context.Background(), 35770609)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(r)
}
