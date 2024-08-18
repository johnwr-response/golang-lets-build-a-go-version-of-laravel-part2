package celeritas

import (
	"fmt"
	"github.com/tsawler/celeritas/filesystems"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func (c *Celeritas) UploadFile(r *http.Request, destination, field string, fs filesystems.FS) error {
	fileName, err := c.getFileToUpload(r, field)
	if err != nil {
		c.ErrorLog.Println(err)
		return err
	}

	if fs != nil {
		// Uploading to a remote filesystem
		err = fs.Put(fileName, destination)
		if err != nil {
			c.ErrorLog.Println(err)
			return err
		}
	} else {
		// Uploading to a (server) local filesystem
		err = os.Rename(fileName, fmt.Sprintf("%s/%s", destination, path.Base(fileName)))
		if err != nil {
			c.ErrorLog.Println(err)
			return err
		}
	}
	return nil
}

func (c *Celeritas) getFileToUpload(r *http.Request, fieldName string) (string, error) {
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
