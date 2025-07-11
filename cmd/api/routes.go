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

	r.Route("/users", func(r chi.Router) {
		r.Get("/", app.getAllUsers)
		r.Post("/", app.createUser)
		r.Get("/{id}", app.getUserByID)
		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthenticatedUser)
			r.Get("/me", app.getCurrentUser)
		})
	})

	r.Route("/posts", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthenticatedUser)
			r.Post("/", app.createPost)
			r.Patch("/{id}", app.updatePost)
		})
		r.Get("/", app.getAllPosts)
		r.Get("/{slug}", app.getPostBySlug)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthenticatedUser)
			r.Post("/", app.createTag)
		})
		r.Get("/", app.getAllTags)
		r.Get("/counts", app.getAllTagsWithPostCount)
		r.Get("/{tag}", app.getAllPostsByTag)
	})

	r.Route("/files", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(app.requireAuthenticatedUser)
			r.Get("/", app.getAllFiles)
			r.Post("/", app.createFile)
		})
		r.Get("/{name}", app.getFileByName)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", app.handleLogin)
		r.Post("/refresh", app.handleRefresh)
	})

	return r
}
