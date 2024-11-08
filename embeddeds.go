package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hello, I am ", p.Name)
}

type Android struct {
	Person // anonymous field (embedded types for is-a relationship)
	Model string
}