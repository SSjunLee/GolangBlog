// @Author ljn 2022/4/25 10:28:00
package web

import (
	"Myblog/api/response"
	"Myblog/common"
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ApiGetPageInPage(c *gin.Context) {
	in := models.BuildPageFromHttpParams(c)
	cnt := models.PostCount(models.KindPage, -1)
	if cnt < 1 {
		response.BzError(c, "未查询到结果")
		return
	}
	args := []string{"id", "title", "summary", "updated", "status", "path"}
	posts := models.PostGetPage(models.KindPage, -1, in.Pid, in.Psize, args...)
	if posts == nil {
		response.BzError(c, "未查询到结果")
		return
	}
	response.Page(c, posts, cnt)
}

func ApiPageDrop(c *gin.Context) {
	in := struct {
		Id int `json:"id"`
	}{}
	_ = c.BindJSON(&in)
	err := models.PostDrop(in.Id)
	if err != nil {
		panic(err)
	}
	response.Ok(c, "删除成功")
}

func ApiPageAdd(c *gin.Context) {
	p := models.Post{}
	err := c.BindJSON(&p)
	if err != nil {
		panic(err)
	}
	p.Created.Format(common.StdDateTime)
	p.Updated = p.Created
	p.Kind = models.KindPage
	err = models.PostAdd(&p)
	if err != nil {
		panic(err)
	}
	response.Ok(c, "添加成功")
}

func ApiPageGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.Error(c, 404, "文章不存在")
		return
	}
	p := models.PostGet(id)
	if p == nil {
		response.Error(c, 404, "文章不存在")
		return
	}
	response.Ok(c, p)
}

func ApiPageEdit(c *gin.Context) {
	p := models.Post{}
	err := c.BindJSON(&p)
	//log.Printf("%+v",p)
	if err != nil {
		panic(err)
	}
	p.HandleRichText()
	p.Updated.Format(common.StdDateTime)
	err = models.PostEdit(&p)
	if err != nil {
		panic(err)
	}
	response.Ok(c, "编辑成功")
}
