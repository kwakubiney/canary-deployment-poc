package main

import (
	"fmt"
	"net/http"
)

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi, I am v1\n")
}

func main() {
	http.HandleFunc("/v1", hi)
	http.ListenAndServe(":8080", nil)
}