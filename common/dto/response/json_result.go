package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResult struct {
	Code    int         `json:"code"`    //状态码（架构层面：中间件/异常处理等）
	Success bool        `json:"success"` //是否成功（业务层面）
	Data    interface{} `json:"data"`    //数据
	Message string      `json:"message"` //消息
}

func Result(code int, success bool, data interface{}, Message string, c *gin.Context) {
	c.JSON(http.StatusOK, JsonResult{
		code,
		success,
		data,
		Message,
	})
}

// Complete 完成
func Complete(data interface{}, err error, c *gin.Context) {
	if err != nil {
		FailWithDetail(nil, err.Error(), c)
	} else {
		SuccessWithDetail(data, "请求成功", c)
	}
}

// CompleteWithMessage 完成/消息
func CompleteWithMessage(err error, c *gin.Context) {
	if err != nil {
		FailWithDetail(nil, err.Error(), c)
	} else {
		SuccessWithDetail(nil, "请求成功", c)
	}
}

// SuccessWithDetail 成功详情
func SuccessWithDetail(data interface{}, message string, c *gin.Context) {
	Result(200, true, data, message, c)
}

// FailWithDetail 失败详情
func FailWithDetail(data interface{}, message string, c *gin.Context) {
	Result(200, false, data, message, c)
}

// FailWithMessage 失败消息
func FailWithMessage(message string, c *gin.Context) {
	FailWithDetail(nil, message, c)
}

// FailWithCode 失败代码消息
func FailWithCode(code int, message string, c *gin.Context) {
	Result(code, false, nil, message, c)
}

//func Success(c *gin.Context) {
//	SuccessWithData(nil, c)
//}
//
//func SuccessWithData(data interface{}, c *gin.Context) {
//	SuccessWithDetailed(data, "请求成功", c)
//}
//

//
//func Fail(c *gin.Context) {
//	FailWithData(nil, c)
//}
//

//
//func FailWithData(data interface{}, c *gin.Context) {
//	FailWithDetailed(data, "请求失败", c)
//}
//
