package main

import "fmt"

func main() {
	fmt.Println(hashedBytes([]byte("test")))
	fmt.Println(hashedBytes([]byte("test2")))

	hash1, err := getHash("test.txt")
	if err != nil { return }

	hash2, err := getHash("test2.txt")
	if err != nil { return }

	fmt.Println(hash1, hash2, hash1 == hash2)
}
