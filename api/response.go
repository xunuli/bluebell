package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//http请求的返回响应

// 响应的结构体
type ResponseData struct {
	//响应码
	Code ResCode `json:"code"`
	//响应消息
	Msg interface{} `json:"msg"`
	//响应数据
	Data interface{} `json:"data"`
}

// 请求处理错误的返回函数
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// 请求处理错误，而且有明确处理逻辑错误的消息的返回函数
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// 请求处理成功，返回带有对应的数据的函数
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
