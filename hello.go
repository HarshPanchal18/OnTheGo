package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// datatypes()

	// variables()

	// userInput()

	// loops()

	// collections()

	println(average([]float64{1.0, 2.0, 3.0}))

	x,y := f()
	println(x, y)

	println(add(1,6,3))
	println(add([]int{1,6,3,8}...))

	// Closure
	add := func(x, y int) int {
		return x + y
	}

	println(add(1,5))

	nextEven := evenGenerator()
	println(nextEven())
	println(nextEven())
	println(nextEven())

	println(factorial(5))

}