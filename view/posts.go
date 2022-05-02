// @Author ljn 2022/5/1 11:44:00
package view

import (
	"Myblog/api/response"
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Post(c *gin.Context) {
	paths := strings.Split(c.Param("path"), ".")
	log.Println(paths)
	if len(paths) != 2 {
		c.Redirect(302, "/404")
		return
	}
	if paths[1] != "html" {
		response.Ok(c, "ok")
		return
	}
	paths[0] = paths[0][1:]
	p := models.GetPostByPath(paths[0])
	//log.Println(p)
	if p == nil {
		c.Redirect(302, "/404")
		return
	}
	naver := p.GetNav()
	cate := models.CateGet(p.CatId)
	wrapper := struct {
		*models.Post
		Cate *models.Cate `json:"cate"`
	}{p, cate}

	//log.Printf("%+v",wrapper.Post)

	render(c, "post.html", gin.H{
		"Post":  wrapper,
		"Naver": naver,
		"Show":  p.Status == 2,
	})

}
