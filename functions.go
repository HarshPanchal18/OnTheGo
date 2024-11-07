package main

import "fmt"

func average(nums []float64) float64 {
	// panic("Not implemented")
	total := 0.0

	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}

// Returning multiple values
func f() (int, int) {
	return 5, 6
}

// Variadic functions
func add(arguments ...int) int {
	total := 0
	for _, v := range arguments {
		total += v
	}
	return total
}

func evenGenerator() func() int {
	i := int(0)

	return func() (value int) { // parameter signature should be match with parent method signature.
		value = i
		i += 2
		return // `value` is returned here.
	}
}

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}