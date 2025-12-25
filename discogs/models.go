package discogs

import "time"

// TODO: the fields here are only relevant to the Serendipity application. We should
// ensure that this client captures the whole data from the payload.

type Artist struct {
	Name        string `json:"name"`
	Anv         string `json:"anv"`
	Join        string `json:"join"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
	ID          int64  `json:"id"`
	ResourceURL string `json:"resource_url"`
	ThumbURL    string `json:"thumbnail_url"`
}

type Community struct {
	Have         int64  `json:"have"`
	Want         int64  `json:"want"`
	Rating       Rating `json:"rating"`
	Submitter    User   `json:"submitter"`
	Contributors []User `json:"contributors"`
	DataQuality  string `json:"data_quality"`
	Status       string `json:"status"`
}

type Company struct {
	Name           string `json:"name"`
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int64  `json:"id"`
	ResourceURL    string `json:"resource_url"`
	ThumbURL       string `json:"thumbnail_url"`
}

type Format struct {
	Name         string   `json:"name"`
	Quantity     string   `json:"qty"`
	Descriptions []string `json:"descriptions"`
	Text         string   `json:"text"`
}

type Identifier struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type Image struct {
	Type        string `json:"type"`
	URI         string `json:"uri"`
	ResourceURL string `json:"resource_url"`
	URI150      string `json:"uri150"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type Label struct {
	Name           string `json:"name"`
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int64  `json:"id"`
	ResourceURL    string `json:"resource_url"`
	ThumbURL       string `json:"thumbnail_url"`
}

type Track struct {
	Position     string   `json:"position"`
	Type         string   `json:"type_"`
	Title        string   `json:"title"`
	ExtraArtists []Artist `json:"extraartists"`
	Duration     string   `json:"duration"`
}

type Video struct {
	URI         string `json:"uri"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
}

type Rating struct {
	Count   int64   `json:"count"`
	Average float32 `json:"average"`
}

type Release struct {
	ID                   int64        `json:"id"`
	Status               string       `json:"status"`
	Year                 int          `json:"year"`
	ResourceURL          string       `json:"resource_url"`
	URI                  string       `json:"uri"`
	ArtistsSort          string       `json:"artists_sort"`
	Labels               []Label      `json:"labels"`
	Companies            []Company    `json:"companies"`
	Formats              []Format     `json:"formats"`
	DataQuality          string       `json:"data_quality"`
	Community            Community    `json:"community"`
	FormatQuantity       int          `json:"format_quantity"`
	DateAdded            time.Time    `json:"date_added"`
	DateChanged          time.Time    `json:"date_changed"`
	NumForSale           int          `json:"num_for_sale"`
	LowestPrice          float32      `json:"lowest_price"`
	MasterId             int64        `json:"master_id"`
	MasterURL            string       `json:"master_url"`
	Title                string       `json:"title"`
	Country              string       `json:"country"`
	ReleaseDate          string       `json:"released"`
	Notes                string       `json:"notes"`
	ReleaseDateFormatted string       `json:"released_formatted"`
	Identifiers          []Identifier `json:"identifiers"`
	Videos               []Video      `json:"videos,omitempty"`
	Genres               []string     `json:"genres"`
	Styles               []string     `json:"styles"`
	Tracklist            []Track      `json:"tracklist"`
	ExtraArtists         []Artist     `json:"extraartists"`
	Images               []Image      `json:"images,omitempty"`
	Thumb                string       `json:"thumb"`
	EstimatedWeight      int          `json:"estimated_weight"`
	BlockedFromSale      bool         `json:"blocked_from_sale"`
	Artists              []Artist     `json:"artists"`
	IsOffensive          bool         `json:"is_offensive"`

	// Still couldn't find a release with a populated "Series" field
	// so this will convert anything until we cna be sure of what this
	// this field contains to properly parse it.
	Series []map[string]any `json:"series"`
}

type User struct {
	Username    string `json:"username"`
	ResourceURL string `json:"resource_url"`
}
