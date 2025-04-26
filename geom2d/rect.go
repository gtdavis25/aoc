package geom2d

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (r Rect) Contains(p Point) bool {
	return r.X <= p.X && p.X < r.X+r.Width && r.Y <= p.Y && p.Y < r.Y+r.Height
}
