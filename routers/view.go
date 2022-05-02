// @Author ljn 2022/4/28 19:06:00
package routers

import (
	"Myblog/view"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"time"
)

func Date(t time.Time, format string) string {
	return t.Format(format)
}

func Str2Html(s string) template.HTML {
	return template.HTML(s)
}

func Str2Js(s string) template.JS {
	return template.JS(s)
}

func Md5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

func RouterView(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"date":     Date,
		"str2Html": Str2Html,
		"str2js":   Str2Js,
		"md5":      Md5,
	})
	r.LoadHTMLGlob("ui/*")
	log.Println(r.FuncMap)
	r.GET("/", view.Index)
	r.GET("/archives", view.Archive)
	r.GET("/post/*path", view.Post)
	r.GET("/cate/:cate", view.CatePost)
	r.GET("/tags", view.TagsView)
	r.GET("/tag/:tag", view.TagPost)
	r.GET("/about", view.About)
	r.GET("/links", view.FriendLink)
}
