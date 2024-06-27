package handler

// ErrCode 定义错误码类型
type ErrCode int

// 定义错误码常量
const (
	ErrCodeSuccess      ErrCode = 0    // 成功
	ErrCodeInternal     ErrCode = 1000 // 默认内部错误
	ErrCodeInvalidParam ErrCode = 1001 // 无效参数
)

// errMsgs 错误消息映射
var errMsgs = map[ErrCode]string{
	ErrCodeSuccess:      "Success",
	ErrCodeInternal:     "Internal Server Fail",
	ErrCodeInvalidParam: "Invalid Parameter",
}
