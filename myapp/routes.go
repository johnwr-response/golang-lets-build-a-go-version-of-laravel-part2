package main

import (
	"github.com/tsawler/celeritas/filesystems/minioFilesystem"
	"log"
	"net/http"

	"github.com/tsawler/celeritas"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.get("/", a.Handlers.Home)

	a.get("/test-minio", func(w http.ResponseWriter, r *http.Request) {
		f := a.App.Filesystems["MINIO"].(minioFilesystem.Minio)

		files, err := f.List("")
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("files:", files)

		for _, file := range files {
			log.Println(file.Key)
		}
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from celeritas
	a.App.Routes.Mount("/celeritas", celeritas.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
