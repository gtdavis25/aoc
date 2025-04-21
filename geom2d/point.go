package geom2d

type Point struct {
	X int
	Y int
}

func (p Point) Add(d Point) Point {
	return Point{
		X: p.X + d.X,
		Y: p.Y + d.Y,
	}
}
