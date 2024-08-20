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
### Trying out our Minio filesystem
- `go mod tidy`
- Video claims to copy test file directly into `./docker-data/minio/testbucket` from IDE, but that did not work
- Had to upload test file through [browser](http://localhost:9001/browser/testbucket)
### Creating a handler to list the remote file system
### Connecting the handler to a route and trying things out
- NOTE: The use of prefix has changed with later versions of Minio. If at root level, set "" as prefix instead of "/"
- NOTE: Files can not be placed directly in docker-data folder, must be uploaded
### Creating handlers to display the upload form
### Creating the handler to process the file upload
### Creating the delete handler

## File systems: sFTP
### Implementing the Put function for sFTP
- SIDENOTE! Updated Makefile in myapp to fix make start and added running tests
- sftp - SFTP support for the go.crypto/ssh package
  [GitHub](https://github.com/pkg/sftp)
  ```shell
  cd celeritas
  go get github.com/pkg/sftp
  cd ..
  ```
### Implementing the List function for sFTP
### Implementing the Delete function for sFTP
### Implementing the Get function for sFTP
- SIDENOTE:
  - Defer statements inside loops *can* cause leaks, specifically in cases where the call's arguments are pointers
    whose pointed-to values are being updated on each iteration. There was no real leak here, but wrapped the defer
    statement in an IIFE (Immediately Invoked Function Expression) nevertheless. This ensures that the object will now
    be closed and there will not be a memory leak.
### Connecting Celeritas to our sFTP file system
### Updating our ListFS handler to support sFTP
- SIDENOTE: Fixed bug in `make restart`
### Updating our PostUploadToFS handler to support sFTP
### Updating our DeleteFromFS handler to support sFTP
### Cleaning up the Get function to avoid resource leaks
- SIDENOTE: Clean up the handling of defer statements in loop

## File systems: WebDAV
### Implementing the Put function for WebDAV
- GoWebDAV - A golang WebDAV client library and command line tool
  [GitHub](https://github.com/studio-b12/gowebdav)
  ```shell
  cd celeritas
  go get github.com/studio-b12/gowebdav
  cd ..
  ```
### Implementing the List function for WebDAV
### Implementing the Delete function for WebDAV
### Implementing the Get function for WebDAV
### Testing things out

## File systems: Amazon S3 Buckets
### Implementing the List function for S3 file systems
- AWS SDK for Go - AWS SDK for the Go programming language.
  [GitHub](https://github.com/aws/aws-sdk-go)
  ```shell
  cd celeritas
  go get github.com/aws/aws-sdk-go/aws
  cd ..
  ```
### Implementing the Put function for S3 file systems
### Implementing the Delete function for S3 file systems
### Implementing the Get function for S3 buckets
### Connecting Celeritas to our S3 file system
### Creating an S3 compatible bucket on Linode
- This can be done on any S3 compatible service
  - On `Linode`, go to `Object Storage` 
  - On `Digital Ocean`, go to `Spaces`
  - On `Amazon`, go to `S3 Buckets`
- Create a bucket
  - Label: `Celeritas`
  - Region: Select a region near you
- Go to access
  - Make ACL `Public Read`
  - Leave CORS `on`
- Create access to you bucket
  - Generate `Access Key`
    - Label: `celeritas`
    - Leave Limited Access `off`
- Copy your Access Key information
  - Access Key: Put this in your `S3_KEY` field in your `.env` file
  - Secret Key Put this in your `SECRET_KEY` field in your `.env` file
- Populate the rest of the `S3` fields in your `.env` file
  - `S3_REGION`: `us-east-1` ## Sample, must match where you selected above
  - `S3_ENDPOINT`: `us-east-1.linodeobjects.com` ## Sample, must match where you selected above
  - `S3_BUCKET`: `celeritas` ## Match what you created above
### Updating our handlers for S3 buckets
### Trying things out
- SIDENOTE: Actual S3 is never tested. Added a check if chosen filesystem is enabled to avoid error

## Building a File System Agnostic File Uploader
### What we'll build
### Adding file systems to the Celeritas type
### Creating the file uploader
- Create files and folders
  ```shell
  ni celeritas/upload.go -type file -Value "package celeritas`n`n"
  ```
### Limiting upload by mime type
- List of commonly used mimetypes: [Common MIME types](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types)
- Mimetype - A fast Golang library for media type and file extension detection, based on magic numbers
  [GitHub](https://github.com/gabriel-vasile/mimetype)
  ```shell
  cd celeritas
  go get github.com/gabriel-vasile/mimetype
  cd ..
  ```
### Adding the mime type and file size limitations to the Celeritas config type
### Setting up handlers and routes to try things out
- SIDENOTE: Used Minio instead of S3
- SIDENOTE2: Also had to run `go mod tidy` on `myapp`
- Create files and folders
  ```shell
  ni myapp/views/celeritas-upload.jet -type file
  ```
### Trying things out

## Improving our Migrations package
### Pop vs. SQL
- [Soda CLI](https://gobuffalo.io/documentation/database/soda/)
- Build and copy cli
  ```shell
  cd celeritas
  make build_cli
  cp dist/celeritas.exe ../myapp/.
  .\celeritas.exe
  #  .\celeritas.exe make migration test
  ```
### Getting started with Pop functions for our migrations code in Celeritas
- Pop - It wraps the absolutely amazing https://github.com/jmoiron/sqlx library, cleans up some of the common
  patterns and workflows usually associated with dealing with databases in Go.
  [GitHub](https://github.com/gobuffalo/pop)
  ```shell
  cd celeritas
  go get github.com/gobuffalo/pop
  cd ..
  ```
### Implementing the CreatePopMigration() function to create up and down migrations
### Implementing the RunPopMigrations() function
### Implementing the PopMigrateDown() function
### Implementing the PopMigrateReset() function
### Making changes in the Celeritas CLI for our pop migrations
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/templates/migrations/migration_up.fizz -type file
  ni celeritas/cmd/cli/templates/migrations/migration_down.fizz -type file
  ```
### Trying out our new make migration command
- Build and copy cli
  ```shell
  cd celeritas
  make build
  cp dist/celeritas.exe ../myapp/.
  ```
- Run cli to test
  ```shell
  cd myapp
  .\celeritas.exe make migration test
  .\celeritas.exe make migration test2 sql
  .\celeritas.exe make migration test3 fizz
  ```
### Ensuring the database is connected before allowing people to make migrations
### Creating a database.yml file and running migrations
### Trying out the migrate command
- First make sure the database is empty
  `drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens; drop table if exists schema_migrations;`
- Build and copy cli
  ```shell
  cd celeritas
  make build
  cp dist/celeritas.exe ../myapp/.
  ```
- Run cli to test
  ```shell
  cd myapp
  .\celeritas.exe make migration test
  .\celeritas.exe migrate
  .\celeritas.exe migrate down
  .\celeritas.exe migrate
  .\celeritas.exe migrate reset
  ```
### Updating the "make auth" command for our Pop integration
### Trying out make auth
- Build and copy cli
  ```shell
  cd celeritas
  make build_cli
  ```
- Run cli to test
  ```shell
  cd myapp
  .\celeritas.exe make auth
  ```

## Social Authentication with OAuth2
### Social Authentication or Single Sign On: an Overview
- [Social login](https://en.wikipedia.org/w/index.php?title=Social_login&oldid=1037502496)
- Goth: Multi-Provider Authentication for Go - Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications.
  [GitHub](https://github.com/markbates/goth)
  ```shell
  cd celeritas
  go get github.com/markbates/goth
  cd ..
  ```
### Getting started with Goth and Social Authentication
### Setting up authentication routes
- NOTE: This includes two commits. The first one with all the files added by the previously run command 
  `celeritas make auth` is not building successfully. This is by purpose to illustrate what was missing in the
  `celeritas make auth` command. This second commit holds the changes needed to fix this. Also adds a `login` route.
### Initializing social sign on
### Implementing the SocialLogin handler
### Implementing the SocialCallback handler
### Connecting our social authentication handlers to routes
### Setting up GitHub for social authentication
- Log in to GitHub
- Go to `Settings` in your profile
- Go to `Developer settings` and then `OAuth Apps`
- Click `Register a new application`
  - Application : `celeritas-oauth`
  - Homepage URL : `http://localhost:4000`
  - Authorization callback URL: `http://localhost:4000/auth/github/callback`
- Store your new client ID into `GITHUB_KEY` in your `.env` file
- Generate and store your new secret into `GITHUB_SECRET` in your `.env` file
- Store your callback URL into `GITHUB_CALLBACK` in your `.env` file
### Trying out the GitHub login functionality
- SIDENOTE: Had to run `go mod tidy`
- `make restart`
### Logging out
### Really logging out
- [Authorizing OAuth apps](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps)
### Trying the socialLogout function
### Adding support for Google login
- Log in to ``Google Cloud Platform``
- Go to `Dashboard`, currently under `Cloud overview` in navigation menu
- Create and select a project
- Find (or search for) the `OAuth consent screen`
- Go to Credentials and create a `OAuth 2.0 Client ID`
  - You need to Configure a Consent Screen first, specify Public and follow the wizard
  - Application Type: `Web application`
  - Name: `Celeritas client 1`
### Updating the auth-handlers.go file for Google to enable login
### Trying out login with Google
- SIDENOTE: Constantly got the error
  ` ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules`
- Had to run the following
  ```shell
  cd myapp
  go get cloud.google.com/go
  go mod tidy
  cd ..
  ```
### Adding the case for logging out of Google in socialLogout()
### Trying things out

## RPC, Graceful Shutdown, and additional changes
### Separating Web and API routes
- SIDENOTE: This was already implemented...
### Getting started with "Maintenance Mode" functionality using RPC
### Starting RPC
### Adding maintenance mode middleware
- SIDENOTE: This file already existed...
- Create files and folders
  ```shell
  ni myapp/public/maintenance.html -type file
  ```
### Updating the CLI for maintenance mode
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/rpc.go -type file -Value "package main`n`n"
  ```
### Testing the maintenance mode functionality
- Build and copy cli
  ```shell
  cd celeritas
  make build_cli
  ```
- Run cli to test
  ```shell
  cd myapp
  .\celeritas.exe down
  .\celeritas.exe up
  ```
- Note: Functionality for `ALLOWED_URLS` in `.env` file is not implemented, but could be
### Graceful Shutdown
- Create files and folders
  ```shell
  ni celeritas/server.go -type file -Value "package celeritas`n`n"
  ```
- NOTE: Functions in myapp/main.go already existed

## Testing
### Adding a simple setup_test.go file to handlers
- Create files and folders
  ```shell
  ni myapp/handlers/setup_test.go -type file -Value "package handlers`n`n"
  ```
- NOTE: The Test Setup File in myapp/handlers already existed with the package variables and the TestMain function
### Adding two functions to our setup_test.go file
- NOTE: The Test Setup File in myapp/handlers already existed with a partly getRoutes and a full getCtx functions
### Adding and running a sample test
- Create files and folders
  ```shell
  ni myapp/handlers/handlers_test.go -type file -Value "package celeritas`n`n"
  ```
- NOTE: `myapp/handlers/handlers_test.go` already existed
- Run tests in `myapp` folder: `go test ./handlers/...`
### Adding some additional tests
### Implementing Laravel Dusk like screen captures
- [Laravel Dusk documentation](https://laravel.com/docs/8.x/dusk)
- [Rod documentation](https://go-rod.github.io/#/)
  - Rod: A Chrome DevTools Protocol driver for web automation and scraping.
    [GitHub](https://github.com/go-rod/rod)
    ```shell
    cd celeritas
    go get github.com/go-rod/rod
    cd ..
    ```
### Writing the screen capture function
- Create files and folders
  ```shell
  ni celeritas/dusk.go -type file -Value "package celeritas`n`n"
  md myapp/screenshots
  ```
### Trying out the screen capture function
- NOTE: This was already added in `myapp/handlers/handlers_test.go`
- SIDENOTE: Had to run `go mod tidy`
- SIDENOTE2: Could not test this as it was flagged by Windows Security as suspicious. So commented out... 
  `Operation did not complete successfully because the file contains a virus or potentially unwanted software.`
  Seems to be connected to something down the chain calling `Leakless(true)` https://github.com/ysmood/leakless
### Writing additional helper functions for testing
- Create files and folders
  ```shell
  ni myapp/views/tester.jet -type file
  ```
- SIDENOTE: Could not test this as it was flagged by Windows Security as suspicious. So commented out...
  `Operation did not complete successfully because the file contains a virus or potentially unwanted software.`
  - Solution 1: Disable `leakless` by creating your own *Launcher and set its `leakless` property to false.
    [](https://pkg.go.dev/github.com/go-rod/rod@v0.79.0/lib/launcher#Launcher.Leakless)
  - Solution 2: Tell your antivirus to ignore the `leakless` binary.
- TODO: Investigate above options

## Final changes to the Celeritas CLI application
### Updating our templates in the CLI, and making some changes to the myapp
### Creating our skeleton app
### Additional updates to the skeleton application and the celeritas project
###  Trying out the "celeritas new <project>" command
- Prepare celeritas-app
  - search and replace from `github.com/tsawler/celeritas` to `github.com/john-wraa/celeritas`
  - run `go get github.com/john-wraa/celeritas`
  - run `go mod tidy`
  - Push changes to GitHub and optionally create a new tag
- Build cli in celeritas project
  ```shell
  cd celeritas
  make build
  cp dist/celeritas.exe ../testapp
  cd ../testapp
  ./celeritas.exe new testapp4
  ```
- Open new app
- run `go mod tidy`
- run `make start`
- SIDENOTE: Also had to close all goland processes but one and run `go clean cache` and `go clean -modcache`






## Repo creation Log
- git init
- git add .
- git commit -m "Initial entry"
- git remote add origin https://github.com/johnwr-response/golang-lets-build-a-go-version-of-laravel-part2.git
- git branch -M main
- git push -u origin main
- git branch -M 01-walkthrough
