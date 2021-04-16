package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/hello", hello)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", mux)
}
