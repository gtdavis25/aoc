package regression_test

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/gtdavis25/aoc/internal/registry"
	"github.com/gtdavis25/aoc/internal/solver"
)

var update = flag.Bool("update", false, "controls whether to update the test data file")

func TestSolvers(t *testing.T) {
	resultCh := make(chan result)
	results := make(map[int]map[int][]any)
	var wg sync.WaitGroup
	for year, solvers := range registry.Solvers {
		results[year] = make(map[int][]any)
		for day, factory := range solvers {
			results[year][day] = make([]any, 2)
			wg.Add(1)
			go func(year, day int, factory solver.Factory) {
				defer wg.Done()
				path := fmt.Sprintf("../../input/%d/%02d.txt", year, day)
				f, err := os.Open(path)
				if err != nil {
					t.Errorf("open input file: %v", err)
					return
				}

				defer f.Close()
				s := factory(solver.Params{})
				if err := s.Solve(testContext{
					r:       f,
					year:    year,
					day:     day,
					results: resultCh,
				}); err != nil {
					t.Errorf("solving %d day %d: %v", year, day, err)
				}
			}(year, day, factory)
		}
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for r := range resultCh {
		results[r.year][r.day][r.part-1] = r.value
	}

	if *update {
		f, err := os.Create("testdata/results.json")
		if err != nil {
			t.Fatalf("opening test data file: %v", err)
		}

		defer f.Close()
		if err := json.NewEncoder(f).Encode(results); err != nil {
			t.Fatalf("writing test data file: %v", err)
		}
	} else {
		f, err := os.Open("testdata/results.json")
		if err != nil {
			t.Fatalf("opening test data file: %v", err)
		}

		defer f.Close()
		var expected map[int]map[int][]int
		if err := json.NewDecoder(f).Decode(&expected); err != nil {
			t.Fatalf("reading test data file: %v", err)
		}

		for year, days := range results {
			if _, ok := expected[year]; !ok {
				continue
			}

			for day, parts := range days {
				if _, ok := expected[year][day]; !ok {
					continue
				}

				for i, got := range parts {
					t.Run(fmt.Sprintf("%d day %d part %d", year, day, i+1), func(t *testing.T) {
						if want := expected[year][day][i]; got != want {
							t.Errorf("got %v, want %v", got, want)
						}
					})
				}
			}
		}
	}
}

type testContext struct {
	r       io.Reader
	year    int
	day     int
	results chan<- result
}

func (c testContext) InputLines() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(c.r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}

	return lines, nil
}

func (c testContext) SetPart1(value any) {
	c.results <- result{
		year:  c.year,
		day:   c.day,
		part:  1,
		value: value,
	}
}

func (c testContext) SetPart2(value any) {
	c.results <- result{
		year:  c.year,
		day:   c.day,
		part:  2,
		value: value,
	}
}

type result struct {
	year  int
	day   int
	part  int
	value any
}
