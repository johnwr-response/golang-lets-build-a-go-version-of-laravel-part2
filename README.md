# golang-lets-build-a-go-version-of-laravel-part2
The followup to "Let's Build a Go Version of Laravel," with support for remote file systems, Social Auth, and more.

## What you'll learn
- How to implement and use Remote Procedure Calls (RPC) in Go
- How to upload files safely in Go
- How to integrate AWS S3 Buckets in a Go application
- How to integrate an FTP/SFTP filesystem in Go
- How to implement social authentication in Go

## Requirements
- A basic understanding of the Go programming language
- A basic understanding of HTML
- A basic understanding of JavaScript

## Course content
- Introduction
- Project Setup
- Setting up our remote file systems
- File systems: Minio
- File systems: sFTP
- File systems: WebDAV
- File systems: Amazon S3 Buckets
- Building a File System Agnostic File Uploader
- Improving our Migrations package
- Social Authentication with OAuth2
- RPC, Graceful Shutdown, and additional changes
- Testing
- Final changes to the Celeritas CLI application

## Description
This is the follow-up to part one of this course and is intended for students who have already taken that course!  
[Let's Build a Go Version of Laravel](https://github.com/johnwr-response/golang-lets-build-a-go-version-of-laravel.git) 

In the first part of this series, we built a re-usable Go module that gave us a lot of functionality, including:
- html, json, and xml response types
- support for Go templates and Jet templates to render pages
- multiple database support
- sessions
- and more.
This time around, we'll improve our Celeritas package and add the following functionality:
- Add support for remote file systems, including Amazon S3 buckets, Minio, sFTP, and WebDAV 
- Add support for Social Authentication using GitHub and Google (and you can add as many more as you like)
- Add support for improved testing, including a Go version of Laravel Dusk package, which takes a browser screenshot
  when testing functionality that renders a web page
- Add support for "maintenance mode" using Remote Procedure Calls (RPC)
- Improve our database migrations to support both raw SQL and soda's Fizz file format 
- Implement file upload functionality (with support for local and remote file systems)
- Separate logic and routes for web and API 
- Make it easy for users to create tests by pre-populating stub test files and the appropriate `setup_test.go` files

By the time that you have completed this course, you will not only have a solid understanding of each of
the things listed above, but also a reusable code base that will help you jump start your next project.

## Who this course is for:
- This course is intended for developers who wish to further their knowledge of using Go to build web applications
- It's also great for PHP & Laravel developers who want to build faster, safer web applications using Go

## Project apps: celeritas and myapp
- Built in Go version 1.23.0
  - Uses the minio.go package for providing S3 compatible object storage
    [link](https://github.com/minio/minio-go)
    - Transiently added packages in `go.mod`
      - `github.com/go-ini/ini`
      - `github.com/goccy/go-json`
      - `github.com/google/uuid`
      - `github.com/klauspost/cpuid/v2`
      - `github.com/minio/md5-simd`
      - `github.com/rs/xid`
  - Uses the sftp package adding SFTP support for the go.crypto/ssh package
    [link](https://github.com/pkg/sftp)
    - Transiently added packages in `go.mod`
      - `github.com/kr/fs`
  - Uses the GoWebDAV package adding Webdav support
    [link](https://github.com/studio-b12/gowebdav)
  - Uses the AWS SDK for Go to support Amazon S3 Buckets
    [link](https://github.com/aws/aws-sdk-go)
  - Uses the Mimetype package for Go to support mimetype detection
    [link](https://github.com/gabriel-vasile/mimetype)
