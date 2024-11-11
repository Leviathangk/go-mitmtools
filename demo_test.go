/*
下面的 handler 可以按需测试
命令：go test -v demo_test.go
*/
package main

import (
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler/req"
	"github.com/Leviathangk/go-mitmtools/mitmtools"
	"testing"
)

const (
	Port     = 8866
	ProxyUrl = ""
)

func TestDemo(t *testing.T) {
	config := mitmtools.NewConfig(
		mitmtools.SetPort(Port),
		mitmtools.SetSslInsecure(true),
		mitmtools.SetProxy(ProxyUrl),
		mitmtools.SetShowLog(true),
	)

	// 打印请求
	config.AddHandler(&req.ShowReq{})

	/*
		查找替换修改
	*/
	// 修改请求 url（注意实际发出的编码的问题）
	//config.AddHandler(&req.ChangeUrl{
	//	Pattern:    "wd=%E7%8B%97",
	//	ReplaceVal: "wd=%E7%8C%AB",
	//})

	// 文件、内容整体替换
	//config.AddHandler(&resp.ReplaceFile{
	//	Pattern: "^https://www.baidu.com/$",
	//	Times:   1, // 0 为无限次
	//	Content: []byte("我不是百度"),
	//})

	// 文件、内容整体替换（仅当请求无 cookie 时）
	//config.AddHandler(&resp.ReplaceFileIfNoCookie{
	//	Pattern: "^https://www.baidu.com/$",
	//	Content: []byte("我不是百度"),
	//})

	// 内容查找替换
	//config.AddHandler(&resp.ReplaceContent{
	//	Pattern:     "^https://www.baidu.com/$",
	//	Times:       2, // 0 为无限次
	//	FindContent: "百度一下，你就知道",
	//	ToContent:   "百度一下，你也不知道",
	//})

	// 在 html 标签开始后插入 script 标签，并添加 js 代码：body 无就 head
	//config.AddHandler(&resp.AddScriptToHead{
	//	Pattern: "^https://www.baidu.com/$",
	//	Content: []byte("console.log('我不是百度');"),
	//})

	// 在 html 标签结束前插入 script 标签，并添加 js 代码：body 无就 head
	//config.AddHandler(&resp.AddScriptToTail{
	//	Pattern: "^https://www.baidu.com/$",
	//	Content: []byte("console.log('我不是百度');"),
	//})

	// 在头部增加内容
	//config.AddHandler(&resp.AddContentToHead{
	//	Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
	//	Content: []byte("console.log(2);"),
	//})

	// 在尾部增加内容
	//config.AddHandler(&resp.AddContentToTail{
	//	Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
	//	Content: []byte("console.log(1);"),
	//})

	/*
		查找
	*/
	// 输出含有指定字符的 url
	//config.AddHandler(&resp.FindContent{
	//	Pattern:        "", // 为空则为任何响应
	//	ContentPattern: "百度一下",
	//})

	// 输出含有指定 响应 Cookie 的 url：匹配的就是 document.cookie 后的那部分
	//config.AddHandler(&resp.FindCookie{
	//	Pattern:    "^https://www.baidu.com/$",
	//	KeyPattern: []string{"BAIDUID"},
	//})

	// 输出含有指定 响应头 的 url：匹配的是响应头的 key
	//config.AddHandler(&resp.FindHeader{
	//	Pattern:    "^https://www.baidu.com/$",
	//	KeyPattern: []string{"Bdqid", "Set-Cookie"},
	//})

	/*
		响应
	*/
	// 添加响应头
	//config.AddHandler(&resp.AddHeader{
	//	Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
	//	Header:  map[string]string{"k": "v"},
	//})

	// 删除响应头
	//config.AddHandler(&resp.RemoveHeader{
	//	Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
	//	Header:  []string{"Last-Modified", "Content-Type"},
	//})

	// 修改响应头
	//config.AddHandler(&resp.ChangeHeader{
	//	Pattern: "^https://www.baidu.com/$",
	//	Header:  map[string][]string{"Bdqid": {"baidu"}},
	//})

	// 修改响应 cookie
	//config.AddHandler(&resp.ChangeCookie{
	//	Pattern: "^https://www.baidu.com/$",
	//	Cookie:  map[string]string{"H_PS_PSSID": "baidu"},
	//})

	/*
		请求
	*/
	// 修改请求头
	//config.AddHandler(&req.ChangeHeader{
	//	Pattern: "^http://127.0.0.1:8877/headerTest$",
	//	Header:  map[string][]string{"X": {"qiandu"}},
	//})

	// 修改请求 cookie
	//config.AddHandler(&req.ChangeCookie{
	//	Pattern: "^http://127.0.0.1:8877/cookieTest$",
	//	Cookie:  map[string]string{"x": "qiandu"},
	//})

	glog.DLogger.Fatalln(mitmtools.Start(config))
}
