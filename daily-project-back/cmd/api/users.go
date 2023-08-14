package main

import (
	"daily-project/internal/data"
	"errors"
	"net/http"
)

func (app *application) showUserNotesHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := app.readIntIdParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	notes, err := app.models.Users.GetNotes(uid)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"notes": notes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
