package main

import (
	"errors"
	"fmt"
	"github.com/S-Medra/GoteBook/internal/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	notes, err := app.notes.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "home.html", &templateData{
		Notes: notes,
	})
}
func (app *application) noteView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	app.render(w, http.StatusOK, "view.html", &templateData{
		Note: note,
	})
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
	}
	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.notes.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/note/view?id=%d", id), http.StatusSeeOther)
}
