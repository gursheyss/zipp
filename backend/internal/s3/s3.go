package s3

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func ConnectToAWS(ctx context.Context) (*s3.Client, error) {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	log.Println("Successfully connected to AWS")

	return client, nil
}

func UploadFile(ctx context.Context, client *s3.Client, file io.Reader, bucket, key string, id string) (string, error) {
	fullKey := id + "/" + key
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(fullKey),
		Body:                 file,
		ServerSideEncryption: types.ServerSideEncryptionAes256,
	})

	if err != nil {
		return "", err
	}

	log.Println("File uploaded successfully")
	return fullKey, nil
}

func DownloadAllFiles(ctx context.Context, client *s3.Client, bucket, id string) (map[string][]byte, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(id + "/"),
	}

	resp, err := client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, err
	}

	files := make(map[string][]byte)
	for _, item := range resp.Contents {
		getObjInput := &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    item.Key,
		}

		getObjResp, err := client.GetObject(ctx, getObjInput)
		if err != nil {
			return nil, err
		}

		defer getObjResp.Body.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, getObjResp.Body)
		if err != nil {
			return nil, err
		}

		files[*item.Key] = buf.Bytes()
	}

	log.Println("All files downloaded successfully")

	return files, nil
}
