package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/tsawler/celeritas"
	"github.com/tsawler/celeritas/filesystems"
	"github.com/tsawler/celeritas/filesystems/minioFilesystem"
	"io"
	"log"
	"mime/multipart"
	"myapp/data"
	"net/http"
	"net/url"
	"os"
)

// Handlers is the type for handlers, and gives access to Celeritas and models
type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

// Home is the handler to render the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) ListFs(w http.ResponseWriter, r *http.Request) {
	var fs filesystems.FS
	var list []filesystems.Listing

	fsType := ""
	if r.URL.Query().Get("fs-type") != "" {
		fsType = r.URL.Query().Get("fs-type")
	}

	curPath := "/"
	if r.URL.Query().Get("curPath") != "" {
		curPath = r.URL.Query().Get("curPath")
		curPath, _ = url.QueryUnescape(curPath)

		if curPath == "/" {
			curPath = ""
		}
	}

	if fsType != "" {
		switch fsType {
		case "MINIO":
			log.Println("Using MINIO for fsType")
			f := h.App.Filesystems["MINIO"].(minioFilesystem.Minio)
			fs = &f
			fsType = "MINIO"
		}

		if fs != nil {
			l, err := fs.List(curPath)
			if err != nil {
				h.App.ErrorLog.Println(err)
				return
			}
			list = l
		}

	}

	vars := make(jet.VarMap)
	vars.Set("list", list)
	vars.Set("fs_type", fsType)
	vars.Set("curPath", curPath)
	err := h.render(w, r, "list-fs", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
		return
	}
}

func (h *Handlers) UploadToFS(w http.ResponseWriter, r *http.Request) {
	fsType := r.URL.Query().Get("type")

	vars := make(jet.VarMap)
	vars.Set("fs_type", fsType)

	err := h.render(w, r, "upload", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) PostUploadToFS(w http.ResponseWriter, r *http.Request) {
	filename, err := getFileToUpload(r, "formFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadType := r.Form.Get("upload-type")
	switch uploadType {
	case "MINIO":
		fs := h.App.Filesystems["MINIO"].(minioFilesystem.Minio)
		err := fs.Put(filename, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	h.App.Session.Put(r.Context(), "flash", "File uploaded successfully!")
	http.Redirect(w, r, "/files/upload?type="+uploadType, http.StatusSeeOther)
}

func getFileToUpload(r *http.Request, fieldName string) (string, error) {
	_ = r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	dst, err := os.Create(fmt.Sprintf("./tmp/%s", header.Filename))
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("./tmp/%s", header.Filename), nil
}
