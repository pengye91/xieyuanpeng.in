package aws

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

func UploadToS3(files []*multipart.FileHeader, contentTypes []string, sizes []int64) error {
	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()
	if err != nil {
		log.LoggerSugar.Errorw("UploadToS3 Error: creds.Get() error",
			"module", "AWS S3",
			"bad credentials", err,
		)
		return err
	}
	cfg := aws.NewConfig().WithRegion(configs.AWS_REGION).WithCredentials(creds)
	s3Session, s3SessionErr := session.NewSession()

	if s3SessionErr != nil {
		log.LoggerSugar.Errorw("UploadToS3 Error: session.NewSession error",
			"module", "AWS S3",
			"bad credentials", s3SessionErr,
		)
		return s3SessionErr
	}
	svc := s3.New(s3Session, cfg)

	for index, item := range files {
		file, err := item.Open()
		if err != nil {
			log.LoggerSugar.Errorw("UploadToS3 Error: file.Open error",
				"module", "application []*multipart.FileHeader",
				"bad credentials", err,
			)
			return err
		}

		path := "/public/images/" + item.Filename
		params := &s3.PutObjectInput{
			Bucket:        aws.String(configs.AWS_S3_BUCKET),
			Key:           aws.String(path),
			Body:          file,
			ContentType:   aws.String(contentTypes[index]),
			ContentLength: aws.Int64(sizes[index]),
		}

		if resp, err := svc.PutObject(params); err != nil {
			log.LoggerSugar.Errorw("UploadToS3 Error: PutObject error",
				"module", "AWS S3 service.PutObject",
				"bad response", err,
			)
			return err
		} else {
			fmt.Printf("response %s", awsutil.StringValue(resp))
		}
	}
	return nil
}
