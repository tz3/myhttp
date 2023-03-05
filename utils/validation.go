package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const (
	scheme1 = "http://"
	scheme2 = "https://"
)

// FixURLScheme adds the "http" scheme to a URL if it doesn't have one
func FixURLScheme(url string) string {
	if !strings.HasPrefix(url, scheme1) && !strings.HasPrefix(url, scheme2) {
		return scheme1 + url
	}
	return url
}

// ValidateURL checks whether a URL is valid
func ValidateURL(rawURL string) error {
	const urlRegex = `^(?i)(http|https):\/\/([^\s\/$.?#]+\.[^\s\/$.?#]+)+([:]\d+)?([\/?].*)?$`
	match, _ := regexp.MatchString(urlRegex, rawURL)
	if !match {
		return fmt.Errorf("invalid URL: %s", rawURL)
	}

	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}

	return nil
}
