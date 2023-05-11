# 简介

Mitmtools 是用 go 语言实现的方便处理响应的工具集  
通过正则匹配 url 并进行对应的处理

# 安装

```
go get github.com/Leviathangk/go-mitmtools@latest
```

# Handler

主要有以下 Handler

## req

请求 handler

- ShowReq：

## resp

响应 handler

- ReplaceFile：替换全部响应体
- ReplaceContent：替换部分响应体
- AddContentToHead：在指定文件头部添加代码
- AddScriptToTail：在指定文件尾部添加代码
- AddHeader：添加指定请求头
- RemoveHeader：移除指定请求头
- AddScriptToHead：在 html 标签（body、head，取其一）开始后插入一个 script 节点，里面是自己的代码
- AddScriptToTail：在 html 标签（body、head，取其一）结束前插入一个 script 节点，里面是自己的代码
- FindContent：输出含有指定字符的 url
- FindCookie：输出含有指定 响应 Cookie 的 url：匹配的就是 document.cookie 后的那部分
- FindHeader：输出含有指定 响应头 的 url：匹配的是响应头的 key

# 案例

```
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
	// opts.AddHandler(&resp.ReplaceFile{
	// 	  Pattern: "https://www.baidu.com/",
	// 	  Content: []byte("我不是百度"),
	// })

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
		Pattern: "^https://.js$",
		Content: []byte("console.log(2);"),
	})

	// 在尾部增加内容
	opts.AddHandler(&resp.AddScriptToTail{
		Pattern: "^https://.js$",
		Content: []byte("console.log(1);"),
	})

	// 添加请求头
	opts.AddHandler(&resp.AddHeader{
		Pattern: "^https://.js$",
		Header:  map[string]string{"k": "v"},
	})

	// 删除请求头
	opts.AddHandler(&resp.RemoveHeader{
		Pattern: "^https://.js$",
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
```