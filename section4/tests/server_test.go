package tests

import (
	"github.com/nade-harlow/go-test-exercise/section3/httpServer/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	// Create a new instance of the HTTP server
	server := httptest.NewServer(http.HandlerFunc(handler.HelloHandler))
	defer server.Close()

	t.Run("Perform a GET request to the server", func(t *testing.T) {
		// Perform a GET request to the server
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("GET request failed: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %v", err)
		}

		// Verify that the response status code is 200 OK
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
		}

		// Verify that the response body contains "Hello, World!"
		expectedBody := "Hello, World!"
		if string(body) != expectedBody {
			t.Errorf("Unexpected response body: got %s, want %s", body, expectedBody)
		}
	})

	t.Run("should fail if a POST request is sent to the server", func(t *testing.T) {
		// Perform a POST request to the server
		resp, err := http.Post(server.URL, "", nil)
		if err != nil {
			t.Fatalf("POST request failed: %v", err)
		}
		defer resp.Body.Close()

		// Verify that the response status code is 200 OK
		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
		}

		// Verify that the response body contains "Hello, World!"
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %v", err)
		}

		expectedBody := "Method Not Allowed"
		if strings.TrimSpace(string(body)) != expectedBody {
			t.Errorf("Unexpected response body: got %s, want %s", body, expectedBody)
		}
	})

}
