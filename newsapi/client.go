package newsapi

import (
	"crypto/tls"
	"net/http"
	"time"
)

type httpDoer interface {
	Do(r *http.Request) (*http.Response, error)
}

// NewHTTPClient returns new HTTP client with timeout
// passed as an argument and without SSL verification.
func NewHTTPClient(seconds int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(seconds) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
