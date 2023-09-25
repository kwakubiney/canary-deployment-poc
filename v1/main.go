package main

import (
	"fmt"
	"net/http"
)

func hi(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	 }
	fmt.Fprintf(w, "hi, I am v1\n")
}

func main() {
	http.HandleFunc("/", hi)
	http.ListenAndServe(":8080", nil)
}