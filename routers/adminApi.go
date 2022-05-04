// @Author ljn 2022/4/28 16:16:00
package routers

import (
	"Myblog/api"
	"Myblog/cmd"
	"github.com/gin-gonic/gin"
)

func AdminApi(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		if !cmd.Config.Nointer {
			admin.Use(LoginInterceptor())
		}
		admin.GET("/userInfo", api.UserInfo)
		admin.GET("/test", api.Test)
		admin.GET("/menu", api.GetMenuList)
		admin.POST("/page/drop", api.PageDrop)
		admin.POST("/page/add", api.PageAdd)
		admin.POST("/page/edit", api.PageEdit)
		admin.POST("/post/drop", api.PostDrop)
		admin.POST("/post/add", api.PostAdd)
		admin.POST("/post/edit", api.PostEdit)
		admin.POST("/cate/edit", api.CateEdit)
		admin.POST("/cate/add", api.CateAdd)
		admin.POST("/cate/drop", api.CateDrop)
		admin.POST("/upload/image/oss", api.ImageUploadOss)
		admin.POST("/upload/image/local", api.ImgUploadLocal)
		admin.POST("/tag/drop", api.TagDrop)
		admin.POST("/tag/edit", api.TagEdit)
		admin.POST("/tag/add", api.TagAdd)
		admin.GET("/meta/get", api.MetaGet)
		admin.POST("/meta/edit", api.MetaEdit)
	}
}
