package web

import (
	"Myblog/api/response"
	"Myblog/common"
	"github.com/gin-gonic/gin"
	"log"
)

func Test(c *gin.Context) {
	log.Println(c.Request)
	str := c.Query("arg")
	log.Println(str)
	log.Println(c.Request.Header.Get(common.JWTHeader))
	response.Ok(c, "ok")
}
