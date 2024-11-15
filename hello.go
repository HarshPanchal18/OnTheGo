package main

import (
	"errors"
	"fmt"
)

func main() {
	// stringpkg()
	// ospkg()
	err := errors.New("not implemented")
	fmt.Printf("err: %v\n", err.Error())
}
