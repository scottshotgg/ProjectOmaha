package main

import (
	"net/http"
	"omaha/handlers"
)

func main() {
	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/app/", handlers.AppHandler)
	http.HandleFunc("/demo/start/", handlers.DemoStartHandler)
	http.HandleFunc("/demo/stop/", handlers.DemoStartHandler)
	http.Handle("/css/", handlers.CssHandler)
	http.Handle("/bower_components/", handlers.BowerHandler)
	http.Handle("/components/", handlers.ComponentsHandler)
	http.ListenAndServe(":8080", nil)
}
