// @Author ljn 2022/5/1 15:15:00
package view

import (
	"Myblog/common"
	"Myblog/models"
	"github.com/gin-gonic/gin"
)

const imgTemplate = `<img class="lazy-load" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==" data-src="$1" alt="$2">`

func lazyLoadImg(raw string) string {
	return common.RegImg.ReplaceAllString(raw, imgTemplate)
}

func renderPage(c *gin.Context, path string) {
	page := models.GetPageByPath(path)
	if page == nil {
		c.Redirect(302, "/")
		return
	}
	page.RichText = lazyLoadImg(page.RichText)
	render(c, "page.html", gin.H{
		"Page": page,
		"Show": page.Status == 2,
	})
}

func About(c *gin.Context) {
	renderPage(c, "about")
}

func FriendLink(c *gin.Context) {
	renderPage(c, "links")
}
