package main

import (
	"net/http"

	"tempMonitor/views/pages"
)

func (app *application) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Home().Render(r.Context(), w)
	if err != nil {

	}

}
