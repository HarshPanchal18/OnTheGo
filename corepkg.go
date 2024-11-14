package main

import (
	"fmt"
	"os"
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

func ospkg() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening test file.")
		return
	}
	fmt.Println("File name: ", file.Name())
	defer file.Close()

	fileContent, error := os.ReadFile("test.txt")
	if error != nil {
		fmt.Println("Error reading test file")
		return
	}
	fmt.Println("File content: ", string(fileContent))

	// get the file-size.
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error opening test file.")
		return
	}
	fmt.Println("File size in bytes: ", stat.Size())

	// Fetch ENV variables through a key.
	fmt.Println(strings.Split(os.Getenv("path"),";"))

}
