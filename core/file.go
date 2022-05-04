// @Author ljn 2022/4/25 21:07:00
package core

import "mime/multipart"

type FileUploader interface {
	Upload(src multipart.File, filename string) (string, error)
}

func InitFileUploader() {
	initOss()
	initLocal()
}
