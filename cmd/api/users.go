package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/karbasia/karbasi.dev/internal/password"
	"github.com/karbasia/karbasi.dev/internal/request"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/validator"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FullName  string              `json:"full_name"`
		Email     string              `json:"email"`
		Password  string              `json:"password"`
		DeletedAt *string             `json:"deleted_at,omitzero"`
		Validator validator.Validator `json:"-"`
	}

	ctx := r.Context()

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	_, found, err := app.store.Users.GetByEmail(ctx, input.Email)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	input.Validator.CheckField(input.FullName != "", "full_name", "Name is required")
	input.Validator.CheckField(len(input.FullName) <= 100, "full_name", "Name must be 100 or fewer characters")

	input.Validator.CheckField(input.Email != "", "email", "Email is required")
	input.Validator.CheckField(validator.Matches(input.Email, validator.RgxEmail), "email", "Must be a valid email address")
	input.Validator.CheckField(!found, "email", "Email is already in use")

	input.Validator.CheckField(input.Password != "", "password", "Password is required")
	input.Validator.CheckField(len(input.Password) >= 8, "password", "Password is too short")
	input.Validator.CheckField(len(input.Password) <= 72, "password", "Password is too long")
	input.Validator.CheckField(validator.NotIn(input.Password, password.CommonPasswords...), "password", "Password is too common")

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	hashedPassword, err := password.Hash(input.Password)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	user := &store.User{
		FullName:       input.FullName,
		Email:          input.Email,
		HashedPassword: hashedPassword,
	}
	err = app.store.Users.Create(ctx, user)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusCreated, user)
}

func (app *application) getUserByID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx := r.Context()

	user, found, err := app.store.Users.GetByID(ctx, id)
	if !found {
		app.notFound(w, r)
		return
	} else if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, user)
	if err != nil {
		app.serverError(w, r, err)
	}

}

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := app.store.Users.GetAll(ctx)
	if err != nil {
		app.serverError(w, r, err)
	}
	err = response.JSON(w, http.StatusOK, users)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := contextGetAuthenticatedUser(r)

	userCore := store.UserCore{
		ID:       user.ID,
		FullName: user.FullName,
	}

	err := response.JSON(w, http.StatusOK, userCore)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
