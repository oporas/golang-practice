package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggo doggo doggo")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "mew mew mew")
}

func main() {
	http.HandleFunc("/dog/", d) //Extra / in the end will catch all paths that starts with /dog
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
