package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		go f(0)
	}

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)

	/*
	* Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation.
	* With a goroutine we return immediately to the next line and don't wait for the function to complete.
	* This is why the call to the Scanln function has been included; without it the program would exit before being given the opportunity to print all the numbers.
	*/
}