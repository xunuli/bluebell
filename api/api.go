package api

import (
	"bluebell/internal/dao/mysql"
	"bluebell/internal/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//请求进来，路由到对应的接口处理函数

// 用户注册
func SignUpHandler(c *gin.Context) {
	//1. 获取参数和参数校验
	req := &service.SignUpRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		//请求参数的格式有误，直接返回响应
		zap.L().Error("SignUp with invaild param", zap.Error(err)) //记录日志
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2. 业务处理，调用注册服务
	if err := service.SignUp(req); err != nil {
		zap.L().Error("service.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3. 返回响应
	ResponseSuccess(c, nil)
}

// 用户登录
func LoginHandler(c *gin.Context) {
	//1.获取参数以及参数校验
	req := &service.LoginRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.业务逻辑处理，登录
	token, err := service.Login(req)
	if err != nil {
		zap.L().Error("service.Login failed", zap.String("username", req.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(c, token)
}

// 查询到所有的社区(community_id, community_name)以列表的形式返回
func CommunityHandler(c *gin.Context) {
	//查询到所有的社区(community_id, community_name)以列表的形式返回
	fmt.Println("正在查询1。。。")
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端的错误暴露出去
		return
	}
	//返回成功和数据
	ResponseSuccess(c, data)
}

// 查询某个社区的具体详情
func CommunityDetailHandler(c *gin.Context) {
	//1.获取社区id
	idStr := c.Param("id") //获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	//2.根据id查询社区详情
	data, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("service.GetCommunityDetail fsiled", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, data)
}
