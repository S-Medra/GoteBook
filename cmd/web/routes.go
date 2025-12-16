package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/note/view", app.noteView)
	mux.HandleFunc("/note/create", app.noteCreate)

	return app.logRequest(secureHeaders(mux))
}
