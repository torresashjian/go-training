package main

import "fmt"

type Number interface{
	int64 | float64
}

// SumInst adds the value of m
func SumInts (m map[string]int64) int64{
	var s int64
	for _, v := range m {
		s +=v
	}
	return s
}

// SumFloats adds the value of m
func SumFloats (m map[string]float64) float64{
	var s float64
	for _, v := range m {
		s +=v
	}
	return s
}

func SumIntsOrFloats [K comparable, V int64 | float64] (m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers [K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main () {
	// Initialize a map for the integer values
	intMap := map[string]int64{
		"first": 32,
		"second": 12,
	}
	
	// Initialize a map for the float values
	intFloat := map[string]float64{
		"first": 34.54,
		"second": 15.53,
	}
	
	fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(intMap),
        SumFloats(intFloat))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](intMap),
		SumIntsOrFloats[string, float64](intFloat))

	fmt.Printf("Generic Sums no types: %v and %v\n",
		SumIntsOrFloats(intMap),
		SumIntsOrFloats(intFloat))

	fmt.Printf("Generic Sums Number: %v and %v\n",
		SumNumbers(intMap),
		SumNumbers(intFloat))
}