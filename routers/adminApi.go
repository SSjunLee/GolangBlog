// @Author ljn 2022/4/28 16:16:00
package routers

import (
	web "Myblog/api"
	"Myblog/cmd"
	"github.com/gin-gonic/gin"
)

func AdminApi(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		if !cmd.Config.Nointer {
			admin.Use(LoginInterceptor())
		}
		admin.GET("/test", web.Test)
		admin.GET("/menu", web.ApiGetMenuList)
		admin.POST("/page/drop", web.ApiPageDrop)
		admin.POST("/page/add", web.ApiPageAdd)
		admin.POST("/page/edit", web.ApiPageEdit)
		admin.POST("/post/drop", web.ApiPostDrop)
		admin.POST("/post/add", web.ApiPostAdd)
		admin.POST("/post/edit", web.ApiPostEdit)
		admin.POST("/cate/edit", web.ApiCateEdit)
		admin.POST("/cate/add", web.ApiCateAdd)
		admin.POST("/cate/drop", web.ApiCateDrop)
		admin.POST("/upload/image", web.ApiImgUpload)
		admin.POST("/tag/drop", web.ApiTagDrop)
		admin.POST("/tag/edit", web.ApiTagEdit)
		admin.POST("/tag/add", web.ApiTagAdd)
		admin.GET("/meta/get", web.ApiMetaGet)
		admin.POST("/meta/edit", web.ApiMetaEdit)
	}
}
