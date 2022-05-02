package token

import (
	"Myblog/cmd"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	_ "github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Auth struct {
	Id     int   `json:"id"`
	RoleId int   `json:"rid"`
	Exp    int64 `json:"exp"`
}

func NewAuth(uid int) Auth {
	newAuth := Auth{
		Id:  uid,
		Exp: time.Now().Add(cmd.Config.JwtExp * time.Minute).Unix(),
	}
	return newAuth
}

func (auth *Auth) Encode(key string) string {
	data, _ := json.Marshal(auth)
	bStr := base64.RawURLEncoding.EncodeToString(data)
	hm := hmac.New(sha1.New, []byte(key))
	hm.Write([]byte(bStr))
	sign := base64.RawURLEncoding.EncodeToString(hm.Sum(nil))
	return bStr + "." + sign
}

func (auth *Auth) Decode(raw, key string) error {
	parts := strings.Split(raw, ".")
	if len(parts) != 2 {
		return errors.New("非法的token " + raw)
	}
	hm := hmac.New(sha1.New, []byte(key))
	hm.Write([]byte(parts[0]))
	sign := base64.RawURLEncoding.EncodeToString(hm.Sum(nil))
	if sign != parts[1] {
		return errors.New("token非法")
	}
	data, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return errors.New("base64解码失败" + err.Error())
	}
	err = json.Unmarshal(data, auth)
	if err != nil {
		return errors.New("json解码失败" + err.Error())
	}
	if time.Now().Unix() > auth.Exp {
		return errors.New("token已过期")
	}
	return nil
}
