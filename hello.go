package main

import "fmt"

func main() {
	fmt.Println(string(sha1Hash([]byte("foo"))))
}
