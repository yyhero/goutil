package configUtil

import (
	"fmt"
	"testing"

	"github.com/yyhero/goutil/xmlUtil"
)

// bool值读取测试
func TestBool(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	var ispost bool
	ispost, errMsg = xmlConfigData.Bool("html/body", "IsPost")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}
	fmt.Println("读取到的值:", ispost)
	if ispost == false {
		t.Error("html/body的isPost读取错误")
		t.Fail()
		return
	}

	ispost = xmlConfigData.DefaultBool("html/body", "IsPost", false)
	if ispost == false {
		t.Error("html/body的isPost读取错误")
		t.Fail()
	}
}

// int值读取测试
func TestInt(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	var id int
	id, errMsg = xmlConfigData.Int("html/body/ul/li/a[@id=1]", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}

	if id != 1 {
		t.Errorf("html/body的isPost读取错误，读取到的值:%v", id)
		t.Fail()
		return
	}

	id = xmlConfigData.DefaultInt("html/body", "id", 2)
	if id != 2 {
		t.Error("TestInt html/body的id读取错误")
		t.Fail()
	}
}

// int64值读取测试
func TestInt64(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	var id int64
	id, errMsg = xmlConfigData.Int64("html/body/ul/li/a[@id=1]", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}

	if id != 1 {
		t.Errorf("TestInt64 html/body/ul/li/a[@id=1]的id读取错误，读取到的值:%v", id)
		t.Fail()
		return
	}

	id = xmlConfigData.DefaultInt64("html/body", "id", 2)
	if id != 2 {
		t.Error("TestInt64 html/body的id读取错误")
		t.Fail()
	}
}

// Float值读取测试
func TestFloat(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	var id float64
	id, errMsg = xmlConfigData.Float("html/body/ul/li/a[@id=1]", "dd")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}

	if id != 1.1 {
		t.Errorf("TestFloat html/body/ul/li/a[@id=1]的id读取错误，读取到的值:%v", id)
		t.Fail()
		return
	}

	id = xmlConfigData.DefaultFloat("html/body", "id", 2)
	if id != 2 {
		t.Error("TestFloat html/body的id读取错误")
		t.Fail()
	}
}

// 字符串读取测试
func TestString(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	var id string
	id, errMsg = xmlConfigData.String("html/body/ul/li/a[@id=1]", "dd")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}

	if id != "1.1" {
		t.Errorf("TestString html/body/ul/li/a[@id=1]的id读取错误，读取到的值:%v", id)
		t.Fail()
		return
	}

	id = xmlConfigData.DefaultString("html/body", "id", "2")
	if id != "2" {
		t.Error("TestString html/body的id读取错误")
		t.Fail()
	}
}

type HelloStruct struct {
	// 连接字符串
	ConnectionString string

	// 最大开启连接数量
	MaxOpenConns int `xml:",attr"`

	// 最大空闲连接数量
	MaxIdleConns int `xml:",attr"`
}

func TestUnmarshal(t *testing.T) {
	data, _ := getxmlConfigData2(`
		<DBConnection>
		<GameServerCenterDB MaxOpenConns="10" MaxIdleConns="5">
			<![CDATA[
				root:moqikaka3312@tcp(10.1.0.10:3312)/2_gsc_develop?charset=utf8&parseTime=true&loc=Local&timeout=60s
			]]>	
		</GameServerCenterDB>
	</DBConnection>
		`)

	val := &HelloStruct{}
	err := data.Unmarshal("/DBConnection/GameServerCenterDB", val)
	if err != nil {
		t.Error(err)
		return
	}

	t.Errorf("和预期值不一致，结果为:Name:%v ", val)
}

type ConnsNum int
type Hello2Struct struct {
	// 连接字符串
	ConnectionString string

	// 最大开启连接数量
	MaxOpenConns ConnsNum

	// 最大空闲连接数量
	MaxIdleConns ConnsNum
}

func TestUnmarshal2(t *testing.T) {
	data, _ := getxmlConfigData2(`
		<DBConnection>
		<GameServerCenterDB MaxOpenConns="10" MaxIdleConns="5">
			<![CDATA[
				root:moqikaka3312@tcp(10.1.0.10:3312)/2_gsc_develop?charset=utf8&parseTime=true&loc=Local&timeout=60s
			]]>	
		</GameServerCenterDB>
	</DBConnection>
		`)

	val := &Hello2Struct{}
	err := data.Unmarshal("/DBConnection/GameServerCenterDB", val)
	if err != nil {
		t.Error(err)
		return
	}

	t.Errorf("和预期值不一致，结果为:Name:%v ", val)
}

func TestUnmarshal3(t *testing.T) {
	data, _ := getxmlConfigData2(`
		<DBConnection>
		<GameServerCenterDB MaxOpenConns="10">
			<![CDATA[
				root:moqikaka3312@tcp(10.1.0.10:3312)/2_gsc_develop?charset=utf8&parseTime=true&loc=Local&timeout=60s
			]]>	
		</GameServerCenterDB>
	</DBConnection>
		`)

	val := &HelloStruct{}
	err := data.Unmarshal("/DBConnection/GameServerCenterDB", val)
	if err != nil {
		t.Error(err)
		return
	}

	t.Errorf("和预期值不一致，结果为:Name:%v ", val)
}

func getxmlConfigData() (xmlConfigData *XmlConfig, errMsg error) {
	content := `
	<html lang="en">
		   <head>
			   <title>Hello</title>
			   <meta name="language" content="en"/>
		   </head>
		   <body IsPost='true'>
				<h1> This is a H1 </h1>
				<ul>
					<li><a id="1" dd='1.1' href="/">Home</a></li>
					<li><a id="2" href="/about">about</a></li>
					<li><a id="3" href="/account">login</a></li>
					<li></li>
				</ul>
				<p>
					Hello,This is an example for gxpath.
				</p>
				<footer>footer script</footer>
		   </body>
		</html>
	`
	var root *xmlUtil.Node
	root, errMsg = xmlUtil.LoadFromString(content)
	if errMsg == nil {
		xmlConfigData = NewXmlConfig()
		xmlConfigData.LoadFromXmlNode(root)
	}

	return
}

func getxmlConfigData2(xml string) (xmlConfigData *XmlConfig, errMsg error) {
	var root *xmlUtil.Node
	root, errMsg = xmlUtil.LoadFromString(xml)
	if errMsg == nil {
		xmlConfigData = NewXmlConfig()
		xmlConfigData.LoadFromXmlNode(root)
	}

	return
}
