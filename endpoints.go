package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(response http.ResponseWriter, request *http.Request) {

	response.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(response,
		`<doctype html>
		<html>
			<head><title>Hello World!</title></head>
			<body>
				Hello World!
			</body>
		</html>`,
	)

	fmt.Println("Page Loaded @ http://localhost:9000/hello")

}