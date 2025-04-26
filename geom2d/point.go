package geom2d

import "iter"

type Point struct {
	X int
	Y int
}

func (p Point) Complement() Point {
	return Point{
		X: -p.X,
		Y: -p.Y,
	}
}

func (p Point) Add(d Point) Point {
	return Point{
		X: p.X + d.X,
		Y: p.Y + d.Y,
	}
}

func (p Point) Subtract(d Point) Point {
	return p.Add(d.Complement())
}

func Up() Point {
	return Point{X: 0, Y: -1}
}

func Down() Point {
	return Point{X: 0, Y: 1}
}

func Left() Point {
	return Point{X: -1, Y: 0}
}

func Right() Point {
	return Point{X: 1, Y: 0}
}

func (p Point) Adjacent() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for _, n := range []Point{
			p.Add(Up()),
			p.Add(Right()),
			p.Add(Down()),
			p.Add(Left()),
		} {
			if !yield(n) {
				return
			}
		}
	}
}
