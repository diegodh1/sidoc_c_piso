package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"time"
	"mime/multipart"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"strconv"
   )

//Config struct
type ConfigAWS struct {
	AccessKeyID string
	SecretAccessKey string
	MyRegion string
}

//Initialize func
func (c *ConfigAWS) Initialize() {
	c.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	c.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	c.MyRegion = "us-east-1"
}   

func ConnectAws() (*session.Session, error) {
	conf := ConfigAWS{}
	conf.Initialize()
	sess, err := session.NewSession(
		&aws.Config{
		 Region: aws.String(conf.MyRegion),
		 Credentials: credentials.NewStaticCredentials(
			conf.AccessKeyID,
			conf.SecretAccessKey,
		  "", // a token will be created when the session it's used.
		 ),
		})
	if err != nil {
		return nil, err
	}
	
	return sess, nil

}

func UploadToAWS(img *multipart.FileHeader) (string, error) {

	sess, er := ConnectAws()
	
	if er != nil {
		return "", er
	}
	uploader := s3manager.NewUploader(sess)
	//fileName := generateFileNameForBucket(img.Filename)
	body, _ := img.Open()
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key: aws.String(img.Filename),
		Body: body,
	})
	if err != nil {
		return "", err
	}

	return img.Filename, nil

}

func generateFileNameForBucket(name string) string {
	curTime := time.Now().UnixNano()
	var strArr []string
	strArr = append(strArr, name)
	strArr = append(strArr, strconv.Itoa(int(curTime)))
	var filename string
	for _, str := range strArr {
		filename = filename + str
	}
	return filename
}   