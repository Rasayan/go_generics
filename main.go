package main

import "fmt"

// Declearing a type contraint
type Number interface {
	int64 | float64
}

// Writing Non-Generic function first
func sumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

// Generic function to handle any kind of incoming input type
func sumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic sums: %v and %v\n",
		sumInts(ints),
		sumFloats(floats))

	fmt.Printf("Generic sums: %v and %v\n",
		sumIntsOrFloats[string, int64](ints),
		sumIntsOrFloats[string, float64](floats))
}
