package cmqUtil

import (
	"errors"
	"github.com/yyhero/goutil/cmqUtil/common"
)

// 批量消费消息请求对象
type batchReceiveRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字(必须)
	queueName string

	// 本次消费的消息数量。取值范围 1-16。(必须)
	numOfMsg int

	// 本次请求的长轮询等待时间(非必须)
	pollingWaitSeconds int
}

// 获取请求方法名
func (this *batchReceiveRequest) GetActionName() string {
	return "BatchReceiveMessage"
}

// SetCommonRequest 设置公共请求对象
func (this *batchReceiveRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *batchReceiveRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName
	paramMap["numOfMsg"] = this.numOfMsg

	// 组装非必须参数
	if this.pollingWaitSeconds > 0 {
		paramMap["pollingWaitSeconds"] = this.pollingWaitSeconds
	}

	return paramMap
}

// 批量消费消息请求返回结果对象
type batchReceiveResponse struct {
	*common.CommonResponse

	// message信息列表
	MsgInfoList []*msgInfo `json:"msgInfoList"`
}

// BatchReceive 批量消费单条消息
// 参数
// numOfMsg:本次消费的消息数量
// 返回值
// msgMap:消息字典(key:消息句柄;value:消息内容)
// err:错误对象
func (this *CMQObject) BatchReceiveMessage(_numOfMsg int) (msgMap map[string]string, err error) {
	msgMap = make(map[string]string)

	// 新建请求对象
	reqObj := &batchReceiveRequest{
		queueName: this.queueName,
		numOfMsg:  _numOfMsg,
	}

	// 新建请求结果对象
	resObj := &batchReceiveResponse{}

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

	// 组装返回
	for _, msgInfo := range resObj.MsgInfoList {
		msgMap[msgInfo.ReceiptHandle] = msgInfo.MsgBody
	}

	return
}
