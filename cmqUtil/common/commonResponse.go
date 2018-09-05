package common

// CommonResponse 公共请求结果对象
type CommonResponse struct {
	// 错误码
	Code ResultStatus `json:"code"`

	// 错误提示信息
	Message string `json:"message"`

	// 服务器生成的请求Id
	RequestId string `json:"requestId"`
}

// IsSuccess 请求结果是否成功
func (this *CommonResponse) IsSuccess() bool {
	return this.Code == ResultStatus_Success
}

// NewCommonResponse 新建公共请求结果对象
// 参数
// code:错误码
// message:错误提示信息
// requestId:服务器生成的请求Id
// 返回值
// *CommonResponse:公共请求结果对象
func NewCommonResponse(code ResultStatus, message, requestId string) *CommonResponse {
	return &CommonResponse{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}
