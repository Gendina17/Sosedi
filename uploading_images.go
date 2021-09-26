package main

import (
	"net/http"
    "io"

	"mime/multipart"
	"bytes"
	"path/filepath"

	"github.com/globalsign/mgo/bson"
	"github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
  BUCKET_NAME = "sosedi-ng-rb"
  REGION = "eu-west-2"
  AccessKeyID = "AKIASJYGCYOCO7GYQ4X5"
  AccessKey = "Xs6mt2DPBaJFKmf4tpVz1DGJpBTRN8aV2EoI6a5G"
  UserKey = "AKIASJYGCYOCKR2SATBN"
  SecretUserKey = "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"   // потом спрячу не смарите плз))
)

func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	tempFileName := bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(BUCKET_NAME),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

func getObject(filename string) string {

  s, _ := session.NewSession(&aws.Config{
		Region: aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, AccessKey, ""),
	})

  resp, _ := s3.New(s).GetObject(&s3.GetObjectInput{
    Bucket: aws.String(BUCKET_NAME),
    Key: aws.String(filename),
  })

  size := int(*resp.ContentLength)

	buffer := make([]byte, size)
	defer resp.Body.Close()
	var bbuffer bytes.Buffer
	for true {
		num, rerr := resp.Body.Read(buffer)
		if num > 0 {
			bbuffer.Write(buffer[:num])
		} else if rerr == io.EOF || rerr != nil {
			break
		}
	}
	return bbuffer.String()
}
//во время редактирования мб пригодится
// func deleteObject(filename string) (resp *s3.DeleteObjectOutput) {
//   fmt.Println("Deleting: ", filename)
//   resp, err := s3session.DeleteObject(&s3.DeleteObjectInput{
//     Bucket: aws.String(BUCKET_NAME),
//     Key: aws.String(filename),
//   })
//
//   if err != nil {
//     panic(err)
//   }
//
//   return resp
// }

func getting_image_from_request(r *http.Request) (string, string) {
  maxSize := int64(1024000)

  err := r.ParseMultipartForm(maxSize)
  if err != nil {
    return "", "Изображение слишком большое"
  }

  file, fileHeader, err := r.FormFile("photo")

	if err != nil {
    return "", "Загрузите изображение"
	}

	defer file.Close()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, AccessKey, ""),
	})

  if err != nil {
    return "", "Ошибка загрузки изображения"
	}

  photo_key, err := UploadFileToS3(s, file, fileHeader)

	if err != nil {
    return "", "К сожалению, нам не удалось загрузить изображение"
	}

  return photo_key, "ok"
}
