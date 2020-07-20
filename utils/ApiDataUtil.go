package utils

import (
	"time"
)

// 统一的响应格式
type ApiData struct {
	Code    int         `json:"code"`    //0成功1异常提示2参数错误3账号冻结4token非法或者过期
	Msg     string      `json:"msg"`     //提示信息
	Data    interface{} `json:"data"`    //数组或对象信息
	SysTime int64       `json:"sysTime"` //系统时间
}

//空结构体
func SuccessObject() ApiData {
	return ApiData{
		0,
		"success",
		struct{}{},
		sysTime(),
	}
}

//空数组
func SuccessArray() ApiData {
	return ApiData{
		0,
		"success",
		[]string{},
		sysTime(),
	}
}

func SuccessMsg(msg string) ApiData {
	return ApiData{
		0,
		msg,
		nil,
		sysTime(),
	}
}

func SuccessReturn(data interface{}) ApiData {
	return ApiData{
		0,
		"success",
		data,
		sysTime(),
	}
}

// 错误提示
func ErrReturn(msg string) ApiData {
	return ApiData{
		1,
		msg,
		nil,
		sysTime(),
	}
}

// 参数错误
func ErrParam(msg string) ApiData {
	return ApiData{
		2,
		msg,
		nil,
		sysTime(),
	}
}

// 账户锁定
func ErrLock(msg string) ApiData {
	return ApiData{
		3,
		msg,
		nil,
		sysTime(),
	}
}

// 账户锁定
func ErrToken(msg string) ApiData {
	return ApiData{
		4,
		msg,
		nil,
		sysTime(),
	}
}


//系统时间,13位毫秒
var sysTime = func() int64 {
	return time.Now().UnixNano() / 1e6
}
