package main

import (
	"daily-project/internal/validator"
	"encoding/json"
	"net/http"
	"time"
)

var connStr = "postgres://greenlight:pa55word@localhost/greenlight"

func (app *application) createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var noteReqBody struct {
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}

	err := json.NewDecoder(r.Body).Decode(&noteReqBody)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	v := validator.New()

	v.Check(len(noteReqBody.Content) > 0, "content", "can't be empty string")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//fmt.Fprintf(w, "%+v\n", input)
	//fmt.Fprintf(w, noteReqBody)

}

// func (app *application) showNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	nid, err := app.readIntIdParam(r)

// 	if err != nil {
// 		http.NotFound(w, r)
// 		return
// 	}

// }
