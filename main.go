package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello it's GoteBook"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Server Starting on port : 8888")
	err := http.ListenAndServe(":8888", mux)
	log.Fatal(err)
}
