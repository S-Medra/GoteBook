package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ": 8888", "PORT")
	flag.Parse()

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	log.Printf("Server Starting on port%s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
