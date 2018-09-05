package cmqUtil

import (
	"errors"
	"moqikaka.com/goutil/cmqUtil/common"
)

// 发送消息请求对象
type sendRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字
	queueName string

	// 消息正文
	msgBody string

	// 延时多久用户才可见该消息
	delaySeconds int
}

// 获取请求方法名
func (this *sendRequest) GetActionName() string {
	return "SendMessage"
}

// SetCommonRequest 设置公共请求对象
func (this *sendRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *sendRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName
	paramMap["msgBody"] = this.msgBody

	// 组装非必须参数
	if this.delaySeconds > 0 {
		paramMap["delaySeconds"] = this.delaySeconds
	}

	return paramMap
}

// 发送消息请求返回结果对象
type sendResponse struct {
	// 公共请求结果
	*common.CommonResponse

	// 服务器生成消息的唯一标识Id
	MsgId string `json:"msgId"`
}

// Send 发送单条消息
// 参数
// message:消息内容
// 返回值
// error:错误对象
func (this *CMQObject) SendMessage(message string) error {
	// 新建请求对象
	reqObj := &sendRequest{
		queueName: this.queueName,
		msgBody:   message,
	}

	// 新建请求结果对象
	resObj := &sendResponse{}

	// 发送请求
	if err := this.action(reqObj, resObj); err != nil {
		return err
	}

	if resObj.IsSuccess() == false {
		return errors.New(resObj.Message)
	}

	return nil
}
