package cmqUtil

import (
	"errors"
	"github.com/yyhero/goutil/cmqUtil/common"
)

// 创建队列请求对象
type createRequest struct {
	// 公共请求参数
	*common.CommonRequest

	// 队列名字
	queueName string

	// 最大堆积消息数
	maxMsgHeapNum int

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	pollingWaitSeconds int

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	visibilityTimeout int

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	maxMsgSize int

	// 消息保留周期。取值范围 60-1296000 秒（1min-15天），默认值 345600 (4 天)。
	msgRetentionSeconds int

	// 队列是否开启回溯消息能力，该参数取值范围 0-msgRetentionSeconds,即最大的回溯时间为消息在队列中的保留周期，0 表示不开启。
	rewindSeconds int

	// 可选参数，死信队列策略配置参数，不填默认不开启死信队列功能。必须是Json格式的字符串（json.dump后的值），包含参数有：deadLetterQueue、policy、maxReceiveCount、maxTimeToLive，具体定义如下字段描述。
	deadLetterPolicy string

	// 死信队列的名称，条件：必须是同地域、同帐号下的队列；该队列本身没有设置死信队列（不允许嵌套）；该队列被指定为其他队列的死信队列的次数未达到上限，目前上限为 6 个队列。
	deadLetterPolicy_deadLetterQueue string

	// 死信策略，0：消息被多次消费未删除；1：Time-To-Live 过期。
	deadLetterPolicy_policy int

	// Policy=0 时是必填参数。最大接收次数，支持设定值为 1~1000 次。
	deadLetterPolicy_maxReceiveCount int

	// Policy=1 时是必填参数。最大未消费过期时间，允许设置 5min-12 小时，单位为秒，且必须小于消息保留周期 msgRetentionSeconds 的值。
	deadLetterPolicy_maxTimeToLive int
}

// 获取请求方法名
func (this *createRequest) GetActionName() string {
	return "CreateQueue"
}

// SetCommonRequest 设置公共请求对象
func (this *createRequest) SetCommonRequest(commonRequest *common.CommonRequest) {
	this.CommonRequest = commonRequest
}

// GetParamMap 获取请求参数字典
// 返回值
// map[string]interface{}:请求参数字典
func (this *createRequest) GetParamMap() map[string]interface{} {
	paramMap := this.BuildParamMap()

	// 组装必须参数
	paramMap["queueName"] = this.queueName

	// 组装非必须参数
	if this.maxMsgHeapNum > 0 {
		paramMap["maxMsgHeapNum"] = this.maxMsgHeapNum
	}

	if this.pollingWaitSeconds > 0 {
		paramMap["pollingWaitSeconds"] = this.pollingWaitSeconds
	}

	if this.visibilityTimeout > 0 {
		paramMap["visibilityTimeout"] = this.visibilityTimeout
	}

	if this.maxMsgSize > 0 {
		paramMap["maxMsgSize"] = this.maxMsgSize
	}

	if this.msgRetentionSeconds > 0 {
		paramMap["msgRetentionSeconds"] = this.msgRetentionSeconds
	}

	if this.rewindSeconds > 0 {
		paramMap["rewindSeconds"] = this.rewindSeconds
	}

	if this.deadLetterPolicy != "" {
		paramMap["deadLetterPolicy"] = this.deadLetterPolicy
	}

	if this.deadLetterPolicy_deadLetterQueue != "" {
		paramMap["deadLetterPolicy::deadLetterQueue"] = this.deadLetterPolicy_deadLetterQueue
	}

	if this.deadLetterPolicy_policy > 0 {
		paramMap["deadLetterPolicy::policy"] = this.deadLetterPolicy_policy
	}

	if this.deadLetterPolicy_maxReceiveCount > 0 {
		paramMap["deadLetterPolicy::maxReceiveCount"] = this.deadLetterPolicy_maxReceiveCount
	}

	if this.deadLetterPolicy_maxTimeToLive > 0 {
		paramMap["deadLetterPolicy::maxTimeToLive"] = this.deadLetterPolicy_maxTimeToLive
	}

	return paramMap
}

// 创建队列请求返回结果对象
type createResponse struct {
	*common.CommonResponse

	// 队列的唯一标识Id
	QueueId string `json:"queueId"`
}

// 创建队列
// 参数
// _queueName:队列名
// 返回值
// error:错误对象
func (this *CMQObject) createQueue(_queueName string) error {
	// 新建请求对象
	reqObj := &createRequest{
		queueName: _queueName,
	}

	// 新建请求结果对象
	resObj := &createResponse{}

	// 发送请求
	if err := this.action(reqObj, &resObj); err != nil {
		return err
	}

	if resObj.IsSuccess() == false {
		return errors.New(resObj.Message)
	}

	return nil
}
