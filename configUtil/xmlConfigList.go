package configUtil

import (
	"fmt"
	"strings"

	"github.com/yyhero/goutil/typeUtil"
)

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// []bool:结果
// error:错误信息
func (this *XmlConfig) BoolList(xpath string, attrName string) (result []bool, errMsg error) {
	result = make([]bool, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		return result, errMsg
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Bool(valItem)
		if errMsg != nil {
			return result, errMsg
		}

		result = append(result, resultItem)
	}

	return result, nil
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
// ifAddDefaultVal:如果某项值转换失败，是否把默认值添加到结果集合中
//　返回值：
// []bool:结果
func (this *XmlConfig) DefaultBoolList(xpath string, attrName string, defaultval bool, ifAddDefaultVal bool) []bool {
	result := make([]bool, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		if ifAddDefaultVal {
			result = append(result, defaultval)
		}

		return result
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Bool(valItem)
		if errMsg != nil {
			if ifAddDefaultVal {
				result = append(result, defaultval)
			}

			continue
		}

		result = append(result, resultItem)
	}

	return result
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// []int:结果
// error:错误信息
func (this *XmlConfig) IntList(xpath string, attrName string) ([]int, error) {
	result := make([]int, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		return result, errMsg
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Int(valItem)
		if errMsg != nil {
			return result, errMsg
		}

		result = append(result, resultItem)
	}

	return result, nil
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
// ifAddDefaultVal:如果某项值转换失败，是否把默认值添加到结果集合中
//　返回值：
// []int:结果
func (this *XmlConfig) DefaultIntList(xpath string, attrName string, defaultval int, ifAddDefaultVal bool) []int {
	result := make([]int, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		if ifAddDefaultVal {
			result = append(result, defaultval)
		}

		return result
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Int(valItem)
		if errMsg != nil {
			if ifAddDefaultVal {
				result = append(result, defaultval)
			}

			continue
		}

		result = append(result, resultItem)
	}

	return result
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// []int64:结果
// error:错误信息
func (this *XmlConfig) Int64List(xpath string, attrName string) ([]int64, error) {
	result := make([]int64, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		return result, errMsg
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Int64(valItem)
		if errMsg != nil {
			return result, errMsg
		}

		result = append(result, resultItem)
	}

	return result, nil
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
// ifAddDefaultVal:如果某项值转换失败，是否把默认值添加到结果集合中
//　返回值：
// []int64:结果
func (this *XmlConfig) DefaultInt64List(xpath string, attrName string, defaultval int64, ifAddDefaultVal bool) []int64 {
	result := make([]int64, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		if ifAddDefaultVal {
			result = append(result, defaultval)
		}

		return result
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Int64(valItem)
		if errMsg != nil {
			if ifAddDefaultVal {
				result = append(result, defaultval)
			}

			continue
		}

		result = append(result, resultItem)
	}

	return result
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// []float64:结果
// error:错误信息
func (this *XmlConfig) FloatList(xpath string, attrName string) ([]float64, error) {
	result := make([]float64, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		return result, errMsg
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Float64(valItem)
		if errMsg != nil {
			return result, errMsg
		}

		result = append(result, resultItem)
	}

	return result, nil
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
// ifAddDefaultVal:如果某项值转换失败，是否把默认值添加到结果集合中
//　返回值：
// []float64:结果
func (this *XmlConfig) DefaultFloatList(xpath string, attrName string, defaultval float64, ifAddDefaultVal bool) []float64 {
	result := make([]float64, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		if ifAddDefaultVal {
			result = append(result, defaultval)
		}

		return result
	}

	// 转换成指定类型
	for _, valItem := range valList {
		resultItem, errMsg := typeUtil.Float64(valItem)
		if errMsg != nil {
			if ifAddDefaultVal {
				result = append(result, defaultval)
			}

			continue
		}

		result = append(result, resultItem)
	}

	return result
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// string:结果
// error:错误信息
func (this *XmlConfig) StringList(xpath string, attrName string) ([]string, error) {
	// 获取值列表
	return this.getValList(xpath, attrName)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
// ifAddDefaultVal:如果某项值转换失败，是否把默认值添加到结果集合中
//　返回值：
// string:结果
func (this *XmlConfig) DefaultStringList(xpath string, attrName string, defaultval string, ifAddDefaultVal bool) []string {
	result := make([]string, 0)

	// 获取值列表
	valList, errMsg := this.getValList(xpath, attrName)
	if errMsg != nil {
		if ifAddDefaultVal {
			result = append(result, defaultval)
		}

		return result
	} else {
		return valList
	}
}

// 获取指定路径的之
// xpath:xpath路径
// attrName:要获取的属性值，如果为空，则返回内部文本
func (this *XmlConfig) getValList(xpath string, attrName string) ([]string, error) {
	result := make([]string, 0)

	targetNodeList := this.root.SelectElements(xpath)
	if targetNodeList == nil {
		return result, fmt.Errorf("no find target node:%v", xpath)
	}

	// 依次获取各个节点
	for _, nodeItem := range targetNodeList {
		val := ""
		if attrName == "" {
			val = strings.TrimSpace(nodeItem.InnerText())
		} else {
			val, _ = nodeItem.SelectAttr(attrName)
		}

		result = append(result, val)
	}

	return result, nil
}
