package main

import (
	"context"
	"net/http"

	"github.com/gtdavis25/aoc/client"
	"github.com/gtdavis25/aoc/download"
)

type Download struct {
	Day DownloadDay `cmd:""`
}

type DownloadDay struct {
	Year       int    `arg:""`
	Day        int    `arg:""`
	Cookie     string `required:""`
	OutputFile string `required:""`
}

func (d *DownloadDay) Run() error {
	httpClient := http.Client{}
	aocClient := client.New(&httpClient, d.Cookie)
	downloadService := download.NewService(aocClient)
	return downloadService.DownloadPuzzleInput(context.Background(), d.Year, d.Day, d.OutputFile)
}
