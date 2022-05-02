package utils

import (
	"math/rand"
	"time"
)

const (
	mask     = 1<<4 - 1
	chars    = "1234567890"
	charsLen = len(chars)
)

var rng = rand.NewSource(time.Now().UnixNano())

/* chars 36个字符
 * rng.Int63() 每次产出64bit的随机数,每次我们使用6bit(2^6=64) 可以使用10次
 */
func RandomDigitStr(ln int) string {
	buf := make([]byte, ln)
	for idx, cache, remain := 0, rng.Int63(), 10; idx < ln; idx++ {
		if remain == 0 {
			cache, remain = rng.Int63(), 10
		}
		buf[idx] = chars[int(cache&mask)%charsLen]
		cache >>= 6
		remain--
	}
	return string(buf)
}
