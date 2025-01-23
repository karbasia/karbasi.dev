package main

import (
	"net/http"

	"github.com/karbasia/karbasi.dev/internal/request"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/validator"
)

type TagInput struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
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

	tags, err := app.store.Tags.GetAll(ctx)
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
