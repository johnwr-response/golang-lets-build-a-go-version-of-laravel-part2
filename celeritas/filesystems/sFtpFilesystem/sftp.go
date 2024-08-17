package sFtpFilesystem

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/tsawler/celeritas/filesystems"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type SFTP struct {
	Host string
	User string
	Pass string
	Port string
}

// getCredentials generates sftp client using the credentials stored in the SFTP type
func (s *SFTP) getCredentials() (*sftp.Client, error) {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User:            s.User,
		Auth:            []ssh.AuthMethod{ssh.Password(s.Pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	cwd, err := client.Getwd()
	if err != nil {
		return nil, err
	}
	log.Println("Current working directory: ", cwd)

	return client, nil
}

// Put transfers a file to the remote file system
func (s *SFTP) Put(fileName, folder string) error {
	client, err := s.getCredentials()
	if err != nil {
		return err
	}
	defer func(client *sftp.Client) {
		_ = client.Close()
	}(client)

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	f2, err := client.Create(path.Base(fileName))
	if err != nil {
		return err
	}
	defer func(f2 *sftp.File) {
		_ = f2.Close()
	}(f2)

	if _, err := io.Copy(f2, f); err != nil {
		return err
	}

	return nil
}

// List returns a listing of all files in the remote bucket with the given prefix, except for files named with a leading .
func (s *SFTP) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	client, err := s.getCredentials()
	if err != nil {
		return listing, err
	}
	defer func(client *sftp.Client) {
		_ = client.Close()
	}(client)

	files, err := client.ReadDir(prefix)
	if err != nil {
		return listing, err
	}
	for _, f := range files {
		var item filesystems.Listing

		if !strings.HasPrefix(f.Name(), ".") {
			b := float64(f.Size())
			kb := b / 1024
			mb := kb / 1024
			item.Key = f.Name()
			item.Size = mb
			item.LastModified = f.ModTime()
			item.IsDir = f.IsDir()
			listing = append(listing, item)
		}
	}

	return listing, nil
}

// Delete removes one or more files from the remote filesystem
func (s *SFTP) Delete(itemsToDelete []string) bool {
	client, err := s.getCredentials()
	if err != nil {
		return false
	}
	defer func(client *sftp.Client) {
		_ = client.Close()
	}(client)

	for _, item := range itemsToDelete {
		deleteErr := client.Remove(item)
		if deleteErr != nil {
			return false
		}
	}

	return true
}

// Get pulls a file from the remote file system and saves it somewhere on our server
func (s *SFTP) Get(destination string, items ...string) error {
	return nil
}
