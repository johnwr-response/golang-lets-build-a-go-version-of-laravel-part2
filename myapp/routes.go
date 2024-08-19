package main

import (
	"net/http"

	"github.com/tsawler/celeritas"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.get("/", a.Handlers.Home)

	a.get("/users/login", a.Handlers.UserLogin)
	a.post("/users/login", a.Handlers.PostUserLogin)

	a.get("/upload", a.Handlers.CeleritasUpload)
	a.post("/upload", a.Handlers.PostCeleritasUpload)

	a.get("/list-fs", a.Handlers.ListFs)

	a.get("/files/upload", a.Handlers.UploadToFS)
	a.post("/files/upload", a.Handlers.PostUploadToFS)

	a.get("/delete-from-fs", a.Handlers.DeleteFromFS)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from celeritas
	a.App.Routes.Mount("/celeritas", celeritas.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
