package regression_test

import (
	"crypto/md5"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/gtdavis25/aoc/internal/registry"
	"github.com/gtdavis25/aoc/internal/solver"
)

var update = flag.Bool("update", false, "Controls whether to update test data files")

func TestSolver(t *testing.T) {
	if *update {
		if err := os.RemoveAll("testdata/TestSolver"); err != nil && !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("cleaning result directory: %v", err)
		}
	}

	if err := os.MkdirAll("testdata/TestSolver", 0777); err != nil && !errors.Is(err, os.ErrExist) {
		t.Fatalf("creating result directory: %v", err)
	}

	for year, factories := range registry.Solvers {
		for day, factory := range factories {
			year := year
			day := day
			solver := factory(solver.Params{})
			t.Run(fmt.Sprintf("%d day %02d", year, day), func(t *testing.T) {
				t.Parallel()
				in, err := os.Open(fmt.Sprintf("../../input/%d/%02d.txt", year, day))
				if err != nil {
					t.Fatalf("opening input file: %v", err)
				}

				defer in.Close()
				md5 := md5.New()
				if err := solver.Solve(in, md5); err != nil {
					t.Fatalf("running solver: %v", err)
				}

				result := md5.Sum(make([]byte, 0, 16))
				if *update {
					out, err := os.Create(fmt.Sprintf("testdata/%s.md5", t.Name()))
					if err != nil {
						t.Fatalf("creating result file: %v", err)
					}

					defer out.Close()
					if _, err := fmt.Fprintf(out, "%x", result); err != nil {
						t.Fatalf("writing result file: %v", err)
					}
				} else {
					out, err := os.Open(fmt.Sprintf("testdata/%s.md5", t.Name()))
					if err != nil {
						t.Fatalf("opening result file: %v", err)
					}

					defer out.Close()
					var want []byte
					if _, err := fmt.Fscanf(out, "%x", &want); err != nil {
						t.Fatalf("reading result file: %v", err)
					}

					if !slices.Equal(want, result) {
						t.Fatalf("got %x, want %x", result, want)
					}
				}
			})
		}
	}
}
