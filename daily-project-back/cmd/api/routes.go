package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/notes/:id", app.showNoteHandler)
	router.HandlerFunc(http.MethodPut, "/v1/notes/:id", app.updateNoteHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/notes/:id", app.deleteNoteHandler)
	router.HandlerFunc(http.MethodPost, "/v1/notes", app.createNoteHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/:id/notes", app.showUserNotesHandler)
	return router
}
