package token

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	auth := Auth{
		Id:     1,
		RoleId: 1000,
		Exp:    time.Now().Add(time.Hour * 24).Unix(),
	}
	t.Log(auth.Encode("key"))
}

func TestVerify(t *testing.T) {
	raw := "eyJpZCI6MSwicmlkIjoxMDAwLCJleHAiOjE2NTA4MTM2MjN9.QEp7k4D0hD7Q6IFNcgWnXS0nEx0"
	auth := Auth{}
	err := auth.Decode(raw, "key")
	t.Log(auth, err)
}
