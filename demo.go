package main

import (
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/mitmtools"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler/req"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler/resp"
	"strconv"
)

const (
	port     = 8866
	proxyUrl = ""
)

func main() {
	opts := &mitmtools.MitmConfig{
		Addr:              ":" + strconv.Itoa(port),
		StreamLargeBodies: 1024 * 1024 * 5,
		SslInsecure:       true,
		Upstream:          proxyUrl,
		ShowLog:           true,
	}

	// 打印请求
	opts.AddHandler(&req.ShowReq{
		Pattern: "",
	})

	// 文件、内容整体替换
	opts.AddHandler(&resp.ReplaceFile{
		Pattern: "https://www.baidu.com/",
		Content: []byte("我不是百度"),
	})

	// 内容查找替换
	opts.AddHandler(&resp.ReplaceContent{
		Pattern:     "^https://www.baidu.com/$",
		FindContent: "百度一下，你就知道",
		ToContent:   "百度一下，你也不知道",
	})

	// 在 html 标签开始后插入 script 标签，并添加 js 代码：body 无就 head
	opts.AddHandler(&resp.AddScriptToHead{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("console.log('我不是百度');"),
	})

	// 在 html 标签结束前插入 script 标签，并添加 js 代码：body 无就 head
	opts.AddHandler(&resp.AddScriptToTail{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("console.log('我不是百度');"),
	})

	// 在头部增加内容
	opts.AddHandler(&resp.AddContentToHead{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Content: []byte("console.log(2);"),
	})

	// 在尾部增加内容
	opts.AddHandler(&resp.AddScriptToTail{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Content: []byte("console.log(1);"),
	})

	// 添加请求头
	opts.AddHandler(&resp.AddHeader{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Header:  map[string]string{"k": "v"},
	})

	// 删除请求头
	opts.AddHandler(&resp.RemoveHeader{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Header:  []string{"Last-Modified", "Content-Type"},
	})

	// 输出含有指定字符的 url
	opts.AddHandler(&resp.FindContent{
		Pattern:        "", // 为空则为任何响应
		ContentPattern: "百度一下",
	})

	// 输出含有指定 响应 Cookie 的 url：匹配的就是 document.cookie 后的那部分
	opts.AddHandler(&resp.FindCookie{
		Pattern:    "^https://www.baidu.com/$",
		KeyPattern: []string{"BAIDUID"},
	})

	// 输出含有指定 响应头 的 url：匹配的是响应头的 key
	opts.AddHandler(&resp.FindHeader{
		Pattern:    "^https://www.baidu.com/$",
		KeyPattern: []string{"Bdqid", "Set-Cookie"},
	})

	glog.Fatal(mitmtools.Start(opts))
}
