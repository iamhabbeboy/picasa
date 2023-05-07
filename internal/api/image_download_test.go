package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func TestDownloadFile(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve a file with some content
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "some content")
	}))
	defer server.Close()

	// Replace the URL with the mock server URL
	url := server.URL

	// Download the file
	resp, err := http.Get(url)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Verify that the file contents are correct
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "some content\n", string(body))
}
