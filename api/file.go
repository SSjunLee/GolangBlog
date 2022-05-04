// @Author ljn 2022/4/25 17:18:00
package api

import (
	"Myblog/api/response"
	"Myblog/core"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"strings"
)

func checkFileType(f *multipart.FileHeader, fileType string) bool {
	if strings.Contains(f.Header.Get("Content-Type"), fileType) {
		return true
	}
	return false
}

func getFile(c *gin.Context) (multipart.File, string, error) {
	f, err := c.FormFile("file")
	if err != nil {
		return nil, "", errors.New("未发现文件")
	}
	if !checkFileType(f, "image") {
		return nil, "", errors.New("必须选择图片")
	}
	src, err := f.Open()
	if err != nil {
		return nil, "", errors.New("文件打开失败")
	}
	return src, f.Filename, nil
}

func upload(uploader core.FileUploader, src multipart.File, filename string, c *gin.Context) {
	url, err := uploader.Upload(src, filename)
	if err != nil {
		log.Println(err)
		response.BzError(c, "上传失败")
		return
	}
	response.Ok(c, url)
}

func ImageUploadOss(c *gin.Context) {
	src, filename, err := getFile(c)
	if err != nil {
		response.BzError(c, err.Error())
		return
	}
	if core.OssUploader == nil {
		response.BzError(c, "请配置oss")
		return
	}
	upload(core.OssUploader, src, filename, c)
}

func ImgUploadLocal(c *gin.Context) {
	src, filename, err := getFile(c)
	if err != nil {
		response.BzError(c, err.Error())
		return
	}
	upload(core.LocalUploader, src, filename, c)
}
