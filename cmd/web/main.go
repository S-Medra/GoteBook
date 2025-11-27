package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	log.Println("Server Starting on port : 8888")
	err := http.ListenAndServe(":8888", mux)
	log.Fatal(err)
}
