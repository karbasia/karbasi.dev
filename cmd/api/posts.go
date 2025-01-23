package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/karbasia/karbasi.dev/internal/request"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/validator"
)

type PostInput struct {
	ID        int                 `json:"-"`
	Title     string              `json:"title"`
	Slug      string              `json:"slug"`
	Content   string              `json:"content"`
	Active    int                 `json:"active"`
	PostedAt  *time.Time          `json:"posted_at"`
	Tags      []store.TagCore     `json:"tags"`
	Validator validator.Validator `json:"-"`
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	input := PostInput{}
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

	ctx := r.Context()
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
		Tags:        input.Tags,
		CreatedByID: user.ID,
		CreatedBy:   userCore,
	}

	err = app.store.Posts.Create(ctx, post)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusCreated, post)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) updatePost(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	postID, err := strconv.Atoi(param)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	input := PostInput{}
	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx := r.Context()
	post := &store.Post{
		ID:       postID,
		Title:    input.Title,
		Slug:     input.Slug,
		Content:  input.Content,
		Active:   input.Active,
		PostedAt: input.PostedAt,
		Tags:     input.Tags,
	}

	err = app.store.Posts.Update(ctx, post)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusCreated, post)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
	}
}
