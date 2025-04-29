package solver

import (
	"io"
)

type Interface interface {
	Solve(io.Reader, io.Writer) error
}

type Factory func(Params) Interface

type Params struct{}
