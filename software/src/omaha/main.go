package main

import (
	"net/http"
	"omaha/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/demo/start/", handlers.DemoStartHandler)
	http.HandleFunc("/demo/stop/", handlers.DemoStartHandler)
	http.Handle("/css/", handlers.CssHandler)
	http.ListenAndServe(":8080", nil)
}
