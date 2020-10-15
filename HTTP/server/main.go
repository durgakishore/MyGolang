package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func main() {

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe("localhost:8080", nil)
}
