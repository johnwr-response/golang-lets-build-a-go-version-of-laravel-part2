package webdavFilesystem

import "github.com/tsawler/celeritas/filesystems"

type WebDAV struct {
	Host string
	User string
	Pass string
}

// Put transfers a file to the remote file system
func (s *WebDAV) Put(fileName, folder string) error {
	return nil
}

// List returns a listing of all files in the remote bucket with the given prefix, except for files named with a leading .
func (s *WebDAV) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	return listing, nil
}

// Delete removes one or more files from the remote filesystem
func (s *WebDAV) Delete(itemsToDelete []string) bool {
	return true
}

// Get pulls a file from the remote file system and saves it somewhere on our server
func (s *WebDAV) Get(destination string, items ...string) error {
	return nil
}
