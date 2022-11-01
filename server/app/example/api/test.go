package api

import (
	"github.com/gin-gonic/gin"
)

type TestApi struct{}

// Demo
// @Summary 测试demo
// @Tags 测试
// @Success 200 {object} response.JsonResult
// @Security ApiKeyAuth
// @param tableNames body []string true "表名"
// @Router /test/demo [post]
func (testApi *TestApi) Demo(c *gin.Context) {

	//jwtUtil := utils.JwtUtil{}
	//tokenStr, err := jwtUtil.GenerateToken(10000, time.Now().Add(time.Second*-5))
	//if err != nil {
	//
	//}
	//fmt.Println(tokenStr)
	//token, err := jwtUtil.ParseToken(tokenStr)
	//if err != nil {
	//	fmt.Println("错误token")
	//	fmt.Println(err)
	//}
	//claims := token.Claims.(*utils.AuthClaims)
	//fmt.Println(claims)
}
