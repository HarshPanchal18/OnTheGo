package main

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, shape := range shapes {
		area += shape.area()
	}

	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64

	for _, shape := range m.shapes {
		area += shape.area()
	}

	return area
}