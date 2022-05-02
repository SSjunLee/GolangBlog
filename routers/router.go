package routers

import (
	"Myblog/cmd"
	"github.com/gin-gonic/gin"
)

func RouterApp() {
	r := gin.New()
	r.Use(CORS(), midrecover())
	r.Static("/static", "static")
	RouterView(r)
	CommonApi(r)
	AdminApi(r)
	_ = r.Run(":" + string(cmd.Config.Port))
}
