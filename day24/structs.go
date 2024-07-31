package day24

type Point struct {
	x, y float64
}

type Vertex struct {
	x, y, z, vx, vy, vz float64
	ident               string
}

func (vtx *Vertex) line() (Point, Point) {
	const T = 1e14
	return Point{vtx.x, vtx.y}, Point{vtx.x + vtx.vx*T, vtx.y + vtx.vy*T}
}

type Pair struct {
	vtx1 Vertex
	vtx2 Vertex
}

func (pair *Pair) lines() (Point, Point, Point, Point) {
	p1, p2 := pair.vtx1.line()
	p3, p4 := pair.vtx2.line()
	return p1, p2, p3, p4
}
