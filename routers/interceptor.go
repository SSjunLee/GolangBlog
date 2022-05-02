package routers

import (
	"Myblog/api/response"
	"Myblog/cmd"
	"Myblog/common"
	"Myblog/core/token"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func loginInterceptorImpl(c *gin.Context) {
	tokenRaw := c.Request.Header.Get(common.JWTHeader)
	if tokenRaw == "" {
		response.PleaseLogin(c)
		return
	}
	auth := token.Auth{}
	err := auth.Decode(tokenRaw, cmd.Config.JwtSecret)
	if err != nil {
		response.PleaseLogin(c)
		return
	}

	if common.IsInTimeRange(time.Unix(auth.Exp, 0), cmd.Config.TokenRefreshMinute) {
		newAuth := token.NewAuth(auth.Id)
		newToken := newAuth.Encode(cmd.Config.JwtSecret)
		log.Println("刷新token")
		response.RefreshToken(c, newToken)
		return
	}
	c.Set(common.CTXUserId, auth.Id)
	log.Printf("用户访问%d api %s", auth.Id, c.Request.URL)
	c.Next()
}

func LoginInterceptor() gin.HandlerFunc {
	return loginInterceptorImpl
}
