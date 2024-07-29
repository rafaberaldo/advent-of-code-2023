package day23

type Direction struct {
	dx, dy int
}

type Point struct {
	x, y, steps int
	Direction
}
