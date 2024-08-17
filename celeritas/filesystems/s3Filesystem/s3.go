package s3Filesystem

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tsawler/celeritas/filesystems"
)

type S3 struct {
	Key      string
	Secret   string
	Region   string
	Endpoint string
	Bucket   string
}

// getCredentials generates s3 client using the credentials stored in the SFTP type
func (s *S3) getCredentials() *credentials.Credentials {
	cred := credentials.NewStaticCredentials(s.Key, s.Secret, "")
	return cred
}

// Put transfers a file to the remote file system
func (s *S3) Put(fileName, folder string) error {
	return nil
}

// List returns a listing of all files in the remote bucket with the given prefix, except for files named with a leading .
func (s *S3) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	cred := s.getCredentials()
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    &s.Endpoint,
		Region:      &s.Region,
		Credentials: cred,
	}))

	svc := s3.New(sess)
	input := &s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(prefix),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		var aErr awserr.Error
		if errors.As(err, &aErr) {
			switch aErr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aErr.Error())
			default:
				fmt.Println(aErr.Error())
			}
		}
		return nil, err
	}

	for _, key := range result.Contents {
		b := float64(*key.Size)
		kb := b / 1024
		mb := kb / 1024
		current := filesystems.Listing{
			Etag:         *key.ETag,
			LastModified: *key.LastModified,
			Key:          *key.Key,
			Size:         mb,
		}
		listing = append(listing, current)
	}

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
