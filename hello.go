package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// datatypes()

	// variables()

	// userInput()

	// loops()

	// collections()

	/*println(average([]float64{1.0, 2.0, 3.0}))

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

	println(factorial(5))*/

	/*
	defer second() // defer - moves the call to second to the end of the function. often used when resources need to be freed in some way.
	first()
	*/

	/*
	For example when we open a file we need to make sure to close it later. With defer:
		f, _ := os.Open(filename)
		defer f.Close()

		This has 3 advantages:
		(1) it keeps our Close call near our Open call so its easier to understand,
		(2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them and
		(3) deferred functions are run even if a run-time panic occurs.
	*/

	/*
	panic("PANIC")
	str := recover() // unreachable code
	println(str)
	*/

	// recover() stops the panic and returns the value that was passed to the call to panic.
	/*defer func() {
		str := recover()
		fmt.Println("Something wrong happened: ", str)
	}()

	panic("PANIC")*/

	/*x := 5
	zero(&x)
	fmt.Println(x)

	// Another way to get a pointer.
	xPtr := new(int) // new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	one(xPtr)
	fmt.Println(*xPtr)
	*/

	// c := new(Circle) // This allocates memory for all the fields, sets each of them to their zero value and returns a pointer.
	c := Circle { x: 0, y: 0, r: 10 } // Or we can leave off the field names if we know the order they were defined
	fmt.Println(c.x, c.y, c.r)

	fmt.Println(circleArea(c))
	fmt.Println(c.area())

	rect := Rectangle{ 2, 2, 4, 4}
	fmt.Println(rect.area())
}