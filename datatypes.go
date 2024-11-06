package main

import "fmt"

func datatypes() {
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("1 + 1 = ", 1.0 + 1.0)

	fmt.Println(`Hello`)
	fmt.Println("Hello"[3]) // 108 - Character is represented by `byte`
	fmt.Println("Hello" + "World")
	fmt.Println(len(`Hello`))

	fmt.Println(true)
	fmt.Println(false)
}

func variables()  {
	var x string = "Hello world"
	fmt.Println(x)

	y := 5
	fmt.Println(y)

	const z string = "Constant"
	fmt.Println(z)


	var (
		a = 5
		b = 40
	)

	fmt.Println(a + b)


	const (
		aa = 5
		bb = 40
	)

	fmt.Println(aa + bb)
}

func userInput() {
	fmt.Print("Enter a number: ")

	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}