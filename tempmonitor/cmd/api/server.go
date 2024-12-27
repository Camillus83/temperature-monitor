package main

import "net/http"

func (app *application) serve() error {
	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
		// Handler: app.routes(),
	}

	app.logger.Info("Starting server")

	srv.ListenAndServe()

	return nil
}
