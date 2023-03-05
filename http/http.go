package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Result represents the result of an HTTP request
type Result struct {
	URL     string
	MD5Hash string
	Error   error
}

// Fetch performs an HTTP GET request to the specified URL
func Fetch(url string) ([]byte, error) {
	client := &http.Client{Timeout: 100 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
