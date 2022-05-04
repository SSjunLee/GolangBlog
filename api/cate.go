// @Author ljn 2022/4/26 15:07:00
package api

import (
	"Myblog/api/response"
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CateAll(c *gin.Context) {
	r := models.GetAllCate()
	response.Ok(c, r)
}

func CateGetPage(c *gin.Context) {
	in := models.BuildPageFromHttpParams(c)
	cnt := models.CateCnt()
	if cnt <= 0 {
		response.BzError(c, "分类为空")
		return
	}
	res, err := models.CateGetPage(in.Pid, in.Psize)
	if err != nil {
		response.BzError(c, "查找分类失败")
		return
	}
	response.Page(c, res, int(cnt))
}

func CateGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		panic(err)
	}
	r := models.CateGet(id)
	if r == nil {
		response.BzError(c, "不存在")
		return
	}
	response.Ok(c, r)
}

func CateEdit(c *gin.Context) {
	cate := models.Cate{}
	err := c.BindJSON(&cate)
	if err != nil {
		panic(err)
	}
	err = models.CateEdit(&cate)
	if err != nil {
		response.BzError(c, "编辑失败")
		return
	}
	response.Ok(c, "编辑成功")
}

func CateAdd(c *gin.Context) {
	cate := models.Cate{}
	err := c.BindJSON(&cate)
	if err != nil {
		panic(err)
	}
	err = models.CateAdd(&cate)
	if err != nil {
		response.BzError(c, "添加失败")
		return
	}
	response.Ok(c, "添加成功")
}

func CateDrop(c *gin.Context) {
	input := struct {
		Id int `json:"id"`
	}{-1}
	err := c.BindJSON(&input)
	if err != nil {
		panic(err)
	}
	err = models.CateDrop(input.Id)
	if err != nil {
		response.BzError(c, "删除失败")
		return
	}
	response.Ok(c, "删除成功")
}
