// @Author ljn 2022/5/1 11:55:00
package view

import (
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CatePost(c *gin.Context) {
	cateName := c.Param("cate")
	cate := models.CateGet(cateName)
	if cate == nil {
		c.Redirect(302, "/")
		return
	}
	//response.Ok(c, c.PostForm("page"))
	pi, _ := strconv.Atoi(c.Query("page"))
	if pi == 0 {
		pi = 1
	}
	ps := models.MetaInfo.PageSize
	posts := models.PostGetPage(models.KindArticle,
		cate.Id,
		pi,
		ps,
		"id", "title", "path", "created", "summary", "updated", "status")
	if len(posts) == 0 {
		c.Redirect(302, "/")
		return
	}
	total := models.PostCount(models.KindArticle, cate.Id)
	naver := models.Naver{}
	if pi > 1 {
		naver.Prev = `/cate/` + cateName + `?page=1`
	}
	if pi*ps < total {
		naver.Next = `/cate/` + cateName + `?page=` + strconv.Itoa(pi+1)
	}
	render(c, "post-cate.html", gin.H{
		"Cate":      cate,
		"CatePosts": posts,
		"Naver":     naver,
	})

}
