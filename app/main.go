package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello you bastard\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
