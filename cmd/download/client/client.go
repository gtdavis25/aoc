package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
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

func (c *Client) GetDaysForYear(ctx context.Context, year int) ([]int, error) {
	url := fmt.Sprintf("%s/%d", baseURL, year)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending http request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http response status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading http response: %w", err)
	}

	var days []int
	re := regexp.MustCompile(`Day (\d+)`)
	for _, match := range re.FindAllSubmatch(data, -1) {
		day, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return nil, fmt.Errorf("parsing day %q: %w", match[1], err)
		}

		days = append(days, day)
	}

	return days, nil
}

func (c *Client) GetYears(ctx context.Context) ([]int, error) {
	url := fmt.Sprintf("%s/events", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending http request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http response status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading http response body: %w", err)
	}

	var years []int
	re := regexp.MustCompile(`\[(\d{4})\]`)
	for _, match := range re.FindAllSubmatch(data, -1) {
		year, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return nil, fmt.Errorf("parsing year %q: %w", match[1], err)
		}

		years = append(years, year)
	}

	return years, nil
}
