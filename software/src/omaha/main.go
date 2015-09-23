package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
