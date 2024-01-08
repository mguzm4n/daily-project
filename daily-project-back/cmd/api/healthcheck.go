package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Write text to response.
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	// Write JSON to response.
	// json := `{
	// 	"status": "available"
	// 	"environment": %q,
	// 	"version": %q
	// }`

	// json = fmt.Sprintf(json, app.config.env, version)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(json))

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	env := envelope{
		"status":      "available",
		"system_info": data,
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
