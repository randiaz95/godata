package main

import (
	"fmt"
)


// Gets a func and a slice of ints and maps that input to an output.
func IntMap(fn func(input int) int, input []int) []int {
	var output []int
	
	for _, value := range input {
		output = append(output, fn(value))	
	}
	
	return output
}

// Gets a func and a slice of float64s and maps that input to an output
func FloatMap(fn func(input float64) float64, input []float64) []float64 {
	var output []float64
	
	for _, value := range input {
		output = append(output, fn(value))
	}
	
	return output
}


func main() {
	
	// Test Increment function on slice of ints.
	fmt.Println(IntMap(incrementInt, []int{1, 2, 3}))
	
	// Test Increment function on slice of floats.
	fmt.Println(FloatMap(incrementFloat, []float64{1.0, 2.0, 3.0}))
	
}

func incrementInt(a int) int {
	return a+1
}
func incrementFloat(a float64) float64 {
	return a+0.01
}
