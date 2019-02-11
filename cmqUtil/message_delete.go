package cmqUtil

import (
	"errors"
	"github.com/yyhero/goutil/cmqUtil/common"
)

// 删除消息请求对象
type deleteRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字(必须)
	queueName string

	// 上次消费返回唯一的消息句柄，用于删除消息。(必须)
	receiptHandle string
}

// 获取请求方法名
func (this *deleteRequest) GetActionName() string {
	return "DeleteMessage"
}

// SetCommonRequest 设置公共请求对象
func (this *deleteRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *deleteRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName
	paramMap["receiptHandle"] = this.receiptHandle

	return paramMap
}

// 删除消息请求返回结果对象
type deleteResponse struct {
	*common.CommonResponse
}

// Delete 删除单条消息
// 参数
// _receiptHandle:消息句柄
// 返回值
// error:错误对象
func (this *CMQObject) DeleteMessage(_receiptHandle string) error {
	// 新建请求对象
	reqObj := &deleteRequest{
		queueName:     this.queueName,
		receiptHandle: _receiptHandle,
	}

	// 新建请求结果对象
	resObj := &deleteResponse{}

	// 发送请求
	if err := this.action(reqObj, &resObj); err != nil {
		return err
	}

	if resObj.IsSuccess() == false {
		return errors.New(resObj.Message)
	}

	return nil
}
