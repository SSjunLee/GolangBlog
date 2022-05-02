// @Author ljn 2022/4/28 16:15:00
package routers

import (
	web "Myblog/api"
	"github.com/gin-gonic/gin"
)

func CommonApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/login", web.ApiLogin)
		api.GET("/vcode", web.ApiVCode)
		api.POST("/register", web.ApiRegister)
		api.GET("/page/page", web.ApiGetPageInPage)
		api.GET("/page/get", web.ApiPageGet)
		api.GET("/post/page", web.ApiGetPostInPage)
		api.GET("/post/home/page", web.ApiPostHomeApi)
		api.GET("/post/get", web.ApiPostGet)
		api.GET("/cate/all", web.ApiCateAll)
		api.GET("/cate/page", web.ApiCateGetPage)
		api.GET("/cate/get", web.ApiCateGet)
		api.GET("/tag/all", web.ApiTagAll)
		api.GET("/tag/page", web.ApiTagPage)

	}
}
