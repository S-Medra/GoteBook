package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello it's GoteBook"))
}
func noteView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a note"))
}
func noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a note"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	log.Println("Server Starting on port : 8888")
	err := http.ListenAndServe(":8888", mux)
	log.Fatal(err)
}
