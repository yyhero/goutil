package common

// IRequest 请求对象接口
type IRequest interface {
	// GetActionName 获取方法名
	GetActionName() string

	// SetCommonRequestObject 设置公共请求对象
	SetCommonRequest(commonRequest *CommonRequest)

	// GetParamMap 获取参数字典
	GetParamMap() map[string]interface{}
}
