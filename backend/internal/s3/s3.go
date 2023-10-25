package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func UploadFile(ctx context.Context, client *s3.Client, file io.Reader, bucket, key string, id string, contentType string) (string, error) {
	fullKey := id + "/" + key
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(fullKey),
		Body:                 file,
		ServerSideEncryption: types.ServerSideEncryptionAes256,
		ContentType:          aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	log.Println("File uploaded successfully")
	return fullKey, nil
}

func DownloadFile(ctx context.Context, client *s3.Client, bucket, id string) ([]byte, string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(id + "/"),
	}

	resp, err := client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, "", err
	}

	if len(resp.Contents) == 0 {
		return nil, "", fmt.Errorf("no files found for id %s", id)
	}

	// Get the first file
	item := resp.Contents[0]
	getObjInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    item.Key,
	}

	getObjResp, err := client.GetObject(ctx, getObjInput)
	if err != nil {
		return nil, "", err
	}

	defer getObjResp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, getObjResp.Body)
	if err != nil {
		return nil, "", err
	}

	log.Println("File downloaded successfully")

	filename := strings.Split(*item.Key, "/")[1]
	log.Println("Filename of download " + filename)

	return buf.Bytes(), filename, nil
}
