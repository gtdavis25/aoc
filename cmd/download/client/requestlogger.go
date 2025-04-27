package client

import (
	"log"
	"net/http"
	"time"
)

type RequestLogger struct {
	roundTripper http.RoundTripper
}

func NewRequestLogger(roundTripper http.RoundTripper) *RequestLogger {
	return &RequestLogger{
		roundTripper: roundTripper,
	}
}

func (l *RequestLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := l.roundTripper.RoundTrip(req)
	if resp != nil {
		log.Printf("%s %s %s %v", req.Method, req.URL, resp.Status, time.Since(start))
	}

	return resp, err
}
