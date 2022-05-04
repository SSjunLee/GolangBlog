package api

import (
	"Myblog/api/response"
	"Myblog/common"
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostHomeApi(c *gin.Context) {
	cid, err := strconv.Atoi(c.Query("cid"))
	if err != nil {
		cid = -1
	}
	in := models.Page{}
	err = c.BindQuery(&in)
	if err != nil {
		panic("参数错误")
	}
	err = in.Check()
	if err != nil {
		panic("参数错误")
	}
	cnt := models.PostCount(models.KindArticle, cid)
	if cnt < 1 {
		response.BzError(c, "没查到数据")
		return
	}
	posts := make([]models.Post, 0)
	posts, err = models.PostGetPageWithTags(models.KindArticle, cid, in.Pid, in.Psize)
	if err != nil {
		panic(err)
	}
	if len(posts) < 1 {
		response.BzError(c, "没查到数据")
		return
	}
	type postsWrapper struct {
		*models.Post
		Cate *models.Cate `json:"cate"`
	}
	pw := make([]postsWrapper, len(posts))
	for i, _ := range posts {
		cate := models.CateGet(posts[i].CatId)
		pw[i].Post = &posts[i]
		pw[i].Cate = cate
	}
	response.Page(c, pw, cnt)
}

func GetPostInPage(c *gin.Context) {
	cid, err := strconv.Atoi(c.Query("cid"))
	if err != nil {
		cid = -1
	}
	in := models.Page{}
	err = c.BindQuery(&in)
	//log.Println(cid, in)
	if err != nil {
		panic("参数错误")
	}
	err = in.Check()
	if err != nil {
		panic("参数错误")
	}
	cnt := models.PostCount(models.KindArticle, cid)
	args := []string{"id", "title", "summary", "updated", "status", "path"}
	if cnt < 1 {
		response.BzError(c, "没查到数据")
		return
	}
	posts := models.PostGetPage(models.KindArticle, cid, in.Pid, in.Psize, args...)

	if posts == nil {
		response.BzError(c, "没查到数据")
		return
	}
	response.Page(c, posts, cnt)
	//response.Ok(c, posts)
}

func PostDrop(c *gin.Context) {
	PageDrop(c)
}

func PostAdd(c *gin.Context) {
	p := models.Post{}
	err := c.BindJSON(&p)
	if err != nil {
		panic(err)
	}
	if p.PathExits() {
		response.BzError(c, "博客路径已存在")
		return
	}
	p.HandleRichText()
	//log.Println(p)
	p.Created.Format(common.StdDateTime)
	p.Updated = p.Created
	p.Kind = models.KindArticle
	err = models.PostAdd(&p)
	if err != nil {
		panic(err)
	}
	response.Ok(c, "添加成功")
}

func PostGet(c *gin.Context) {
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
	//tags := models.GetTagsByPostId(id)
	/*if len(tags) > 0 {
		p.Tags = tags
		//copy(p.Tags,tags)
	}*/
	response.Ok(c, p)
}

func PostEdit(c *gin.Context) {
	PageEdit(c)
}
