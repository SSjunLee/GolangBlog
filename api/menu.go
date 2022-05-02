package web

import (
	"Myblog/api/response"
	"Myblog/common"
	"Myblog/models"
	"github.com/gin-gonic/gin"
)

func GetMenuList(id int) []models.Menu {
	return models.FetchMenuByUser(id)
}

func ApiGetMenuList(c *gin.Context) {
	id, _ := c.Get(common.CTXUserId)
	menus := GetMenuList(id.(int))
	response.Ok(c, menus)
}
