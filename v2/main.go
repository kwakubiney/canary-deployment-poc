package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello i am v2\n")
}

func main() {
	http.HandleFunc("/v2", hello)
	http.ListenAndServe(":8080", nil)
}