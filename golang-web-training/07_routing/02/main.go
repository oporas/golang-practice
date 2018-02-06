package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggo doggo doggo")
}

type hotcat int

func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "mew mew mew")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d) //Extra / in the end will catch all paths that starts with /dog
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
