package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type m struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SysError(c *gin.Context) {
	Error(c, 500, "系统错误")
}

func BzError(c *gin.Context, msg string) {
	c.JSON(200, &m{
		Code: 888,
		Msg:  msg,
	})
	c.Abort()
}

const (
	loginPlease = 333
	refresh     = 401
)

func Error(c *gin.Context, code int, msg string) {
	c.JSON(200, &m{
		Code: code,
		Msg:  msg,
	})
	c.Abort()
}

func NotFound(c *gin.Context) {
	Error(c, 404, "没找到")
}

func PleaseLogin(c *gin.Context) {
	Error(c, loginPlease, "请登录")
}

func Page(c *gin.Context, data interface{}, len int) {
	c.JSON(200, &m{
		Code: 200,
		Data: struct {
			List interface{} `json:"list"`
			Len  int         `json:"len"`
		}{List: data,
			Len: len,
		},
	})
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(200, &m{
		Code: 200,
		Data: data,
	})
}

func RefreshToken(c *gin.Context, token string) {
	c.JSON(200, &m{
		Code: refresh,
		Data: token,
	})
	c.Abort()
}

func Forbidden(c *gin.Context, msg string) {
	Error(c, http.StatusForbidden, msg)
}
