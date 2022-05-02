// @Author ljn 2022/4/26 15:32:00
package web

import (
	"Myblog/api/response"
	"Myblog/models"
	"github.com/gin-gonic/gin"
)

func ApiTagAll(c *gin.Context) {
	r := models.GetAllTags()
	response.Ok(c, r)
}

func ApiTagPage(c *gin.Context) {
	in := models.Page{}
	err := c.BindQuery(&in)
	if err != nil {
		panic(err)
	}
	cnt := models.GetTagCnt()
	if cnt <= 0 {
		response.BzError(c, "标签为空")
		return
	}
	tags, err := models.TagGetPage(in.Pid, in.Psize)
	if err != nil {
		response.BzError(c, "没查到标签")
		return
	}
	response.Page(c, tags, int(cnt))
}

func ApiTagDrop(c *gin.Context) {
	input := struct {
		Id int `json:"id"`
	}{-1}
	err := c.BindJSON(&input)
	if err != nil {
		panic(err)
	}
	err = models.TagDrop(input.Id)
	if err != nil {
		response.BzError(c, "删除失败")
		return
	}
	response.Ok(c, "删除成功")
}

func ApiTagEdit(c *gin.Context) {
	tag := models.Tag{}
	err := c.BindJSON(&tag)
	if err != nil {
		panic(err)
	}
	err = models.TagEdit(&tag)
	if err != nil {
		response.BzError(c, "编辑失败")
		return
	}
	response.Ok(c, "编辑成功")
}

func ApiTagAdd(c *gin.Context) {
	tag := models.Tag{}
	err := c.BindJSON(&tag)
	if err != nil {
		panic(err)
	}
	err = models.TagAdd(&tag)
	if err != nil {
		response.BzError(c, "添加失败")
		return
	}
	response.Ok(c, "添加成功")
}
