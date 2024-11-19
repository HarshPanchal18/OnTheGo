package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "The maximum value") // Arg name, Default value, and Usage.
	flag.Parse()
	fmt.Println(rand.Intn(*maxp)) // Generate a number in 0..max

	// Any additional non-flag arguments can be retrieved with flag.Args() which returns a []string.
}
