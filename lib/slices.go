package _slices

import (
	"aoc2023/assert"
	"strconv"
	"strings"
)

// Compare two slices of comparable items, index is also considered.
func Compare[T comparable](a, b []T) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) is not a valid slice of s.
// Same as slices.Delete() but does not mutate the original.
func Delete[T any](s []T, i, j int) []T {
	_ = s[i:j:len(s)] // bounds check

	if i == j {
		return s
	}

	var newSlice []T
	for idx := range s {
		if i <= idx && idx <= j-1 {
			continue
		}
		newSlice = append(newSlice, s[idx])
	}
	return newSlice
}

// Panics on error.
func StrToInt(value []string) []int {
	var result []int
	for _, v := range value {
		num, err := strconv.Atoi(strings.Trim(v, " "))
		assert.Assert(err == nil, "error converting string to int!")
		result = append(result, num)
	}
	return result
}

// Panics on error.
func StrToFloat(value []string) []float64 {
	var result []float64
	for _, v := range value {
		num, err := strconv.ParseFloat(strings.Trim(v, " "), 64)
		assert.Assert(err == nil, "error converting string to int!")
		result = append(result, num)
	}
	return result
}

// Filter slice using callback function, returning a new slice.
// Callback function receives the index and the item as parameter.
func Filter[T any](slice []T, fn func(int, T) bool) []T {
	var result []T
	for i, item := range slice {
		if fn(i, item) {
			result = append(result, item)
		}
	}
	return result
}
