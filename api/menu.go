package api

import (
	"Myblog/api/response"
	"Myblog/common"
	"Myblog/models"
	"github.com/gin-gonic/gin"
)

func getMenuList(id int) []models.Menu {
	return models.FetchMenuByUser(id)
}

func GetMenuList(c *gin.Context) {
	id, _ := c.Get(common.CTXUserId)
	menus := getMenuList(id.(int))
	response.Ok(c, menus)
}
