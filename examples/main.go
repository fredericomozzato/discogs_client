package main

import (
	"fmt"

	"github.com/fredericomozzato/discogs_client/discogs"
)

func main() {
	c := discogs.NewClient()
	r := c.GetRelease(35770609)

	fmt.Println(r)
}
