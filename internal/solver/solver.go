package solver

import (
	"bufio"
	"io"
	"time"

	"github.com/gtdavis25/aoc/internal/option"
)

type Interface interface {
	Solve(Context) error
}

type Context interface {
	InputLines() ([]string, error)
	SetPart1(any)
	SetPart2(any)
}

type SolverContext struct {
	r     io.Reader
	start time.Time

	Part1    option.Option[any]
	Part2    option.Option[any]
	Duration time.Duration
}

func NewContext(r io.Reader) *SolverContext {
	return &SolverContext{
		r:     r,
		start: time.Now(),
	}
}

func (s *SolverContext) InputLines() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(s.r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (s *SolverContext) SetPart1(result any) {
	s.Part1 = option.New(result)
	s.Duration = time.Since(s.start)
}

func (s *SolverContext) SetPart2(result any) {
	s.Part2 = option.New(result)
	s.Duration = time.Since(s.start)
}

type Result struct {
	Part1 int
	Part2 int
}

type Factory func(Params) Interface

type Params struct{}
