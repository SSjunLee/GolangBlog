package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func ShaEncode(raw, secrete string) string {
	hm := hmac.New(sha1.New, []byte(secrete))
	hm.Write([]byte(raw))
	return base64.RawURLEncoding.EncodeToString(hm.Sum(nil))
}
