package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gtdavis25/aoc/solver"
)

type Solve struct {
	Day SolveDay `cmd:""`
}

type SolveDay struct {
	Year      int `arg:""`
	Day       int `arg:""`
	InputFile string
}

func (s *SolveDay) Run() error {
	params := solver.Params{}
	solver := GetSolver(s.Year, s.Day, params)
	if solver == nil {
		return fmt.Errorf("no solver for %d day %d", s.Year, s.Day)
	}

	inputFile := s.InputFile
	if inputFile == "" {
		inputFile = fmt.Sprintf("input/%d/%02d.txt", s.Year, s.Day)
	}

	f, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("opening input file: %w", err)
	}

	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}

	start := time.Now()
	result, err := solver.Solve(lines)
	if err != nil {
		return fmt.Errorf("%d day %d: %w", s.Year, s.Day, err)
	}

	duration := time.Since(start)
	fmt.Printf("part 1: %d\n", result.Part1)
	fmt.Printf("part 2: %d\n", result.Part2)
	fmt.Printf("duration: %v\n", duration)
	return nil
}
