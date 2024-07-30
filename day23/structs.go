package day23

type Direction struct {
	dx, dy int
}

type Point struct {
	x, y, steps int
	dir         Direction
}

func (p *Point) state() Point {
	// basically the point without steps
	return Point{x: p.x, y: p.y}
}
