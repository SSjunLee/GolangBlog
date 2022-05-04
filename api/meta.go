// @Author ljn 2022/5/2 8:47:00
package api

import (
	"Myblog/api/response"
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"log"
)

func MetaGet(c *gin.Context) {
	meta := models.MetaGet(1)
	if meta == nil {
		response.NotFound(c)
		return
	}
	response.Ok(c, meta)
}

func MetaEdit(c *gin.Context) {
	meta := models.Meta{}
	err := c.BindJSON(&meta)
	log.Printf("%+v", meta)
	if err != nil {
		log.Println(err)
		response.Error(c, 500, "输入错误")
		return
	}
	meta.Id = 1
	_ = models.MetaEdit(&meta)
	models.MetaInfo.Load()
	response.Ok(c, "ok")
}
