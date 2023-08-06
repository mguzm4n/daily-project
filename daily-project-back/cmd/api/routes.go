package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/notes", app.showNoteHandler)
	router.HandlerFunc(http.MethodPost, "/v1/notes", app.createNoteHandler)
	return router
}
