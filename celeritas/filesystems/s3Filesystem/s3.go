package s3Filesystem

import "github.com/tsawler/celeritas/filesystems"

type S3 struct {
	Key      string
	Secret   string
	Region   string
	Endpoint string
	Bucket   string
}

// Put transfers a file to the remote file system
func (s *S3) Put(fileName, folder string) error {
	return nil
}

// List returns a listing of all files in the remote bucket with the given prefix, except for files named with a leading .
func (s *S3) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	return listing, nil
}

// Delete removes one or more files from the remote filesystem
func (s *S3) Delete(itemsToDelete []string) bool {
	return true
}

// Get pulls a file from the remote file system and saves it somewhere on our server
func (s *S3) Get(destination string, items ...string) error {
	return nil
}
