# 简介

Mitmtools 是用 go 语言实现的方便处理响应的工具集  
通过正则匹配 url 并进行对应的处理

# 安装

```
go get github.com/Leviathangk/go-mitmtools@latest
```

# Handler

主要分为请求处理、响应处理

## req

请求 handler

- ShowReq：打印请求
- ChangeHeader：修改请求头（注意大小写）
- ChangeCookie：修改请求 Cookie（注意大小写）

## resp

响应 handler

- ReplaceFile：替换全部响应体
- ReplaceContent：替换部分响应体
- AddContentToHead：在指定文件头部添加代码
- AddContentToTail：在指定文件尾部添加代码
- AddHeader：添加指定请求头
- RemoveHeader：移除指定请求头
- AddScriptToHead：在 html 标签（body、head，取其一）开始后插入一个 script 节点，里面是自己的代码
- AddScriptToTail：在 html 标签（body、head，取其一）结束前插入一个 script 节点，里面是自己的代码
- FindContent：输出含有指定字符的 url
- FindCookie：输出含有指定 响应 Cookie 的 url：匹配的就是 document.cookie 后的那部分
- FindHeader：输出含有指定 响应头 的 url：匹配的是响应头的 key
- ChangeHeader：修改响应头
- ChangeCookie：修改响应 Cookie

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
	opts.AddHandler(&resp.AddContentToTail{
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
	
	// 修改响应头
	opts.AddHandler(&resp.ChangeHeader{
		Pattern: "^https://www.baidu.com/$",
		Header:  map[string][]string{"Bdqid": {"baidu"}},
	})
	
	// 修改响应 cookie
	opts.AddHandler(&resp.ChangeCookie{
		Pattern: "^https://www.baidu.com/$",
		Cookie:  map[string]string{"H_PS_PSSID": "baidu"},
	})
	
	// 修改请求头
	opts.AddHandler(&req.ChangeHeader{
		Pattern: "^http://127.0.0.1:8877/headerTest$",
		Header:  map[string][]string{"X": {"qiandu"}},
	})
	
	// 修改请求 cookie
	opts.AddHandler(&req.ChangeCookie{
		Pattern: "^http://127.0.0.1:8877/cookieTest$",
		Cookie:  map[string]string{"x": "qiandu"},
	})

	glog.Fatal(mitmtools.Start(opts))
}
```