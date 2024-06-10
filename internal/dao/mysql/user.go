package mysql

import (
	"bluebell/internal/model"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

//把每一步对数据库中user表的操作封装成函数
//供服务层根据业务需求调用

// 封装错误
var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvaildPassword = errors.New("用户或密码错误")
)

// 检查用户是否已经存在于数据库中
func CheckUserExist(username string) error {
	//根据用户名查询数据库
	var user *model.User
	mdb := GetDB()
	row := mdb.Model(user).Where("username = ?", username).Find(&user).RowsAffected
	fmt.Println(user, row)
	if row > 0 {
		//说明该用户已经存在
		return ErrorUserExist
	}
	return nil
}

// 向数据库中插入一条新的用户记录
func InsertUser(user *model.User) error {
	//数据库中的密码不能显示出来，需要加密
	user.Password = encryptPassword(user.Password)
	//将该记录创建至数据库表中
	err := GetDB().Model(user).Create(user).Error
	return err
}

// 对密码进行加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("bluebell.com"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// 登录校验
func Login(user *model.User) error {
	oPassword := user.Password
	//根据用户信息去查询是否已经注册到数据库中
	mdb := GetDB()
	result := mdb.Model(user).Where("username = ?", user.UserName).Find(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		//未查到，直接返回用户不存在
		return ErrorUserNotExist
	}

	//判断密码是否正确，需要转化成数据库中的密码格式判断
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvaildPassword
	}
	return nil
}
