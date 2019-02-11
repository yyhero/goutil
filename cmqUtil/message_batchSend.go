package cmqUtil

import (
	"errors"
	"fmt"
	"github.com/yyhero/goutil/cmqUtil/common"
)

// 批量发送消息请求对象
type batchSendRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字
	queueName string

	// 消息正文列表
	msgBodyList []string

	// 延时多久用户才可见该消息
	delaySeconds int
}

// 获取请求方法名
func (this *batchSendRequest) GetActionName() string {
	return "BatchSendMessage"
}

// SetCommonRequest 设置公共请求对象
func (this *batchSendRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *batchSendRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName

	// 组装消息列表
	for index, msg := range this.msgBodyList {
		key := fmt.Sprintf("msgBody.%d", index)
		paramMap[key] = msg
	}

	// 组装非必须参数
	if this.delaySeconds > 0 {
		paramMap["delaySeconds"] = this.delaySeconds
	}

	return paramMap
}

// 批量发送消息请求返回结果对象
type batchSendResponse struct {
	*common.CommonResponse

	// message信息列表
	MsgInfoList []*msgInfo `json:"msgInfoList"`
}

// BatchSend 批量发送消息
// 参数
// messageList:消息内容列表
// 返回值
// error:错误对象
func (this *CMQObject) BatchSendMessage(messageList []string) error {
	// 新建请求对象
	reqObj := &batchSendRequest{
		queueName:   this.queueName,
		msgBodyList: messageList,
	}

	// // 新建请求结果对象
	resObj := &batchSendResponse{}

	// 发送请求
	if err := this.action(reqObj, &resObj); err != nil {
		return err
	}

	if resObj.IsSuccess() == false {
		return errors.New(resObj.Message)
	}

	return nil
}
