package client

import (
	"context"
	"net/http"
	"time"
)

type RateLimiter struct {
	roundTripper http.RoundTripper
	ch           chan struct{}
}

func NewRateLimiter(roundTripper http.RoundTripper, rps float64) (_ *RateLimiter, stop func()) {
	ctx, stop := context.WithCancel(context.Background())
	ch := make(chan struct{})
	delay := time.Duration(float64(time.Second) / rps)
	go func() {
		for {
			select {
			case <-ch:
			case <-ctx.Done():
				return
			}

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return
			}
		}
	}()

	return &RateLimiter{
		roundTripper: roundTripper,
		ch:           ch,
	}, stop
}

func (r *RateLimiter) RoundTrip(req *http.Request) (*http.Response, error) {
	select {
	case r.ch <- struct{}{}:
		return r.roundTripper.RoundTrip(req)

	case <-req.Context().Done():
		return nil, req.Context().Err()
	}
}
