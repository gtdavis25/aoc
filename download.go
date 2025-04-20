package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gtdavis25/aoc/client"
	"github.com/gtdavis25/aoc/download"
)

type Download struct {
	All  DownloadAll  `cmd:""`
	Year DownloadYear `cmd:""`
	Day  DownloadDay  `cmd:""`
}

type DownloadAll struct {
	Cookie          string `required:""`
	OutputDirectory string
}

func (d *DownloadAll) Run() error {
	outputDirectory := d.OutputDirectory
	if outputDirectory == "" {
		outputDirectory = "input"
	}

	rateLimiter, stop := client.NewRateLimiter(http.DefaultTransport, 2)
	defer stop()
	httpClient := http.Client{
		Transport: rateLimiter,
	}

	aocClient := client.New(&httpClient, d.Cookie)
	downloadService := download.NewService(aocClient)
	return downloadService.DownloadAll(context.Background(), outputDirectory)
}

type DownloadYear struct {
	Year            int    `arg:""`
	Cookie          string `required:""`
	OutputDirectory string
}

func (d *DownloadYear) Run() error {
	outputDirectory := d.OutputDirectory
	if outputDirectory == "" {
		outputDirectory = fmt.Sprintf("input/%d", d.Year)
	}

	rateLimiter, stop := client.NewRateLimiter(http.DefaultTransport, 2)
	defer stop()
	httpClient := http.Client{
		Transport: rateLimiter,
	}

	aocClient := client.New(&httpClient, d.Cookie)
	downloadService := download.NewService(aocClient)
	return downloadService.DownloadYear(context.Background(), d.Year, outputDirectory)
}

type DownloadDay struct {
	Year       int    `arg:""`
	Day        int    `arg:""`
	Cookie     string `required:""`
	OutputFile string
}

func (d *DownloadDay) Run() error {
	outputFile := d.OutputFile
	if outputFile == "" {
		outputFile = fmt.Sprintf("input/%d/%02d.txt", d.Year, d.Day)
	}

	rateLimiter, stop := client.NewRateLimiter(http.DefaultTransport, 2)
	defer stop()
	httpClient := http.Client{
		Transport: rateLimiter,
	}

	aocClient := client.New(&httpClient, d.Cookie)
	downloadService := download.NewService(aocClient)
	return downloadService.DownloadDay(context.Background(), d.Year, d.Day, outputFile)
}
