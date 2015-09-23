package main

import (
	"net/http"
	"omaha/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.Handle("/css/", handlers.CssHandler)
	http.ListenAndServe(":8080", nil)
}
