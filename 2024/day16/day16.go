package day16

import (
	"container/heap"
	"fmt"
	"iter"
	"slices"

	"github.com/gtdavis25/aoc/internal/geom2d"
	"github.com/gtdavis25/aoc/internal/solver"
)

type Solver struct{}

func NewSolver(_ solver.Params) *Solver {
	return &Solver{}
}

func (s *Solver) Solve(context solver.Context) error {
	lines, err := context.InputLines()
	if err != nil {
		return err
	}

	paths, err := getMinimumScore(lines)
	if err != nil {
		return err
	}

	var part1 int
	positions := make(map[geom2d.Point]struct{})
	for p := range paths {
		if part1 != 0 && p.score > part1 {
			break
		}

		part1 = p.score
		for _, pos := range p.points {
			positions[pos] = struct{}{}
		}
	}

	context.SetPart1(part1)
	context.SetPart2(len(positions))
	return nil
}

type path struct {
	points []geom2d.Point
	score  int
}

func getMinimumScore(lines []string) (iter.Seq[path], error) {
	start, err := getStartNode(lines)
	if err != nil {
		return nil, err
	}

	return func(yield func(path) bool) {
		seen := make(map[state]int)
		q := queue{start}
		for len(q) > 0 {
			n := heap.Pop(&q).(node)
			if lines[n.pos.Y][n.pos.X] == 'E' && !yield(getPath(n)) {
				return
			}

			if min, ok := seen[n.state]; ok && n.score > min {
				continue
			}

			seen[n.state] = n.score
			if nextPos := n.pos.Add(n.vel); lines[nextPos.Y][nextPos.X] != '#' {
				heap.Push(&q, node{
					state: state{
						pos: nextPos,
						vel: n.vel,
					},
					score: n.score + 1,
					prev:  &n,
				})
			}

			for _, next := range []node{
				{
					state: state{
						pos: n.pos,
						vel: geom2d.Point{
							X: -n.vel.Y,
							Y: n.vel.X,
						},
					},
					score: n.score + 1000,
					prev:  &n,
				},
				{
					state: state{
						pos: n.pos,
						vel: geom2d.Point{
							X: n.vel.Y,
							Y: -n.vel.X,
						},
					},
					score: n.score + 1000,
					prev:  &n,
				},
			} {
				if min, ok := seen[next.state]; ok && next.score > min {
					continue
				}

				heap.Push(&q, next)
			}
		}
	}, nil
}

func getPath(endNode node) path {
	var points []geom2d.Point
	for n := &endNode; n != nil; n = n.prev {
		if !slices.Contains(points, n.pos) {
			points = append(points, n.pos)
		}
	}

	return path{
		points: points,
		score:  endNode.score,
	}
}

type node struct {
	state
	score int
	prev  *node
}

func getStartNode(lines []string) (node, error) {
	for y, line := range lines {
		for x := range len(line) {
			if line[x] == 'S' {
				return node{
					state: state{
						pos: geom2d.Point{
							X: x,
							Y: y,
						},
						vel: geom2d.Right(),
					},
					score: 0,
				}, nil
			}
		}
	}

	return node{}, fmt.Errorf("no start position")
}

type state struct {
	pos geom2d.Point
	vel geom2d.Point
}

type queue []node

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].score < q[j].score
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(val any) {
	*q = append(*q, val.(node))
}

func (q *queue) Pop() any {
	n := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return n
}
