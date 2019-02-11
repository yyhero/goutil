package cmqUtil

import (
	"errors"
	"fmt"
	"github.com/yyhero/goutil/cmqUtil/common"
)

// 批量删除消息请求对象
type batchDeleteRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字
	queueName string

	// 上次消费消息时返回的消息句柄列表
	receiptHandleList []string
}

// 获取请求方法名
func (this *batchDeleteRequest) GetActionName() string {
	return "BatchDeleteMessage"
}

// SetCommonRequestObject 设置公共请求对象
func (this *batchDeleteRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *batchDeleteRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName

	// 组装消费消息时返回的消息句柄列表
	for index, msg := range this.receiptHandleList {
		key := fmt.Sprintf("receiptHandle.%d", index)
		paramMap[key] = msg
	}

	return paramMap
}

// 批量删除消息请求返回结果对象
type batchDeleteResponse struct {
	*common.CommonResponse

	// 无法成功删除的错误列表
	ErrorList []*errorInfo `json:"errorList"`
}

// 错误信息对象
type errorInfo struct {
	// 错误码
	Code int `json:"code"`

	// 错误提示信息
	Message string `json:"message"`

	// 对应删除失败的消息句柄
	ReceiptHandle string `json:"receiptHandle"`
}

// BatchDelete 批量删除消息
// 参数
// _receiptHandleList:消息句柄列表
// 返回值
// errorMap:删除错误的字典(key:删除失败的消息句柄;value:删除失败的原因)
// err:错误对象
func (this *CMQObject) BatchDeleteMessage(_receiptHandleList []string) (errorMap map[string]string, err error) {
	errorMap = make(map[string]string)

	// 新建请求对象
	reqObj := &batchDeleteRequest{
		queueName:         this.queueName,
		receiptHandleList: _receiptHandleList,
	}

	// 新建请求结果对象
	resObj := &batchDeleteResponse{}

	// 发送请求
	err = this.action(reqObj, &resObj)
	if err != nil {
		return
	}

	if resObj.IsSuccess() == false {
		err = errors.New(resObj.Message)
		return
	}

	// 组装返回
	for _, errInfo := range resObj.ErrorList {
		errorMap[errInfo.ReceiptHandle] = errInfo.Message
	}

	return
}
