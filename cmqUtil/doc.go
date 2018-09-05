package cmqUtil

/*
腾讯云cmq操作助手类
*/

// Author:L
// Create:2018-5-24 11:42:20
// History:
// 1、2018-5-29 10:22:20
// -- 重构代码(L)
// 2、2018-5-29 16:11:20
// -- 调整包结构、优化代码

/**************************使用说明*************************

1、新建cmq操作对象(详细见下面的API概览)
|
|--a、发送单条消息:调用SendMessage接口
|
|--b、批量发送消息:调用BatchSendMessage接口
|
|--c、消费单条消息:调用ReceiveMessage接口
|
|--d、批量消费消息:调用BatchReceiveMessage接口
|
|--e、删除单条消息:调用DeleteMessage接口
|
|--f、批量删除消息:调用BatchDeleteMessage接口
|
|--g、提供设置异步处理自动推送消息模式
--|
--|--调用SetAsyncHandleFunc设置异步处理方法->调用StartAsyncHandleFunc开始处理消息(消息会间隔一定时间主动调用设置的方法,需要方法处理后返回处理结果即处理成功的消息句柄列表)

**********************************************************/

/**************************API概览**************************

1、新建CMQ操作对象
	cmqUtil.NewCMQObject(_url, _queueName, _secretId, _secretKey, _signatureMethod string) *CMQObject
	参数说明
	_url:请求url,腾讯云提供
	_queueName:队列名
	_secretId:腾讯云API密钥Id
	_secretKey:腾讯云API密钥Key
	_signatureMethod:请求签名加密方法(支持HmacSHA256和HmacSHA1)
	返回值
	*CMQObject:CMQ操作对象

 **********************************************************/
