package utils

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
)

func UploadToS3(files []*multipart.FileHeader, contentTypes []string, sizes []int64) {
	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion(configs.AWS_REGION).WithCredentials(creds)
	s3Session, s3SessionErr := session.NewSession()

	if s3SessionErr != nil {
		fmt.Println(s3SessionErr)
		return
	}
	svc := s3.New(s3Session, cfg)

	for index, item := range files {
		file, err := item.Open()
		if err != nil {
			fmt.Println(err)
			return
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
			fmt.Printf("bad response: %s", err)
			return
		} else {
			fmt.Printf("response %s", awsutil.StringValue(resp))
		}
	}
}
