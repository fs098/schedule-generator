package main

import (
	"math"
)

// Filter == filter in haskell
func Filter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, val := range input {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

// FilterByMin Filters all the elements of the given input that have the minimum
// amount of something, as defined by f
func FilterByMin[T any](input []T, f func(T) int) ([]T, int) {
	min := math.MaxInt
	for _, i := range input {
		val := f(i)
		if min > val {
			min = val
		}
	}

	result := Filter(input, func(a T) bool {
		return f(a) == min
	})
	return result, min
}
