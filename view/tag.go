// @Author ljn 2022/5/1 14:07:00
package view

import (
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func TagsView(c *gin.Context) {
	tags := models.GetAllTagState()
	render(c, "tags.html", gin.H{
		"Tags": tags,
	})
}

func TagPost(c *gin.Context) {
	tagName := c.Param("tag")
	if tagName == "" {
		c.Redirect(302, "/")
		return
	}

	tag := models.TagGet(tagName)
	if tag == nil {
		c.Redirect(302, "/")
		return
	}
	pi, _ := strconv.Atoi(c.Query("page"))
	if pi == 0 {
		pi = 1
	}
	posts := models.TagPostPage(tag.Id,
		pi,
		models.MetaInfo.PageSize,
		"id", "title", "summary", "path", "created", "updated")

	if posts == nil {
		c.Redirect(302, "/")
		return
	}

	sum := models.TagPostCount(tag.Id)
	naver := models.Naver{}
	if pi > 1 {
		naver.Prev = "/tag/" + tagName + "?page=1"
	}
	if pi*models.MetaInfo.PageSize < sum {
		naver.Next = "/tag/" + tagName + "?page=" + strconv.Itoa(pi+1)
	}
	/*
		response.Ok(c, gin.H{
			"Tag":      tag,
			"TagPosts": posts,
			"Naver":    naver,
		})*/

	render(c, "post-tag.html", gin.H{
		"Tag":      tag,
		"TagPosts": posts,
		"Naver":    naver,
	})
}
