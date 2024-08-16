# Let's Build a Go version of Laravel - Part 2

## Introduction
### Introduction
- Let's build a Go version of Laravel: Part II
  - This course should *only* be taken *after* the first course
  - We'll build on core functionality
- What we'll cover
  - Remote file systems: Minio, sFTP, WebDAV, Amazon S3 buckets
  - A file system agnostic file uploader
  - Improving migrations to support Fizz
  - Social authentication: GitHun and Google
  - RPC calls
  - Simplify testing
  - Screen capture in testing
  - Browser emulation in testing
### About me
### Asking for help
### Installing Go
- Download Go [here](https://go.dev/dl/)
- WinGet
  - Install: ```winget install --id GoLang.Go```
  - Upgrade: ```winget upgrade --id GoLang.Go```
- Verify  
  ```go version```
### Installing an IDE
- Visual Studio Code
  - Install: ```winget install --id Microsoft.VisualStudioCode```
  - Upgrade: ```winget upgrade --id Microsoft.VisualStudioCode```
  - Add extensions:
    - [Go]
      - Also, press `shift+ctl` and search for `Go: Install/Update Tools`
        - Click on it, select all associated checkboxes and click OK to install them
    - [goTemplate-syntax]
- GoLand
  - Install using toolbox: ```winget install --id JetBrains.Toolbox```
  - Install directly: ```winget install --id JetBrains.GoLand```

## Project Setup
### Setting up our project
- Run the necessary setup commands
  ```shell
  cd celeritas
  make build
  cd dist
  ./celeritas.exe new myapp
  mv myapp ../..
  ```
- Either setup [go-workspace](go.work) manually or tell the ide to do it for you
### Making sure everything works
- Fixed Makefile, replaced leading spaces with tab
- Changed output folder to `dist` when building
- Fixed `c.App.ListenAndServe() used as value` by removing the error check as ListenAndServe does not return one
- Removed reference to `github.com/tsawler/celeritas` in top `require` section of `go.mod` 
- `go mod tidy`
- `make start`
- Added `HOST_INTERFACE=localhost` in `.env`
- Used `HOST_INTERFACE` when starting server to avoid Firewall prompt in Windows
- Create files and folders
  ```shell
  md celeritas/testFolder
  ni celeritas/testFolder/test.go -type file -Value "package testFolder`n`n"
  ```
- Check new test route [link](http://localhost:4000/test-route)

## Setting up our remote file systems
### What we're going to create
### Setting up our remote file systems using Docker
- Create files and folders
  ```shell
  md myapp/docker-data/postgres
  md myapp/docker-data/redis
  md myapp/docker-data/mariadb
  md myapp/docker-data/sftp
  md myapp/docker-data/home
  md myapp/docker-data/minio
  ```
- Run Docker containers
  ```shell
  docker compose -f ./myapp/docker-compose.yml up  -d
  ```
### Configuring Minio
- Minio - Multi-Cloud Object Storage
  [Link](https://min.io/)
- Go to local instance [Link](http://localhost:9001) and log in
- Create a Bucket called `testbucket`
### Configuring sFTP
- SFTPGo -  SFTP/FTP/WebDAV server. S3, GCS, Azure and local fs as storage backends
  [Link](https://hub.docker.com/r/drakkan/sftpgo)
- Go to local instance [Link](http://localhost:8080) and create an `admin` user
  - Username: admin
  - Password: password
- Create a new user by clicking `Users` and `Add`
  - Username: sftp
  - Password: password
  - Set to `Root directory` to `/mnt/data`
  - Make sure an Asterisk (*) is selected under `ACLs` and `Permissions`
  - Click `Save`
### Setting up a type for file systems
- Create files and folders
  ```shell
  md celeritas/filesystems
  ni celeritas/filesystems/filesystems.go -type file -Value "package filesystems`n`n"
  ```
## File systems: Minio
### Getting started with Minio: connecting and the Put function
- [Minio documentation](https://min.io/docs/minio/kubernetes/upstream/index.html?ref=docs-redirect)
- minio-go - MinIO Go client SDK for S3 compatible object storage
  [GitHub](https://github.com/minio/minio-go)
  ```shell
  cd celeritas
  go get github.com/minio/minio-go/v7
  go get github.com/minio/minio-go/v7/pkg/credentials
  cd ..
  ```
- Create files and folders
  ```shell
  md celeritas/filesystems/minioFilesystem
  ni celeritas/filesystems/minioFilesystem/minio.go -type file -Value "package minioFilesystem`n`n"
  ```
### Implementing the List function in Minio
### Implementing the Delete function in Minio
### Implementing the Get function in Minio
### Creating stub filesystems for the other three types
- Create files and folders
  ```shell
  md celeritas/filesystems/s3Filesystem
  md celeritas/filesystems/sFtpFilesystem
  md celeritas/filesystems/webdavFilesystem
  ni celeritas/filesystems/s3Filesystem/s3.go -type file -Value "package s3Filesystem`n`n"
  ni celeritas/filesystems/sFtpFilesystem/sftp.go -type file -Value "package sFtpFilesystem`n`n"
  ni celeritas/filesystems/webdavFilesystem/webdav.go -type file -Value "package webdavFilesystem`n`n"
  ```
### Adding filesystems to Celeritas









## File systems: sFTP
## File systems: WebDAV
## File systems: Amazon S3 Buckets
## Building a File System Agnostic File Uploader
## Improving our Migrations package
## Social Authentication with OAuth2
## RPC, Graceful Shutdown, and additional changes
## Testing
## Final changes to the Celeritas CLI application

## Repo creation Log
- git init
- git add .
- git commit -m "Initial entry"
- git remote add origin https://github.com/johnwr-response/golang-lets-build-a-go-version-of-laravel-part2.git
- git branch -M main
- git push -u origin main
- git branch -M 01-walkthrough
