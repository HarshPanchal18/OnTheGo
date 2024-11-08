package main

import "math"

type Circle struct {
	x, y, r float64
}

func circleArea(circle Circle) float64 {
	return math.Pi * circle.r * circle.r
}

// Methods - Similar to extension methods for structs
func (circle *Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}