package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gtdavis25/aoc/cmd/download/client"
	"github.com/gtdavis25/aoc/cmd/download/service"
)

const outputDirectory = "input"

var (
	cookie = flag.String("cookie", "", "the session cookie")
	debug  = flag.Bool("debug", false, "whether to print debug logs")
)

func main() {
	if err := download(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func download() error {
	flag.Parse()
	if cookie == nil || *cookie == "" {
		return fmt.Errorf("cookie is required")
	}

	transport := http.DefaultTransport
	if *debug {
		transport = client.NewRequestLogger(transport)
	}

	transport, stop := client.NewRateLimiter(transport, 2)
	defer stop()
	httpClient := http.Client{
		Transport: transport,
	}

	ctx := context.Background()
	client := client.New(&httpClient, *cookie)
	service := service.New(client, outputDirectory)
	switch args := flag.Args(); len(args) {
	case 0:
		return service.DownloadAll(ctx)

	case 1:
		year, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("could not parse %q as year", args[0])
		}

		return service.DownloadYear(ctx, year)

	case 2:
		year, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("could not parse %q as year", args[0])
		}

		day, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("could not parse %q as day", args[1])
		}

		return service.DownloadDay(ctx, year, day)

	default:
		fmt.Fprintf(os.Stderr, "usage: %s [year [day]]\n", os.Args[0])
		os.Exit(1)
	}

	return nil
}
