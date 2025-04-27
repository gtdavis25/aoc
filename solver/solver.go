package solver

type Interface interface {
	Solve([]string) (Result, error)
}

type Result struct {
	Part1 int
	Part2 int
}

type Factory func(Params) Interface

type Params struct{}
