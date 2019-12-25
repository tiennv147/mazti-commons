package s3

import (
	"errors"
	"github.com/tiennv147/mazti-commons/config"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	logger = log.New(os.Stdout, "s3", log.LstdFlags|log.LUTC|log.Llongfile)
)

type Service struct {
	region             string
	bucket             string
	profileCredentials *credentials.Credentials
}

func NewService(conf *config.S3) (*Service, error) {
	if conf.Region == "" {
		return nil, errors.New("missing region configuration")
	}
	if conf.Bucket == "" {
		return nil, errors.New("missing bucket configuration")
	}
	return &Service{
		region:             conf.Region,
		bucket:             conf.Bucket,
		profileCredentials: credentials.NewSharedCredentials(conf.ProfileFilePath, conf.Profile),
	}, nil
}

func (s Service) Upload(fileName string, file io.ReadSeeker) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.region),
	})

	output, err := s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
		ACL:    aws.String("private"),
		Body:   file,
	})

	if output != nil {
		logger.Print("Etag", output.ETag)
	}

	return err
}
