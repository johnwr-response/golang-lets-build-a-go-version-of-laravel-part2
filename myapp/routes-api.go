package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	return r
}
