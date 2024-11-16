package main

import "fmt"

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
