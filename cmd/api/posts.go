package main

import (
	"net/http"
	"time"

	"github.com/karbasia/karbasi.dev/internal/request"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/validator"
)

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string              `json:"title"`
		Slug      string              `json:"slug"`
		Content   string              `json:"content"`
		Active    int                 `json:"active"`
		PostedAt  *time.Time          `json:"posted_at"`
		Validator validator.Validator `json:"-"`
	}

	ctx := r.Context()

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Title != "", "title", "A title is required")
	input.Validator.CheckField(input.Slug != "", "slug", "A slug is required")
	input.Validator.CheckField(input.Content != "", "content", "The post must require some content")

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	_, found, err := app.store.Posts.GetBySlug(ctx, input.Slug)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	input.Validator.CheckField(!found, "slug", "A post with the same slug already exists.")

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	user := contextGetAuthenticatedUser(r)

	userCore := store.UserCore{
		ID:       user.ID,
		FullName: user.FullName,
	}

	post := &store.Post{
		Title:       input.Title,
		Slug:        input.Slug,
		Content:     input.Content,
		Active:      input.Active,
		PostedAt:    input.PostedAt,
		CreatedByID: user.ID,
		CreatedBy:   userCore,
	}

	app.store.Posts.Create(ctx, post)
	err = response.JSON(w, http.StatusCreated, post)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := app.store.Posts.GetAll(ctx)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, posts)
	if err != nil {
		app.serverError(w, r, err)
	}
}
