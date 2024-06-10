package service

import (
	"bluebell/internal/dao/mysql"
	"bluebell/internal/model"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// 注册
func SignUp(req *SignUpRequest) error {
	//1.判断用户是否已经存在
	if err := mysql.CheckUserExist(req.Username); err != nil {
		return err
	}
	//2.基于雪花算法生成用户UUID
	userID := snowflake.GenID()
	//3.构造一个User实例
	user := &model.User{
		UserId:   userID,
		UserName: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
		Email:    req.Email,
	}
	//保存进数据库
	err := mysql.InsertUser(user)

	return err
}

// 登录
func Login(req *LoginRequest) (token string, err error) {
	user := &model.User{
		UserName: req.Username,
		Password: req.Password,
	}
	//根据用户的登录信息，去校验是否用户已经注册，并拿到对应的id
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	//生成JWT
	tokenstring, err := jwt.GenToken(user.UserId, user.UserName)
	return tokenstring, err
}
