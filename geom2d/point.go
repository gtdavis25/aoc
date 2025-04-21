package geom2d

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
