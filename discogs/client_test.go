package discogs

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
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
}
