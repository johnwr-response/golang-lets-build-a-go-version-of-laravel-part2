package sFtpFilesystem

import "github.com/tsawler/celeritas/filesystems"

type SFTP struct {
	Host string
	User string
	Pass string
	Port string
}

// Put transfers a file to the remote file system
func (s *SFTP) Put(fileName, folder string) error {
	return nil
}

// List returns a listing of all files in the remote bucket with the given prefix, except for files named with a leading .
func (s *SFTP) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	return listing, nil
}

// Delete removes one or more files from the remote filesystem
func (s *SFTP) Delete(itemsToDelete []string) bool {
	return true
}

// Get pulls a file from the remote file system and saves it somewhere on our server
func (s *SFTP) Get(destination string, items ...string) error {
	return nil
}
