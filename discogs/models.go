package discogs

// TODO: the fields here are only relevant to the Serendipity application. We should
// ensure that this client captures the whole data from the payload.

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"resource_url"`
}

type Image struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type Track struct {
	Duration string `json:"duration"`
	Position string `json:"position"`
	Title    string `json:"title"`
	Type     string `json:"type_"`
}

type Video struct {
	URI string `json:"uri"`
}

type Release struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Artists   []Artist `json:"artists"`
	Year      int      `json:"year"`
	Genres    []string `json:"genres"`
	Country   string   `json:"country"`
	Images    []Image  `json:"images,omitempty"`
	Tracklist []Track  `json:"tracklist"`
	URI       string   `json:"uri"`
	Videos    []Video  `json:"videos,omitempty"`
}
