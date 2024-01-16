package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/notes/:id", app.showNoteHandler)
	router.HandlerFunc(http.MethodGet, "/v1/notes", app.listNotesHandler)
	router.HandlerFunc(http.MethodPut, "/v1/notes/:id", app.updateNoteHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/notes/:id", app.deleteNoteHandler)
	router.HandlerFunc(http.MethodPost, "/v1/notes", app.createNoteHandler)

	// routes for users
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/notes", app.showUserNotesHandler)

	return app.recoverPanic(
		app.rateLimit(router),
	)
}
