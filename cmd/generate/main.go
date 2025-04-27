package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"regexp"
	"strconv"
)

const outputFile = "internal/registry/registry.go"

func main() {
	if err := generate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generate() error {
	rootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return fmt.Errorf("enumerating root directory: %w", err)
	}

	daysByYear := make(map[int][]int)
	re := regexp.MustCompile(`\d{4}`)
	for _, entry := range entries {
		if entry.IsDir() && re.MatchString(entry.Name()) {
			year, err := strconv.Atoi(entry.Name())
			if err != nil {
				return fmt.Errorf("parsing %q as year: %w", entry.Name(), err)
			}

			daysByYear[year] = []int{}
		}
	}

	re = regexp.MustCompile(`day(\d+)`)
	for year := range daysByYear {
		subdirectories, err := os.ReadDir(fmt.Sprint(year))
		if err != nil {
			return fmt.Errorf("enumerate %d: %w", year, err)
		}

		for _, subdirectory := range subdirectories {
			if subdirectory.IsDir() && re.MatchString(subdirectory.Name()) {
				day, err := strconv.Atoi(subdirectory.Name()[3:])
				if err != nil {
					return fmt.Errorf("parsing %q as day: %w", subdirectory.Name(), err)
				}

				daysByYear[year] = append(daysByYear[year], day)
			}
		}
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}

	defer f.Close()
	var b bytes.Buffer
	tmpl := template.New("solvers")
	template.Must(tmpl.Parse(registryTemplate))
	tmpl.Execute(&b, templateInput{
		DaysByYear: daysByYear,
	})

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("formatting output file: %w", err)
	}

	for len(formatted) > 0 {
		n, err := f.Write(formatted)
		if err != nil {
			return fmt.Errorf("writing output file: %w", err)
		}

		formatted = formatted[n:]
	}

	return nil
}

//go:embed registry.go.template
var registryTemplate string

type templateInput struct {
	DaysByYear map[int][]int
}
