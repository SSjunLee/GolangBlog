// @Author ljn 2022/5/4 10:32:00
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

type fileUploaderLocal struct {
}

func (u *fileUploaderLocal) Upload(src multipart.File, filename string) (string, error) {
	dir := time.Now().Format("201201/02")[:6]
	_ = os.MkdirAll("./static/upload/"+dir, 0755)
	suffix := path.Ext(filename)
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

var LocalUploader *fileUploaderLocal

func initLocal() {
	LocalUploader = &fileUploaderLocal{}
}
