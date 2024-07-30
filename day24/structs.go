package day24

type Vertex struct {
	x, y, z, vx, vy, vz float64
	ident               string
}

func (v *Vertex) line() Line {
	const T = 1e14
	return Line{v.x, v.y, v.x + v.vx*T, v.y + v.vy*T}
}

type Line struct {
	x1, y1, x2, y2 float64
}

type Point struct {
	x, y float64
}

type Pair struct {
	a Vertex
	b Vertex
}
