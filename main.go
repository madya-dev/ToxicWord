package main

import (
	"log"
	"net/http"
	"toxicword/handler"
)

func main() {
	// create mux
	mux := http.NewServeMux()
	// routing with mux
	mux.HandleFunc("/", handler.MainHandler)
	mux.HandleFunc("/say", handler.SayHandler)
	// directory manipulation
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// port
	port := ":8080"
	// show on log
	log.Printf("Starting on http://localhost%s", port)
	// run server
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)

}