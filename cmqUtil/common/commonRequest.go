package common

// CommonRequest 公共请求参数对象
type CommonRequest struct {
	// Action 指令接口名称(必须)
	Action string

	// Region 地域参数
	Region string

	// Timestamp 当前UNIX时间戳(必须)
	Timestamp uint64

	// Nonce 随机正整数(必须)
	Nonce uint32

	// SecretId 在云API密钥上申请的标识身份的SecretId(必须)
	SecretId string

	// SignatureMethod 签名方式(非必须)
	SignatureMethod string

	// Token 临时证书所用的Token(非必须)
	Token string
}

// SetRegion 设置区域参数
func (this *CommonRequest) SetRegion(region string) {
	this.Region = region
}

// SetToken 设置token
func (this *CommonRequest) SetToken(token string) {
	this.Token = token
}

// BuildParamMap 组装请求参数字典
// 返回值
// map[string]interface{}:请求参数字
func (this *CommonRequest) BuildParamMap() map[string]interface{} {
	result := make(map[string]interface{})

	// 组装必要参数
	result["Action"] = this.Action
	result["Timestamp"] = this.Timestamp
	result["Nonce"] = this.Nonce
	result["SecretId"] = this.SecretId
	result["SignatureMethod"] = this.SignatureMethod

	// 区域
	if this.Region != "" {
		result["Region"] = this.Region
	}

	// Token
	if this.Token != "" {
		result["Token"] = this.Token
	}

	return result
}

// NewCommonRequest 新建公共请求参数对象
// 参数
// action:指令接口名称
// timestamp:当前UNIX时间戳
// nonce:随机正整数
// secretId:在云API密钥上申请的标识身份的SecretId
// signatureMethod:签名方式
// 返回值
// *CommonRequest:公共请求参数对象
func NewCommonRequest(action string, timestamp uint64, nonce uint32, secretId, signatureMethod string) *CommonRequest {
	result := &CommonRequest{
		Action:          action,
		Timestamp:       timestamp,
		Nonce:           nonce,
		SecretId:        secretId,
		SignatureMethod: signatureMethod,
	}

	return result
}
