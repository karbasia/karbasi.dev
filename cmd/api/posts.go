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

type PostInput struct {
	ID        int                 `json:"-"`
	Title     string              `json:"title"`
	Slug      string              `json:"slug"`
	Headline  string              `json:"headline"`
	Content   string              `json:"content"`
	Active    bool                `json:"active"`
	PostedAt  *string             `json:"posted_at,omitzero,omitempty"`
	DeletedAt *string             `json:"deleted_at,omitzero,omitempty"`
	Tags      []store.Tag         `json:"tags"`
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

	// Verify the tags and create them if they do not exist
	for i := 0; i < len(input.Tags); i++ {
		if input.Tags[i].ID == 0 {
			app.store.Tags.Create(ctx, &input.Tags[i])
		}
	}

	post := &store.Post{
		Title:       input.Title,
		Slug:        input.Slug,
		Headline:    input.Headline,
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

	// Verify the tags and create them if they do not exist
	for i := 0; i < len(input.Tags); i++ {
		if input.Tags[i].ID == 0 {
			app.store.Tags.Create(ctx, &input.Tags[i])
		}
	}

	post := &store.Post{
		ID:        postID,
		Title:     input.Title,
		Slug:      input.Slug,
		Headline:  input.Headline,
		Content:   input.Content,
		Active:    input.Active,
		PostedAt:  input.PostedAt,
		DeletedAt: input.DeletedAt,
		Tags:      input.Tags,
	}

	err = app.store.Posts.Update(ctx, post)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, post)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	showDeleted, err := strconv.ParseBool(r.URL.Query().Get("showDeleted"))
	if err != nil {
		showDeleted = false
	}
	posts, err := app.store.Posts.GetAll(ctx, showDeleted)

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

func (app *application) getPostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	ctx := r.Context()

	post, found, err := app.store.Posts.GetBySlug(ctx, slug)
	if !found {
		app.notFound(w, r)
		return
	} else if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, post)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
