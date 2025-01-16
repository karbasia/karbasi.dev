package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.NotFound(app.notFound)
	r.MethodNotAllowed(app.methodNotAllowed)

	r.Use(app.logAccess)
	r.Use(app.recoverPanic)
	r.Use(app.authenticate)

	r.Route("/posts", func(r chi.Router) {

	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", app.getAllUsers)
		r.Post("/", app.createUser)
		r.Get("/{id}", app.getUserByID)
	})

	r.Get("/status", app.status)

	r.Post("/authentication-tokens", app.createAuthenticationToken)

	r.Group(func(r chi.Router) {
		r.Use(app.requireAuthenticatedUser)

		r.Get("/protected", app.protected)
	})

	return r
}
