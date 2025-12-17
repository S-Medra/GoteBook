package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/S-Medra/GoteBook/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	notes, err := app.notes.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Notes = notes

	app.render(w, http.StatusOK, "home.html", data)
}
func (app *application) noteView(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	note, err := app.notes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Note = note

	app.render(w, http.StatusOK, "view.html", data)
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for writing a new note"))
}

func (app *application) noteCreatePost(w http.ResponseWriter, r *http.Request) {

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.notes.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/note/view/%d", id), http.StatusSeeOther)
}
