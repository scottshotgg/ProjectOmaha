package main

import (
	"net/http"
	"omaha/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/demo/write/", handlers.DemoWriteHandler)
	http.Handle("/css/", handlers.CssHandler)
	http.ListenAndServe(":8080", nil)
}
