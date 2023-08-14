package main

import (
	"daily-project/internal/data"
	"daily-project/internal/validator"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var connStr = "postgres://greenlight:pa55word@localhost/greenlight"

func (app *application) createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var noteReqBody struct {
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&noteReqBody)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	note := &data.Note{
		UserID:  1,
		Content: noteReqBody.Content,
	}

	v := validator.New()

	data.ValidateNote(v, note)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Notes.Insert(note)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send location of new resource URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", note.ID))

	// Write a JSON response with a 201 Created status code, the note data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"note": note}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showNoteHandler(w http.ResponseWriter, r *http.Request) {
	nid, err := app.readIntIdParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	note, err := app.models.Notes.Get(nid)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"note": note}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	nid, err := app.readIntIdParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	note, err := app.models.Notes.Get(nid)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	var input struct {
		Content string `json:"content"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	note.Content = input.Content

	v := validator.New()
	data.ValidateNote(v, note)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Notes.Update(note)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Write the updated movie record in a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"note": note}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	nid, err := app.readIntIdParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = app.models.Notes.Delete(nid)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Write the updated movie record in a JSON response.
	strRes := fmt.Sprintf("note with id = %d deleted correctly", nid)
	err = app.writeJSON(w, http.StatusOK, envelope{"msg": strRes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
