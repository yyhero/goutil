package common

// ResultStatus 部分错误码
type ResultStatus int32

const (
	// ResultStatus_Success 成功
	ResultStatus_Success ResultStatus = 0

	// ResultStatus_NoMessage 不存在消息
	ResultStatus_NoMessage ResultStatus = 10020
)
