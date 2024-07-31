package day24

import (
	"aoc2023/assert"
	"fmt"
	"math"
	"os"
	"time"
)

func Part2() int {
	started := time.Now()
	input, err := os.ReadFile("day24/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	var vertices = parseInput(input)
	var pairs = createPairs(vertices)
	var result = simulate2(pairs)

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

// https://www.reddit.com/r/adventofcode/comments/18q40he/
func simulate2(pairs []Pair) int {
	var equationsXY [][]float64
	var equationsXZ [][]float64

	for _, pair := range pairs[:4] {
		vtx1 := pair.vtx1
		vtx2 := pair.vtx2

		var XY = func() []float64 {
			x := vtx2.vy - vtx1.vy
			y := vtx1.vx - vtx2.vx
			vx := vtx1.y - vtx2.y
			vy := vtx2.x - vtx1.x
			rhs := vtx2.x*vtx2.vy - vtx2.y*vtx2.vx - vtx1.x*vtx1.vy + vtx1.y*vtx1.vx
			return []float64{x, y, vx, vy, rhs}
		}

		var XZ = func() []float64 {
			x := vtx2.vz - vtx1.vz
			z := vtx1.vx - vtx2.vx
			vx := vtx1.z - vtx2.z
			vz := vtx2.x - vtx1.x
			rhs := vtx2.x*vtx2.vz - vtx2.z*vtx2.vx - vtx1.x*vtx1.vz + vtx1.z*vtx1.vx
			return []float64{x, z, vx, vz, rhs}
		}

		equationsXY = append(equationsXY, XY())
		equationsXZ = append(equationsXZ, XZ())

	}

	rXY := solve(equationsXY)
	rXZ := solve(equationsXZ)

	x := math.Round(rXY[0])
	y := math.Round(rXY[1])
	z := math.Round(rXZ[1])
	// vx := math.Round(rXZ[2])
	// vy := math.Round(rXY[3])
	// vz := math.Round(rXZ[3])
	// fmt.Println(x, y, z, vx, vy, vz)

	return int(x + y + z)
}

// https://en.wikipedia.org/wiki/Gaussian_elimination
func solve(eqs [][]float64) []float64 {
	var n = len(eqs)

	// Forward Elimination
	for i := range n {
		// Pivoting
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(eqs[k][i]) > math.Abs(eqs[maxRow][i]) {
				maxRow = k
			}
		}
		if maxRow != i {
			eqs[i], eqs[maxRow] = eqs[maxRow], eqs[i]
		}

		// Make the diagonal element of row i to be 1 (pivoting step)
		pivot := eqs[i][i]
		assert.Assert(pivot != 0, "matrix is singular")
		for j := i; j <= n; j++ {
			eqs[i][j] /= pivot
		}

		// Make the elements below the pivot element to be 0
		for k := i + 1; k < n; k++ {
			factor := eqs[k][i]
			for j := i; j <= n; j++ {
				eqs[k][j] -= factor * eqs[i][j]
			}
		}
	}

	// Back Substitution
	var result = make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		result[i] = eqs[i][n]
		for j := i + 1; j < n; j++ {
			result[i] -= eqs[i][j] * result[j]
		}
	}

	return result
}
