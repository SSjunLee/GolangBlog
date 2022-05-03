package web

import (
	"Myblog/api/response"
	"Myblog/cmd"
	"Myblog/common"
	"Myblog/common/utils"
	"Myblog/core/token"
	"Myblog/core/vcode"
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

const vcodeSecret = "v.c.o.d.e"

func ApiUserInfo(c *gin.Context) {
	uid, exits := c.Get(common.CTXUserId)
	if !exits {
		response.PleaseLogin(c)
		return
	}

	u := models.FetchUser(uid)
	if u == nil {
		response.PleaseLogin(c)
		return
	}
	response.Ok(c, u)
}

func ApiLogin(c *gin.Context) {
	in := struct {
		Username, Password, Vcode, Vreal string
	}{}
	err := c.ShouldBind(&in)
	if err != nil {
		panic(err)
	}

	if utils.ShaEncode(in.Vcode, vcodeSecret) != in.Vreal {
		response.BzError(c, "验证码不正确")
		return
	}
	if in.Username == "" || len(in.Username) > 18 {
		response.BzError(c, "用户名密码不正确")
		return
	}

	user := models.FetchUser(in.Username)
	if user == nil || user.Password != in.Password {
		response.BzError(c, "用户名密码不正确")
		return
	}
	auth := token.NewAuth(user.Id)
	response.Ok(c, auth.Encode(cmd.Config.JwtSecret))
}

func ApiRegister(c *gin.Context) {
	in := struct {
		Username, Password string
	}{}
	_ = c.ShouldBind(&in)
	if UsernameExits(in.Username) {
		response.BzError(c, "用户已存在")
		return
	}
	user := &models.User{Name: in.Username, Password: in.Password,
		Created: time.Now()}
	res := models.Db.Create(user)
	if res.Error != nil {
		panic(res.Error)
	}
	auth := token.NewAuth(user.Id)
	response.Ok(c, auth.Encode(cmd.Config.JwtSecret))
}

func ApiVCode(c *gin.Context) {
	rd := utils.RandomDigitStr(4)
	//log.Println(rd)
	vreal := utils.ShaEncode(rd, vcodeSecret)
	out := struct {
		VCode string `json:"vcode"`
		VReal string `json:"vreal"`
	}{
		VCode: vcode.NewImage(rd).Base64(),
		VReal: vreal,
	}
	response.Ok(c, out)
}

func UsernameExits(username string) bool {
	res := models.Db.Select("name").Where("name = ?", username).Find(&models.User{})
	if res.Error != nil {
		log.Panic(res.Error)
	}
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
