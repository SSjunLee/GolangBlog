// @Author ljn 2022/4/28 16:15:00
package routers

import (
	"Myblog/api"
	"github.com/gin-gonic/gin"
)

func CommonApi(r *gin.Engine) {
	group := r.Group("/api")
	{
		group.POST("/login", api.Login)
		group.GET("/vcode", api.VCode)
		group.POST("/register", api.Register)
		group.GET("/page/page", api.GetPageInPage)
		group.GET("/page/get", api.PageGet)
		group.GET("/post/page", api.GetPostInPage)
		group.GET("/post/home/page", api.PostHomeApi)
		group.GET("/post/get", api.PostGet)
		group.GET("/cate/all", api.CateAll)
		group.GET("/cate/page", api.CateGetPage)
		group.GET("/cate/get", api.CateGet)
		group.GET("/tag/all", api.TagAll)
		group.GET("/tag/page", api.TagPage)

	}
}
