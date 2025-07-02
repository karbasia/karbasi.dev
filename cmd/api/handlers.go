package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/karbasia/karbasi.dev/internal/password"
	"github.com/karbasia/karbasi.dev/internal/request"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/validator"

	"github.com/pascaldekloe/jwt"
)

type AuthResponse struct {
	AccessToken        string          `json:"access_token"`
	AccessTokenExpiry  string          `json:"access_token_expiry"`
	RefreshToken       string          `json:"refresh_token,omitempty"`
	RefreshTokenExpiry string          `json:"refresh_token_expiry,omitempty"`
	UserInfo           *store.UserCore `json:"user_info"`
}

func (app *application) handleLogin(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email     string              `json:"email"`
		Password  string              `json:"password"`
		Validator validator.Validator `json:"-"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	ctx := r.Context()
	user, found, err := app.store.Users.GetByEmail(ctx, input.Email)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	input.Validator.CheckField(input.Email != "", "email", "Email is required")
	input.Validator.CheckField(found, "Email", "email address could not be found")

	if found {
		passwordMatches, err := password.Matches(input.Password, user.HashedPassword)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		input.Validator.CheckField(input.Password != "", "password", "Password is required")
		input.Validator.CheckField(passwordMatches, "password", "Password is incorrect")
	}

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	userID := strconv.Itoa(user.ID)
	accessExpiry := time.Now().Add(30 * time.Minute)
	accessToken, err := app.generateToken(accessExpiry, app.config.jwt.accessSecretKey, userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	refreshExpiry := time.Now().Add(7 * 24 * time.Hour)
	refreshToken, err := app.generateToken(refreshExpiry, app.config.jwt.refreshSecretKey, userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := AuthResponse{
		AccessToken:        string(accessToken),
		AccessTokenExpiry:  accessExpiry.Format(time.RFC3339),
		RefreshToken:       string(refreshToken),
		RefreshTokenExpiry: refreshExpiry.Format(time.RFC3339),
		UserInfo:           &store.UserCore{ID: user.ID, FullName: user.FullName},
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) handleRefresh(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	claims, err := jwt.HMACCheck([]byte(input.RefreshToken), []byte(app.config.jwt.refreshSecretKey))
	if err != nil {
		app.invalidRefreshToken(w, r)
		return
	}

	if !claims.Valid(time.Now()) {
		app.invalidRefreshToken(w, r)
		return
	}

	if claims.Issuer != app.config.baseURL {
		app.invalidRefreshToken(w, r)
		return
	}

	if !claims.AcceptAudience(app.config.baseURL) {
		app.invalidRefreshToken(w, r)
		return
	}

	ctx := r.Context()
	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		app.invalidRefreshToken(w, r)
		return
	}
	user, _, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	accessExpiry := time.Now().Add(30 * time.Minute)
	accessToken, err := app.generateToken(accessExpiry, app.config.jwt.accessSecretKey, claims.Subject)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := AuthResponse{
		AccessToken:       string(accessToken),
		AccessTokenExpiry: accessExpiry.Format(time.RFC3339),
		UserInfo:          &store.UserCore{ID: user.ID, FullName: user.FullName},
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) generateToken(expiry time.Time, secret, subject string) ([]byte, error) {
	var claims jwt.Claims
	claims.Subject = subject

	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(expiry)

	claims.Issuer = app.config.baseURL
	claims.Audiences = []string{app.config.baseURL}

	return claims.HMACSign(jwt.HS256, []byte(secret))
}
