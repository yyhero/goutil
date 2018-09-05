package cmqUtil

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/goutil/logUtil"
	"testing"
	"time"
)

var (
	cmqObject *CMQObject
)

const (
	// 原始请求url
	con_SendMessageUrl = "http://cmq-queue-cd.api.qcloud.com"

	// 队列名
	con_QueueName = "moqikaka_test1"

	// API密钥Id
	con_SecretId = ""

	// API密钥key
	con_SecretKey = ""

	// 签名加密类型
	con_SignatureMethod = "HmacSHA256"
)

func init() {
	cmqObject = NewCMQObject(con_SendMessageUrl, con_QueueName, con_SecretId, con_SecretKey, con_SignatureMethod)

	logUtil.SetLogPath("E:\\TestProjects\\src\\moqikaka.com\\goutil\\cmqUtil\\Log")
}

type sendObject struct {
	Id int

	Name string
}

func TestSend(context *testing.T) {
	sendObj := &sendObject{
		Id:   1,
		Name: "测试中午",
	}

	data, _ := json.Marshal(sendObj)

	// 发送单条消息
	if err := cmqObject.SendMessage(string(data)); err != nil {
		fmt.Printf("发送失败:%v\r\n", err)
	} else {
		fmt.Println("发送消息成功")
	}
}

func TestBatchSend(context *testing.T) {
	if err := cmqObject.BatchSendMessage([]string{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9", "test10", "test11", "test12", "test13", "test14", "test15", "test16"}); err != nil {
		fmt.Printf("发送失败:%v\r\n", err)
	} else {
		fmt.Println("发送消息成功")
	}
}

func TestReceiveAndDelete(context *testing.T) {
	// 消费单条消息
	msg, receiptHandle, err := cmqObject.ReceiveMessage()
	if err != nil {
		fmt.Printf("消费单条消息失败,err=%v\r\n", err)
	} else {
		fmt.Printf("消费单条内容为:%s\r\n", msg)
	}

	// 删除单条消息
	if err := cmqObject.DeleteMessage(receiptHandle); err != nil {
		fmt.Printf("删除单条消息失败,err=%v\r\n", err)
	} else {
		fmt.Println("删除单条消息成功!")
	}
}

func TestBatchReceiveAndDelete(context *testing.T) {
	result := make([]string, 0)

	// 批量消费消息
	msgMap, err := cmqObject.BatchReceiveMessage(16)
	if err != nil {
		fmt.Printf("批量消费消息失败,err=%v\r\n", err)
	} else {
		fmt.Println("批量消费消息成功")

		for key, value := range msgMap {
			fmt.Printf("消息内容为:%s\r\n", value)
			result = append(result, key)
		}
	}

	// 批量删除消息
	errorMap, err := cmqObject.BatchDeleteMessage(result)
	if err != nil {
		fmt.Printf("批量删除消息失败,err=%v\r\n", err)
	} else if errorMap != nil && len(errorMap) > 0 {
		for handle, errMsg := range errorMap {
			fmt.Printf("批量删除消息失败,消息句柄=%s,失败原因=%s\r\n", handle, errMsg)
		}
	} else {
		fmt.Println("批量删除消息成功")
	}
}

func TestHandle(context *testing.T) {
	// 设置异步处理方法
	cmqObject.SetAsyncHandleFunc(asyncHandle, 2, 5, 2)

	// 开始异步处理
	cmqObject.StartAsyncHandleFunc(3)

	for {
		fmt.Println("异步处理消息中...")

		time.Sleep(time.Second * 1)
	}
}

func asyncHandle(data map[string]string) []string {
	result := make([]string, 0, len(data))

	for key, value := range data {
		fmt.Printf("处理消息:%s\r\n", value)

		result = append(result, key)
	}

	return result
}
