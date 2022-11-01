package middleware

import (
	"gin-fast-admin/server/common/dto/response"
	"gin-fast-admin/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.FailWithCode(http.StatusUnauthorized, "无权限访问，请求未携带token", context)
			context.Abort() //结束后续操作
			return
		}

		// 解析token
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			response.FailWithCode(http.StatusUnauthorized, "无权限访问，错误token", context)
			context.Abort() //结束后续操作
			return
		}
		context.Set("claims", claims)
		context.Next()
	}
}
