// @Author ljn 2022/5/4 10:49:00
package core

import (
	"Myblog/cmd"
	"Myblog/common/utils"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type fileUploaderOss struct {
	BucketUrl string
	SecretID  string
	SecretKey string
	Url       url.URL
	Client    *cos.Client
}

const (
	imageBase = "image"
)

var OssUploader *fileUploaderOss

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		log.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		log.Printf("ERROR: Code: %v\n", e.Code)
		log.Printf("ERROR: Message: %v\n", e.Message)
		log.Printf("ERROR: Resource: %v\n", e.Resource)
		log.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		log.Printf("ERROR: %v\n", err)
		// ERROR
	}
}

func (f *fileUploaderOss) Upload(src multipart.File, filename string) (string, error) {
	suffix := filename[strings.Index(filename, "."):]
	dir := time.Now().Format("201201/02")[:6]
	newFilename := fmt.Sprintf(`%s/%s/%s%s`, imageBase, dir, utils.RandomDigitStr(10), suffix)
	log.Println(newFilename)
	if _, err := f.Client.Object.Put(context.Background(), newFilename, src, nil); err != nil {
		log_status(err)
		return "", err
	}
	return f.BucketUrl + "/" + newFilename, nil

}

func initOss() {

	u, err := url.Parse(cmd.Config.Oss.BucketUrl)
	if err != nil {
		log.Println(err)
		return
	}
	OssUploader = &fileUploaderOss{
		BucketUrl: cmd.Config.Oss.BucketUrl,
		SecretID:  cmd.Config.Oss.SecretID,
		SecretKey: cmd.Config.Oss.SecretKey,
	}
	b := &cos.BaseURL{BucketURL: u}
	OssUploader.Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 COS_SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: OssUploader.SecretID,
			// 环境变量 COS_SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: OssUploader.SecretKey,
			// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

}
