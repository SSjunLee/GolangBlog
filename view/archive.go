// @Author ljn 2022/5/1 11:44:00
package view

import (
	"Myblog/models"
	"github.com/gin-gonic/gin"
)

func Archive(c *gin.Context) {
	archive, err := models.PostArchive()
	if err != nil {
		c.Redirect(302, "/")
		return
	}
	render(c, "archive.html", gin.H{
		"Archives": archive,
	})
}
