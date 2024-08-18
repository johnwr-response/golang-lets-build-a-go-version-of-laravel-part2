package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/tsawler/celeritas"
	"github.com/tsawler/celeritas/filesystems"
	"github.com/tsawler/celeritas/filesystems/minioFilesystem"
	"github.com/tsawler/celeritas/filesystems/s3Filesystem"
	"github.com/tsawler/celeritas/filesystems/sFtpFilesystem"
	"github.com/tsawler/celeritas/filesystems/webdavFilesystem"
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

	//err := h.App.UploadFile(r, "uploaded-file", h.App.Minio)
	//if err != nil {
	//}

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

	_, fsTypeExists := h.App.Filesystems[fsType]
	if fsTypeExists {
		switch fsType {
		case "MINIO":
			log.Println("Using MINIO for fsType")
			f := h.App.Filesystems["MINIO"].(minioFilesystem.Minio)
			fs = &f
			fsType = "MINIO"
		case "SFTP":
			log.Println("Using SFTP for fsType")
			f := h.App.Filesystems["SFTP"].(sFtpFilesystem.SFTP)
			fs = &f
			fsType = "SFTP"
		case "WEBDAV":
			log.Println("Using WEBDAV for fsType")
			f := h.App.Filesystems["WEBDAV"].(webdavFilesystem.WebDAV)
			fs = &f
			fsType = "WEBDAV"
		case "S3":
			log.Println("Using S3 for fsType")
			f := h.App.Filesystems["S3"].(s3Filesystem.S3)
			fs = &f
			fsType = "S3"
		}

		if fs != nil {
			l, err := fs.List(curPath)
			if err != nil {
				h.App.ErrorLog.Println(err)
				return
			}
			list = l
		}
	} else {
		log.Println(fmt.Sprintf("Filesystem %s not enabled", fsType))
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
	case "SFTP":
		fs := h.App.Filesystems["SFTP"].(sFtpFilesystem.SFTP)
		err := fs.Put(filename, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "WEBDAV":
		fs := h.App.Filesystems["WEBDAV"].(webdavFilesystem.WebDAV)
		err := fs.Put(filename, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "S3":
		fs := h.App.Filesystems["S3"].(s3Filesystem.S3)
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

func (h *Handlers) DeleteFromFS(w http.ResponseWriter, r *http.Request) {
	var fs filesystems.FS
	fsType := r.URL.Query().Get("fs_type")
	item := r.URL.Query().Get("file")

	curPath := "/"
	if r.URL.Query().Get("curPath") != "" {
		curPath = r.URL.Query().Get("curPath")
		curPath, _ = url.QueryUnescape(curPath)

		if curPath == "/" {
			curPath = ""
		}
	}

	switch fsType {
	case "MINIO":
		f := h.App.Filesystems["MINIO"].(minioFilesystem.Minio)
		fs = &f
	case "SFTP":
		f := h.App.Filesystems["SFTP"].(sFtpFilesystem.SFTP)
		fs = &f
	case "WEBDAV":
		f := h.App.Filesystems["WEBDAV"].(webdavFilesystem.WebDAV)
		fs = &f
	case "S3":
		f := h.App.Filesystems["S3"].(s3Filesystem.S3)
		fs = &f
	}

	if fs != nil {
		deleted := fs.Delete([]string{item})
		if deleted {
			h.App.Session.Put(r.Context(), "flash", fmt.Sprintf("%s was deleted", item))
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/list-fs?fs-type=%s&curPath=%s", fsType, curPath), http.StatusSeeOther)
}
