package main

import "math"

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (rect *Rectangle) area() float64 {
	length := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	width := distance(rect.x1, rect.y1, rect.x2, rect.y1)

	return length * width
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}