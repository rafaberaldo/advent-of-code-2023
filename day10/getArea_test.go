package day10

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	values := []Point{
		{1, 6},
		{3, 1},
		{7, 2},
		{4, 4},
		{8, 6},
	}
	expect := 18
	result := getArea(values)
	if result != expect {
		t.Fatalf(`getArea(%#v) = %#v, expected %#v`, values, result, expect)
	}
}
