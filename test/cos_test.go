package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse("https://getcharzp-1256268070.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/exampleobject.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/1ff6a037-409d-445a-86cc-6dbca2b29c87.jpeg", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://getcharzp-1256268070.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/exampleobject2.jpg"

	f, err := os.ReadFile("./img/1ff6a037-409d-445a-86cc-6dbca2b29c87.jpeg")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}
