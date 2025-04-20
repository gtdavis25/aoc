package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strconv"
)

//go:generate go run . generate --output-file solvers.go

type Generate struct {
	OutputFile string `required:""`
}

func (g *Generate) Run() error {
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

	f, err := os.Create(g.OutputFile)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}

	defer f.Close()
	tmpl := template.New("solvers")
	template.Must(tmpl.Parse(Template))
	tmpl.Execute(f, TemplateInput{
		DaysByYear: daysByYear,
	})

	return nil
}

//go:embed solvers.go.template
var Template string

type TemplateInput struct {
	DaysByYear map[int][]int
}
