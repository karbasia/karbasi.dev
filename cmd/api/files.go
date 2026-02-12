package main

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/karbasia/karbasi.dev/internal/pagination"
	"github.com/karbasia/karbasi.dev/internal/response"
	"github.com/karbasia/karbasi.dev/internal/store"
)

func (app *application) createFile(w http.ResponseWriter, r *http.Request) {
	multiPartFile, fileheader, err := r.FormFile("file")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer multiPartFile.Close()

	fileData, err := io.ReadAll(multiPartFile)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	filename := fileheader.Filename

	ctx := r.Context()

	_, found, err := app.store.Files.GetByName(ctx, fileheader.Filename)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Append the current time to the file name
	if found {
		extension := path.Ext(filename)
		extIndex := strings.LastIndex(filename, extension)
		fileBase := filename[:extIndex]
		filename = fmt.Sprintf("%s-%d%s", fileBase, time.Now().Unix(), extension)
	}

	file := &store.File{
		Name:    filename,
		Content: fileData,
	}

	err = app.store.Files.Create(ctx, file)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// Do not return the entire content
	file.Content = nil

	err = response.JSON(w, http.StatusCreated, file)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getFileByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	ctx := r.Context()
	file, found, err := app.store.Files.GetByName(ctx, name)
	if !found {
		app.notFound(w, r)
		return
	} else if err != nil {
		app.serverError(w, r, err)
		return
	}
	_, err = w.Write(file.Content)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getAllFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := pagination.FromRequest(r)
	result, err := app.store.Files.GetAll(ctx, params)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, result.Items, result.Pagination)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
