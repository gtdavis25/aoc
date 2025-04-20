package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://adventofcode.com"

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HTTPClient
	cookie     string
}

func New(httpClient HTTPClient, cookie string) *Client {
	return &Client{
		httpClient: httpClient,
		cookie:     cookie,
	}
}

func (c *Client) GetPuzzleInput(ctx context.Context, year, day int) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s/%d/day/%d/input", baseURL, year, day)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating http request: %w", err)
	}

	req.Header.Set("Cookie", c.cookie)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected http response status: %s", resp.Status)
	}

	return resp.Body, nil
}
