package middleware

import (
	"github.com/gin-gonic/gin"
	"go-fast-admin/server/common/dto/response"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()

			response.FailWithCode(http.StatusInternalServerError, ErrorToString(r), c)
			c.Abort() //终止操作
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
