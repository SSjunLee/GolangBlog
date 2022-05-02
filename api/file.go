// @Author ljn 2022/4/25 17:18:00
package web

import (
	"Myblog/api/response"
	"Myblog/cmd"
	"Myblog/core"
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

func ApiImgUpload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		response.BzError(c, "未发现文件")
		return
	}
	if !checkFileType(f, "image") {
		response.BzError(c, "必须选择图片")
		return
	}
	var uploader core.FileUploader
	if cmd.Config.Image == "local" {
		uploader = core.NewLocalUploader(f)
	} else {
		panic("未指定图片保存位置")
	}
	url, err := uploader.Upload()
	if err != nil {
		response.BzError(c, "上传失败")
		log.Println(err)
		return
	}
	response.Ok(c, url)
}
