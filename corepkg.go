package main

import (
	"fmt"
	"strings"
)

func stringpkg() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"1", "a"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e-f", "-"),
		strings.ToUpper("test"),
	)

	// To convert a string to a slice of bytes (and vice-versa)
	arr := []byte("test")
	fmt.Println(arr)

	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
}