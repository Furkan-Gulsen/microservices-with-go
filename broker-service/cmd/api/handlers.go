package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello from the Broker!",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
