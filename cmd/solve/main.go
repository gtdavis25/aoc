package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gtdavis25/aoc/internal/registry"
	"github.com/gtdavis25/aoc/internal/solver"
)

var inputFile = flag.String("input-file", "", "the path to the file containing the puzzle input")

func main() {
	if err := solve(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func solve() error {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		return fmt.Errorf("usage: %s <year> <day>", os.Args[0])
	}

	year, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("parsing %q as year: %w", args[0], err)
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("parsing %q as day: %w", args[1], err)
	}

	params := solver.Params{}
	s := registry.GetSolver(year, day, params)
	if s == nil {
		return fmt.Errorf("no solver for %d day %d", year, day)
	}

	if *inputFile == "" {
		*inputFile = fmt.Sprintf("input/%d/%02d.txt", year, day)
	}

	f, err := os.Open(*inputFile)
	if err != nil {
		return fmt.Errorf("opening input file: %w", err)
	}

	defer f.Close()
	start := time.Now()
	if err := s.Solve(f, os.Stdout); err != nil {
		return fmt.Errorf("%d day %d: %w", year, day, err)
	}

	fmt.Printf("duration: %v\n", time.Since(start))
	return nil
}
