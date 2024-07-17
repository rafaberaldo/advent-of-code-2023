package day09

import (
	"testing"
)

func TestSameElements(t *testing.T) {
	values := []int{1, 1, 1, 1, 1}
	result := sameElements(values)
	if !result {
		t.Fatalf(`sameElements(%#v) = %#v, expected %#v`, values, result, true)
	}
}

func TestNotSameElements(t *testing.T) {
	values := []int{1, 1, 1, 1, 0}
	result := sameElements(values)
	if result {
		t.Fatalf(`sameElements(%#v) = %#v, expected %#v`, values, result, false)
	}
}

func TestNotSameElements2(t *testing.T) {
	values := []int{0, 1, 1, 1, 1}
	result := sameElements(values)
	if result {
		t.Fatalf(`sameElements(%#v) = %#v, expected %#v`, values, result, false)
	}
}

func TestNotSameElements3(t *testing.T) {
	values := []int{0, 1, 2, 3, 4}
	result := sameElements(values)
	if result {
		t.Fatalf(`sameElements(%#v) = %#v, expected %#v`, values, result, false)
	}
}
