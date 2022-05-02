// @Author ljn 2022/4/25 16:28:00 
package routers

import (
	"Myblog/api/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
)


func midrecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				stack := make([]byte, 1<<10)
				l := runtime.Stack(stack, false)
				_, _ = os.Stdout.Write(stack[:l])
				_ = c.Error(err)
				response.SysError(c)
			}
		}()
		c.Next()
	}
}
