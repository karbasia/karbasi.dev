package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/karbasia/karbasi.dev/internal/request"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/validator"
)

type TagInput struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	DeletedAt *string             `json:"deleted_at,omitzero"`
	Validator validator.Validator `json:"-"`
}

func (app *application) createTag(w http.ResponseWriter, r *http.Request) {
	input := TagInput{}
	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Name != "", "name", "A name is required")

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	tag := &store.Tag{Name: input.Name}

	ctx := r.Context()
	err = app.store.Tags.Create(ctx, tag)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusCreated, tag)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getAllTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var tags []store.Tag

	showDeleted, err := strconv.ParseBool(r.URL.Query().Get("showDeleted"))
	if err != nil {
		showDeleted = false
	}

	tags, err = app.store.Tags.GetAll(ctx, showDeleted)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, tags)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getAllPostsByTag(w http.ResponseWriter, r *http.Request) {
	tag := chi.URLParam(r, "tag")

	ctx := r.Context()

	posts, err := app.store.Posts.GetAllByTag(ctx, tag)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, posts)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getAllTagsWithPostCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tags, err := app.store.Tags.GetAllByPostCount(ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, tags)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
