package discogs

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetRelease(t *testing.T) {
	t.Run("builds request to correct endpoint", func(t *testing.T) {
		var requestMethod string
		var requestPath string
		var requestUserAgent string

		handler := func(w http.ResponseWriter, r *http.Request) {
			requestMethod = r.Method
			requestPath = r.URL.Path
			requestUserAgent = r.Header.Get("User-Agent")
			w.Write([]byte(`{"foo":"bar"}`))
		}

		testServer := httptest.NewServer(http.HandlerFunc(handler))
		defer testServer.Close()

		url := testServer.URL
		timeout := 30
		userAgent := "test app"
		client := NewClient(url, timeout, userAgent)

		ctx := context.Background()
		release, err := client.GetRelease(ctx, 1234)

		if err != nil {
			t.Fatalf("error should be nil, instead got %v", err)
		}

		if release == nil {
			t.Error("Release should not be nil")
		}

		expectedMethod := "GET"
		expectedPath := "/releases/1234"

		if requestMethod != expectedMethod {
			t.Errorf("got method %q want %q", requestMethod, expectedMethod)
		}
		if requestPath != expectedPath {
			t.Errorf("got path %q want %q\n", requestPath, expectedPath)
		}
		if requestUserAgent != userAgent {
			t.Errorf("got User-Agent %q want %q", requestUserAgent, userAgent)
		}
	})

	t.Run("correctly parses JSON data to Release type", func(t *testing.T) {
		jsonData, err := os.ReadFile("../../testdata/responses/release_minimal.json")
		if err != nil {
			t.Fatal("could not open JSON file")
		}

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		}

		testServer := httptest.NewServer(http.HandlerFunc(handler))
		defer testServer.Close()

		url := testServer.URL
		timeout := 20
		userAgent := "test"
		client := NewClient(url, timeout, userAgent)
		ctx := context.Background()

		release, err := client.GetRelease(ctx, 1234)
		if err != nil {
			t.Fatalf("Client error: %s", err)
		}
		if release == nil {
			t.Fatal("want Release got nil")
		}
		if release.ID != int64(12345) {
			t.Errorf("got ID %d, want %d", release.ID, int64(12345))
		}
		if release.Status != "Accepted" {
			t.Errorf("got Status %q, want %q", release.Status, "Accepted")
		}
		if release.Year != 2025 {
			t.Errorf("got Year %d, want %d", release.Year, 2025)
		}
		if release.ResourceURL != "https://api.discogs.com/releases/12345" {
			t.Errorf("got ResourceURL %q, want %q", release.ResourceURL, "https://api.discogs.com/releases/12345")
		}
		if release.URI != "https://www.discogs.com/release/12345-Test-Artist-Test-Album" {
			t.Errorf("got URI %q, want %q", release.URI, "https://www.discogs.com/release/12345-Test-Artist-Test-Album")
		}

		if len(release.Artists) != 1 {
			t.Errorf("got Arists length %d, want %d", len(release.Artists), 1)
		}
		artist := release.Artists[0]
		if artist.Name != "Test Artist" {
			t.Errorf("got Artist Name %q, want %q", artist.Name, "Test Artist")
		}
		if artist.Anv != "" {
			t.Errorf("got Artist Env %q, want empty string", artist.Anv)
		}
		if artist.ID != 100 {
			t.Errorf("got Artist ID %d, want %d", artist.ID, 100)
		}
		if artist.ResourceURL != "https://api.discogs.com/artists/100" {
			t.Errorf("got Artist ResourceURL %q, want %q", artist.ResourceURL, "https://api.discogs.com/artists/100")
		}

		if release.ArtistsSort != "Test Artist" {
			t.Errorf("got ArtistsSort %q, want %q", release.ArtistsSort, "Test Artist")
		}

		if len(release.Labels) != 1 {
			t.Errorf("got Labels length %d, want %d", len(release.Labels), 1)
		}
		label := release.Labels[0]
		if label.Name != "Test Label" {
			t.Errorf("got Label Name %q, want %q", label.Name, "Test Label")
		}
		if label.Catno != "TEST001" {
			t.Errorf("got Label Catno %q, want %q", label.Catno, "TEST001")
		}
		if label.EntityType != "1" {
			t.Errorf("got Label Entity Type %q, want %q", label.EntityType, "1")
		}
		if label.EntityTypeName != "Label" {
			t.Errorf("got Label Entity Type Name %q, want %q", label.EntityTypeName, "Label")
		}
		if label.ID != 200 {
			t.Errorf("got Label ID %d, want %d", label.ID, 200)
		}
		if label.ResourceURL != "https://api.discogs.com/labels/200" {
			t.Errorf("got Label Resource URL %q, want %q", label.ResourceURL, "https://api.discogs.com/labels/200")
		}

		if len(release.Companies) != 1 {
			t.Errorf("got Companies length %d, want %d", len(release.Companies), 1)
		}
		company := release.Companies[0]
		if company.Name != "Test Company" {
			t.Errorf("got Company Name %q, want %q", company.Name, "Test Company")
		}
		if company.EntityType != "13" {
			t.Errorf("got Company Entity Type %q, want %q", company.EntityType, "13")
		}
		if company.EntityTypeName != "Phonographic Copyright (p)" {
			t.Errorf("got Company Entity Type Name %q, want %q", company.EntityTypeName, "Phonographic Copyright (p)")
		}
		if company.ID != 300 {
			t.Errorf("got Company ID %d, want %d", company.ID, 300)
		}
		if company.ResourceURL != "https://api.discogs.com/labels/300" {
			t.Errorf("got Company Resource URL %q, want %q", company.ResourceURL, "https://api.discogs.com/labels/300")
		}

		if len(release.Formats) != 1 {
			t.Errorf("got Formats length %d, want %d", len(release.Formats), 1)
		}
		format := release.Formats[0]
		if format.Name != "Vinyl" {
			t.Errorf("got Format Name %q, want %q", format.Name, "Vinyl")
		}
		if format.QTY != "1" {
			t.Errorf("got Format QTY %q, want %q", format.QTY, "1")
		}
		if len(format.Descriptions) != 2 {
			t.Errorf("got Format Descriptions length %d, want %d", len(format.Descriptions), 2)
		}
		if format.Descriptions[0] != "LP" {
			t.Errorf("got first Format Description %q, want %q", format.Descriptions[0], "LP")
		}
		if format.Descriptions[1] != "Album" {
			t.Errorf("got first Format Description %q, want %q", format.Descriptions[1], "Album")
		}
		if format.Text != "Black" {
			t.Errorf("got Format Text %q, want %q", format.Text, "Black")
		}

		if release.DataQuality != "Correct" {
			t.Errorf("got Data Quality %q, want %q", release.DataQuality, "Correct")
		}

		community := release.Community
		if community.Have != 150 {
			t.Errorf("got Community Have %d, want %d", community.Have, 150)
		}
		if community.Want != 300 {
			t.Errorf("got Community Want %d, want %d", community.Want, 300)
		}

		rating := community.Rating
		if rating.Count != 50 {
			t.Errorf("got Community Rating Count %d, want %d", rating.Count, 50)
		}
		if rating.Average != 4.5 {
			t.Errorf("got Community Rating Average %f, want %f", rating.Average, 4.5)
		}

		submitter := community.Submitter
		if submitter.Username != "test_user" {
			t.Errorf("got Community Submitter Username %q, want %q", submitter.Username, "test_user")
		}
		if submitter.ResourceURL != "https://api.discogs.com/users/test_user" {
			t.Errorf("got Community Submitter ResourceUrl %q, want %q", submitter.ResourceURL, "https://api.discogs.com/users/test_user")
		}

		contributors := community.Contributors
		if len(contributors) != 1 {
			t.Errorf("got Community Contributors length %d, want %d", len(contributors), 1)
		}
		contributor := contributors[0]
		if contributor.Username != "contributor1" {
			t.Errorf("got Community Contributor Username %q, want %q", contributor.Username, "contributer1")
		}
		if contributor.ResourceURL != "https://api.discogs.com/users/contributor1" {
			t.Errorf("got Community Contributor Resource URL %q, want %q", contributor.ResourceURL, "https://api.discogs.com/users/contributor1")
		}
		if community.DataQuality != "Correct" {
			t.Errorf("got Community Data Quality %q, want %q", community.DataQuality, "Correct")
		}
		if community.Status != "Accepted" {
			t.Errorf("got Community Status %q, want %q", community.Status, "Accepted")
		}

		if release.FormatQuantity != 1 {
			t.Errorf("got Format Quantity %d, want %d", release.FormatQuantity, 1)
		}

		parsedDateAdded, err := time.Parse(time.RFC3339, "2020-05-15T10:30:00-07:00")
		if err != nil {
			t.Fatal("failed to parse release.DateAdded")
		}
		if !release.DateAdded.Equal(parsedDateAdded) {
			t.Errorf("got Date Added %v, want %v", release.DateAdded, "2020-05-15T10:30:00-07:00")
		}

		parsedDateChanged, err := time.Parse(time.RFC3339, "2020-06-20T14:45:00-07:00")
		if err != nil {
			t.Fatal("failed to parse release.DateChanged")
		}
		if !release.DateChanged.Equal(parsedDateChanged) {
			t.Errorf("got Date Changed %v, want %v", release.DateChanged, parsedDateChanged)
		}
		if release.NumForSale != 10 {
			t.Errorf("got NumForSale %d, want %d", release.NumForSale, 10)
		}
		if release.LowestPrice != 25.99 {
			t.Errorf("got Lowest Price %f, want %f", release.LowestPrice, 25.99)
		}
		if release.MasterID != 54321 {
			t.Errorf("got Master ID %d, want %d", release.MasterID, 54321)
		}
		if release.MasterURL != "https://api.discogs.com/masters/54321" {
			t.Errorf("got Master URL %q, want %q", release.MasterURL, "https://api.discogs.com/masters/54321")
		}
		if release.Title != "Test Album" {
			t.Errorf("got Title %q, want %q", release.Title, "Test Album")
		}
		if release.Country != "US" {
			t.Errorf("got Country %q, want %q", release.Country, "US")
		}
		if release.ReleaseDate != "2017-04-07" {
			t.Errorf("got Release Date %q, want %q", release.ReleaseDate, "2017-04-07")
		}
		if release.Notes != "Test release notes." {
			t.Errorf("got Notes %q, want %q", release.Notes, "Test release notes.")
		}
		if release.ReleaseDateFormatted != "07 April 2017" {
			t.Errorf("got Release Date Formatted %q, want %q", release.ReleaseDateFormatted, "07 April 2017")
		}

		identifiers := release.Identifiers
		if len(identifiers) != 1 {
			t.Errorf("got Identifiers length %d, want %d", len(identifiers), 1)
		}
		identifier := identifiers[0]
		if identifier.Type != "Barcode" {
			t.Errorf("got Identifier Type %q, want %q", identifier.Type, "Barcode")
		}
		if identifier.Value != "123456789012" {
			t.Errorf("got Identifier Value %q, want %q", identifier.Value, "123456789012")
		}
		if identifier.Description != "" {
			t.Errorf("got Identifier Description %q, want empty string", identifier.Description)
		}

		videos := release.Videos
		if len(videos) != 1 {
			t.Errorf("got Videos length %d, want %d", len(videos), 1)
		}
		video := videos[0]
		if video.URI != "https://www.youtube.com/watch?v=test123" {
			t.Errorf("got Video URI %q, want %q", video.URI, "https://www.youtube.com/watch?v=test123")
		}
		if video.Title != "Test Video" {
			t.Errorf("got Video Title %q, want %q", video.Title, "Test Video")
		}
		if video.Description != "Official Music Video" {
			t.Errorf("got Video Description %q, want %q", video.Description, "Official Music Video")
		}
		if video.Duration != 240 {
			t.Errorf("got Video Duration %d, want %d", video.Duration, 240)
		}
		if video.Embed != true {
			t.Errorf("got Video Embed %v, want %v", video.Embed, true)
		}

		genres := release.Genres
		if len(genres) != 2 {
			t.Errorf("got Genres length %d, want %d", len(genres), 2)
		}
		if genres[0] != "Rock" {
			t.Errorf("got Genre %q, want %q", genres[0], "Rock")
		}
		if genres[1] != "Electronic" {
			t.Errorf("got Genre %q, want %q", genres[1], "Electronic")
		}

		styles := release.Styles
		if len(styles) != 2 {
			t.Errorf("got Styles length %d, %d", len(styles), 2)
		}
		if styles[0] != "Indie Rock" {
			t.Errorf("got Style %q, want %q", styles[0], "Indie Rock")
		}
		if styles[1] != "Synth-pop" {
			t.Errorf("got Style %q, want %q", styles[1], "Synth-pop")
		}

		tracklist := release.Tracklist
		if len(tracklist) != 2 {
			t.Errorf("got Tracklist length %d, want %d", len(tracklist), 2)
		}
		track1 := tracklist[0]
		if track1.Position != "A1" {
			t.Errorf("got Track 1 Position %q, want %q", track1.Position, "A1")
		}
		if track1.Type != "track" {
			t.Errorf("got Track Type %q, want %q", track1.Type, "track")
		}
		if track1.Title != "Test Track 1" {
			t.Errorf("got Track 1 Title %q, want %q", track1.Title, "Test Track 1")
		}
		if track1.Duration != "3:45" {
			t.Errorf("got Track 1 Duration %q, want %q", track1.Duration, "3:45")
		}
		if len(track1.ExtraArtists) != 1 {
			t.Errorf("got Track 1 ExtraArtists length %d, want %d", len(track1.ExtraArtists), 1)
		}
		trackArtist := track1.ExtraArtists[0]
		if trackArtist.Name != "Test Producer" {
			t.Errorf("got Track 1 ExtraArtist Name %q, want %q", trackArtist.Name, "Test Producer")
		}
		if trackArtist.Role != "Producer" {
			t.Errorf("got Track 1 ExtraArtist Role %q, want %q", trackArtist.Role, "Producer")
		}
		if trackArtist.ID != 400 {
			t.Errorf("got Track 1 ExtraArtist ID %d, want %d", trackArtist.ID, 400)
		}

		track2 := tracklist[1]
		if track2.Position != "A2" {
			t.Errorf("got Track 2 Position %q, want %q", track2.Position, "A2")
		}
		if track2.Type != "track" {
			t.Errorf("got Track 2 Type %q, want %q", track2.Type, "track")
		}
		if track2.Title != "Test Track 2" {
			t.Errorf("got Track 2 Title %q, want %q", track2.Title, "Test Track 2")
		}
		if track2.Duration != "4:20" {
			t.Errorf("got Track 2 Duration %q, want %q", track2.Duration, "4:20")
		}
		if len(track2.ExtraArtists) != 0 {
			t.Errorf("got Track 2 ExtraArtists length %d, want %d", len(track2.ExtraArtists), 0)
		}

		extraArtists := release.ExtraArtists
		if len(extraArtists) != 1 {
			t.Errorf("got ExtraArtists length %d, want %d", len(extraArtists), 1)
		}
		extraArtist := extraArtists[0]
		if extraArtist.Name != "Test Mixer" {
			t.Errorf("got ExtraArtist Name %q, want %q", extraArtist.Name, "Test Mixer")
		}
		if extraArtist.Role != "Mixed By" {
			t.Errorf("got ExtraArtist Role %q, want %q", extraArtist.Role, "Mixed By")
		}
		if extraArtist.ID != 500 {
			t.Errorf("got ExtraArtist ID %d, want %d", extraArtist.ID, 500)
		}

		images := release.Images
		if len(images) != 1 {
			t.Errorf("got Images length %d, want %d", len(images), 1)
		}
		image := images[0]
		if image.Type != "primary" {
			t.Errorf("got Image Type %q, want %q", image.Type, "primary")
		}
		if image.URI != "https://img.discogs.com/test-image.jpg" {
			t.Errorf("got Image URI %q, want %q", image.URI, "https://img.discogs.com/test-image.jpg")
		}
		if image.ResourceURL != "https://img.discogs.com/test-image.jpg" {
			t.Errorf("got Image ResourceURL %q, want %q", image.ResourceURL, "https://img.discogs.com/test-image.jpg")
		}
		if image.URI150 != "https://img.discogs.com/test-image-150.jpg" {
			t.Errorf("got Image URI150 %q, want %q", image.URI150, "https://img.discogs.com/test-image-150.jpg")
		}
		if image.Width != 600 {
			t.Errorf("got Image Width %d, want %d", image.Width, 600)
		}
		if image.Height != 600 {
			t.Errorf("got Image Height %d, want %d", image.Height, 600)
		}

		if release.Thumb != "https://img.discogs.com/test-thumb.jpg" {
			t.Errorf("got Thumb %q, want %q", release.Thumb, "https://img.discogs.com/test-thumb.jpg")
		}
		if release.EstimatedWeight != 230 {
			t.Errorf("got EstimatedWeight %d, want %d", release.EstimatedWeight, 230)
		}
		if release.BlockedFromSale != false {
			t.Errorf("got BlockedFromSale %v, want %v", release.BlockedFromSale, false)
		}
		if release.IsOffensive != false {
			t.Errorf("got IsOffensive %v, want %v", release.IsOffensive, false)
		}
		if len(release.Series) != 0 {
			t.Errorf("got Series length %d, want %d", len(release.Series), 0)
		}
	})
}
