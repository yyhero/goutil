package cmqUtil

import (
	"encoding/json"
	"fmt"
	"github.com/yyhero/goutil/cmqUtil/common"
	"github.com/yyhero/goutil/logUtil"
	"github.com/yyhero/goutil/mathUtil"
	"github.com/yyhero/goutil/webUtil"
	"time"
)

const (
	// 请求方法类型
	con_Method = "GET"

	// 请求url后缀
	con_UrlSuffix = "/v2/index.php"
)

// CMQObject CMQ对象
type CMQObject struct {
	// url地址
	url string

	// 队列名称
	queueName string

	// API密钥Id
	secretId string

	// API密钥key
	secretKey string

	// 请求方法类型
	method string

	// 签名加密类型
	signatureMethod string

	// 异步处理方法
	asyncHandleFunc func(map[string]string) []string

	// 异步处理每次处理的消息条数
	asyncHandleNumOfMsg int

	// 默认异步处理休眠时间
	asyncHandleSleepTime time.Duration

	// 异步处理成功多少次休眠一次
	asyncSleepSuccessCount int
}

//--------------------------------------------内部方法--------------------------------------------//

// 执行操作
func (this *CMQObject) action(messageInterface common.IRequest, responseObj interface{}) error {
	// 随机数字
	randObj := mathUtil.GetRand()

	// 组装公共请求对象
	commonObj := common.NewCommonRequest(messageInterface.GetActionName(), uint64(time.Now().Unix()), randObj.Uint32(), this.secretId, this.signatureMethod)
	messageInterface.SetCommonRequest(commonObj)

	// 组装请求url
	paramMap := messageInterface.GetParamMap()
	url, err := common.BuildUrl(this.method, this.url, this.secretKey, paramMap)
	if err != nil {
		return err
	}

	// 发送请求
	result, err := webUtil.GetWebData(url, nil)
	if err != nil {
		return err
	}

	// 解析请求结果
	err = json.Unmarshal(result, responseObj)

	return err
}

// 异步处理
func (this *CMQObject) asyncHandle(sleepTime *time.Duration, successCount *int) {
	defer func() {
		if err := recover(); err != nil {
			logUtil.ErrorLog("定时处理消息异常,err=%v", err)
		}
	}()

	logUtil.InfoLog("定时处理消息,当前时间:%v", time.Now())

	if this.asyncHandleFunc != nil {
		// 消费消息
		msgMap, err := this.BatchReceiveMessage(this.asyncHandleNumOfMsg)
		if err != nil {
			logUtil.ErrorLog("定时处理消息->批量消费消息错误,err=%v", err)

			return
		}

		// 处理消息
		successList := this.asyncHandleFunc(msgMap)

		// 删除消息
		if successList != nil && len(successList) > 0 {
			_, err = this.BatchDeleteMessage(successList)
			if err != nil {
				logUtil.ErrorLog("定时处理消息->批量删除消息错误,err=%v", err)

				return
			}

			// 消息处理成功后不休眠
			*sleepTime = 0
			*successCount++
		}
	}
}

//-----------------------------------------------------------------------------------------------//

//--------------------------------------------外部方法--------------------------------------------//

// GetUrl 获取url
func (this *CMQObject) GetUrl() string {
	return this.url
}

// GetQueueName 获取队列名称
func (this *CMQObject) GetQueueName() string {
	return this.queueName
}

// SetAsyncHandleFunc 设置异步处理方法
// 参数
// handleFunc:异步处理方法
// _asyncHandleNumOfMsg:异步处理每次处理的消息条数
// _asyncHandleSleepTime:默认异步处理休眠时间
// _asyncSleepSuccessCount:异步处理成功多少次休眠一次
func (this *CMQObject) SetAsyncHandleFunc(handleFunc func(map[string]string) []string, _asyncHandleNumOfMsg int, _asyncHandleSleepTime time.Duration, _asyncSleepSuccessCount int) {
	this.asyncHandleFunc = handleFunc
	this.asyncHandleNumOfMsg = _asyncHandleNumOfMsg
	this.asyncHandleSleepTime = _asyncHandleSleepTime
	this.asyncSleepSuccessCount = _asyncSleepSuccessCount
}

// StartAsyncHandleFunc 开始异步处理消息
// 参数
// runTimeNum:异步处理消息的协程数量
func (this *CMQObject) StartAsyncHandleFunc(runTimeNum int) {
	for i := 0; i < runTimeNum; i++ {
		go func() {
			// 设置休眠时间、处理成功次数
			sleepTime := this.asyncHandleSleepTime
			successCount := 0

			for {
				// 处理消息
				this.asyncHandle(&sleepTime, &successCount)

				// 当成功一定次数后,休眠一次
				if successCount >= this.asyncSleepSuccessCount {
					// 重置休眠时间、处理成功次数
					sleepTime = this.asyncHandleSleepTime
					successCount = 0
				}

				time.Sleep(sleepTime)
			}
		}()
	}
}

//----------------------------------------------------------------------------------------------//

//--------------------------------------------构造方法--------------------------------------------//

// NewCMQObject 新建CMQ对象
// 参数
// _url:请求url
// _queueName:队列名称
// _secretId:API密钥Id
// _secretKey:API密钥key
// _signatureMethod:签名加密类型
// 返回值
// *CMQObject:CMQ对象
func NewCMQObject(_url, _queueName, _secretId, _secretKey, _signatureMethod string) *CMQObject {
	return &CMQObject{
		url:             fmt.Sprintf("%s%s", _url, con_UrlSuffix),
		queueName:       _queueName,
		secretId:        _secretId,
		secretKey:       _secretKey,
		method:          con_Method,
		signatureMethod: _signatureMethod,
	}
}

// NewCMQObject1 新建CMQ对象(腾讯云不存在队列时)
// 参数
// _url:请求url
// _queueName:队列名称
// _secretId:API密钥Id
// _secretKey:API密钥key
// _signatureMethod:签名加密类型
// 返回值
// result:CMQ对象
// err:错误对象
func NewCMQObject1(_url, _queueName, _secretId, _secretKey, _signatureMethod string) (result *CMQObject, err error) {
	cmqObj := &CMQObject{
		url:             fmt.Sprintf("%s%s", _url, con_UrlSuffix),
		queueName:       _queueName,
		secretId:        _secretId,
		secretKey:       _secretKey,
		method:          con_Method,
		signatureMethod: _signatureMethod,
	}

	err = cmqObj.createQueue(_queueName)
	if err == nil {
		result = cmqObj
	}

	return
}

//----------------------------------------------------------------------------------------------//
