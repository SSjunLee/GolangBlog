// @Author ljn 2022/4/25 21:07:00
package core

import (
	"Myblog/common/utils"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"
)

type FileUploader interface {
	Upload() (string, error)
}

type LocalFileUploader struct {
	FileHeader *multipart.FileHeader
}

func (u *LocalFileUploader) Upload() (string, error) {
	src, err := u.FileHeader.Open()
	if err != nil {
		return "", errors.New("文件打开失败")
	}
	dir := time.Now().Format("201201/02")[:6]
	_ = os.MkdirAll("./static/upload/"+dir, 0755)
	suffix := path.Ext(u.FileHeader.Filename)
	name := "static/upload/" + dir + "/" + utils.RandomDigitStr(10) + suffix
	log.Println("上传图片 " + name)
	dst, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return "", errors.New("文件创建失败")
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return "/" + name, nil
}

func NewLocalUploader(h *multipart.FileHeader) FileUploader {
	return &LocalFileUploader{
		FileHeader: h,
	}
}
