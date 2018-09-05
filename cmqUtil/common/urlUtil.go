package common

import (
	"errors"
	"fmt"
	"net/url"

	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/securityUtil"
	"moqikaka.com/goutil/stringUtil"
	"sort"
)

// BuildUrl 组装请求url
// 参数
// method:请求方法
// _url:原始请求url
// secretKey:密钥的key
// paramMap:参数字典
// 返回值
// string:组装好的请求url
func BuildUrl(method, _url, secretKey string, paramMap map[string]interface{}) (resultUrl string, err error) {
	// 解析原始请求字符串(去除头部的http://或https://)
	signatureUrlStr := stringUtil.Split(_url, []string{"://"})[1]

	// 参数key排序
	paramKeys := sort.StringSlice{}
	for key := range paramMap {
		paramKeys = append(paramKeys, key)
	}
	sort.Sort(paramKeys)

	paramStr := ""
	for _, key := range paramKeys {
		paramStr += fmt.Sprintf("&%s=%v", key, paramMap[key])
	}
	paramStr = paramStr[1:]

	// 读取签名方法
	signatureMethod, exists := paramMap["SignatureMethod"]
	if exists == false {
		logUtil.ErrorLog("签名方法不存在,url=%s", _url)
		err = errors.New("签名方法不存在!")

		return
	}

	// 计算请求签名
	signatureUrlStr = fmt.Sprintf("%s%s?%s", method, signatureUrlStr, paramStr)
	signature, err := calcSignature(signatureMethod.(string), signatureUrlStr, secretKey)
	if err != nil {
		return
	}

	if signature == "" {
		logUtil.ErrorLog("计算得到的签名为空,url=%s", _url)
		err = errors.New("计算得到的签名为空!")

		return
	}

	// 组装最终请求url
	resultUrl = fmt.Sprintf("%s?%s&Signature=%s", _url, paramStr, signature)

	return
}

// 计算签名
// 参数
// signatureMethod:签名方法
// url:请求url
// secretKey:密钥key
// 返回值
// string:计算得到的签名
func calcSignature(signatureMethod, _url, secretKey string) (result string, err error) {
	switch signatureMethod {
	case "HmacSHA256":
		data, err := securityUtil.HmacSha256(_url, secretKey)
		if err != nil {
			return "", err
		}

		// 设置请求签名
		result = url.QueryEscape(string(stringUtil.Base64Encode2(data)))

	case "HmacSHA1":
		data, err := securityUtil.HmacSha1(_url, secretKey)
		if err != nil {
			return "", err
		}

		// 设置请求签名
		result = url.QueryEscape(string(stringUtil.Base64Encode2(data)))

	default:
		err = errors.New("不支持的签名方式")
	}

	return
}
