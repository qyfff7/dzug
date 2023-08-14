package models

type RespCode int32

const (
	CodeSuccess RespCode = 0 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误，请重试",
	CodeUserExist:       "用户名已存在，请重试",
	CodeUserNotExist:    "用户名不存在，请重试",
	CodeInvalidPassword: "用户名或密码错误，请重试",
	CodeServerBusy:      "服务繁忙，请稍后重试",
	CodeNeedLogin:       "请登录后再进行操作",
	CodeInvalidToken:    "当前token无效,请重新登录",
}

func (c RespCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
