package api

//响应码

type ResCode int64

const (
	CodeSuccess         ResCode = 1000 //请求处理成功
	CodeInvalidParam    ResCode = 1001 //无效的请求参数
	CodeUserExist       ResCode = 1002 //用户已经存在
	CodeUserNotExist    ResCode = 1003 //用户不存在
	CodeInvalidPassword ResCode = 1004 //密码无效
	CodeServerBusy      ResCode = 1005 //服务忙
	CodeNeedLogin       ResCode = 1006 //未登录
	CodeInvalidToken    ResCode = 1007 //无效的Token
)

// 为错误码定义错误的原因描述
var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "请求处理成功",
	CodeInvalidParam:    "无效的请求参数",
	CodeUserExist:       "用户已经存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "用户未登录",
	CodeInvalidToken:    "无效的Token",
}

// 根据错误码，返回对应的消息
func (cd ResCode) Msg() string {
	msg, ok := codeMsgMap[cd]
	if !ok {
		//如果找不到对应的错误码，就返回繁忙
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
