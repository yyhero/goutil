package cmqUtil

import (
	"errors"
	"moqikaka.com/goutil/cmqUtil/common"
)

// 消费消息请求对象
type receiveRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字(必须)
	queueName string

	// 本次请求的长轮询等待时间(非必须)
	pollingWaitSeconds int
}

// 获取请求方法名
func (this *receiveRequest) GetActionName() string {
	return "ReceiveMessage"
}

// SetCommonRequestObject 设置公共请求对象
func (this *receiveRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *receiveRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName

	// 组装非必须参数
	if this.pollingWaitSeconds > 0 {
		paramMap["pollingWaitSeconds"] = this.pollingWaitSeconds
	}

	return paramMap
}

// 消息对象
type msgInfo struct {
	// 本次消费的消息正文
	MsgBody string `json:"msgBody"`

	// 服务器生成消息的唯一标识Id
	MsgId string `json:"msgId"`

	// 每次消费返回唯一的消息句柄。用于删除该消息，仅上一次消费时产生的消息句柄能用于删除消息。
	ReceiptHandle string `json:"receiptHandle"`

	// 消费被生产出来，进入队列的时间
	EnqueueTime int64 `json:"enqueueTime"`

	// 第一次消费该消息的时间
	FirstDequeueTime int64 `json:"firstDequeueTime"`

	// 消息的下次可见（可再次被消费）时间
	NextVisibleTime int64 `json:"nextVisibleTime"`

	// 消息被消费的次数
	DequeueCount int64 `json:"dequeueCount"`
}

// 消费消息请求返回结果对象
type receiveResponse struct {
	*common.CommonResponse

	// 消息对象
	*msgInfo
}

// Receive 消费单条消息
// 返回值
// msg:消息内容
// receiptHandle:消息句柄
// err:错误对象
func (this *CMQObject) ReceiveMessage() (msg, receiptHandle string, err error) {
	// 新建请求对象
	reqObj := &receiveRequest{
		queueName: this.queueName,
	}

	// 新建请求结果对象
	resObj := &receiveResponse{}

	// 发送请求
	err = this.action(reqObj, &resObj)
	if err != nil {
		return
	}

	// 忽略掉没有消息的错误
	if resObj.Code == common.ResultStatus_NoMessage {
		return
	}

	if resObj.IsSuccess() == false {
		err = errors.New(resObj.Message)
		return
	}

	msg = resObj.MsgBody
	receiptHandle = resObj.ReceiptHandle

	return
}
