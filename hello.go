package main

import "net/http"

func main() {
	http.HandleFunc("/hello", hello) // handle a URL route (/hello) by calling the given function.

	http.Handle("/", http.FileServer(http.Dir("."))) // a File server that expose current directory.

	http.ListenAndServe(":9000", nil)
}
