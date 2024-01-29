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

- ReplaceFile：替换全部响应体（可记次，默认无限次）
- ReplaceFileIfNoCookie：替换全部响应体（仅当无 cookie 时替换）
- ReplaceContent：替换部分响应体（可记次，默认无限次）
- ReplaceContentIfNoCookie：替换部分响应体（仅当无 cookie 时替换）
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
)

const (
	port     = 8866
	proxyUrl = ""
)

func main() {
	config := mitmtools.NewConfig(
		mitmtools.SetAddr("", port),
		mitmtools.SetSslInsecure(true),
		mitmtools.SetProxy(proxyUrl),
		mitmtools.SetShowLog(true),
	)

	// 打印请求
	config.AddHandler(&req.ShowReq{
		Pattern: "",
	})

	// 文件、内容整体替换
	config.AddHandler(&resp.ReplaceFile{
		Pattern: "^https://www.baidu.com/$",
		Times:   1, // 0 为无限次
		Content: []byte("我不是百度"),
	})

	// 文件、内容整体替换（仅当请求无 cookie 时）
	config.AddHandler(&resp.ReplaceFileIfNoCookie{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("我不是百度"),
	})

	// 内容查找替换
	config.AddHandler(&resp.ReplaceContent{
		Pattern:     "^https://www.baidu.com/$",
		Times:       1, // 0 为无限次
		FindContent: "百度一下，你就知道",
		ToContent:   "百度一下，你也不知道",
	})

	// 内容查找替换（仅当请求无 cookie 时）
	config.AddHandler(&resp.ReplaceContentIfNoCookie{
		Pattern:     "^https://www.baidu.com/$",
		FindContent: "百度一下，你就知道",
		ToContent:   "百度一下，你也不知道",
	})

	// 在 html 标签开始后插入 script 标签，并添加 js 代码：body 无就 head
	config.AddHandler(&resp.AddScriptToHead{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("console.log('我不是百度');"),
	})

	// 在 html 标签结束前插入 script 标签，并添加 js 代码：body 无就 head
	config.AddHandler(&resp.AddScriptToTail{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("console.log('我不是百度');"),
	})

	// 在头部增加内容
	config.AddHandler(&resp.AddContentToHead{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Content: []byte("console.log(2);"),
	})

	// 在尾部增加内容
	config.AddHandler(&resp.AddContentToTail{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Content: []byte("console.log(1);"),
	})

	// 添加请求头
	config.AddHandler(&resp.AddHeader{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Header:  map[string]string{"k": "v"},
	})

	// 删除请求头
	config.AddHandler(&resp.RemoveHeader{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Header:  []string{"Last-Modified", "Content-Type"},
	})

	// 输出含有指定字符的 url
	config.AddHandler(&resp.FindContent{
		Pattern:        "", // 为空则为任何响应
		ContentPattern: "百度一下",
	})

	// 输出含有指定 响应 Cookie 的 url：匹配的就是 document.cookie 后的那部分
	config.AddHandler(&resp.FindCookie{
		Pattern:    "^https://www.baidu.com/$",
		KeyPattern: []string{"BAIDUID"},
	})

	// 输出含有指定 响应头 的 url：匹配的是响应头的 key
	config.AddHandler(&resp.FindHeader{
		Pattern:    "^https://www.baidu.com/$",
		KeyPattern: []string{"Bdqid", "Set-Cookie"},
	})

	// 修改响应头
	config.AddHandler(&resp.ChangeHeader{
		Pattern: "^https://www.baidu.com/$",
		Header:  map[string][]string{"Bdqid": {"baidu"}},
	})

	// 修改响应 cookie
	config.AddHandler(&resp.ChangeCookie{
		Pattern: "^https://www.baidu.com/$",
		Cookie:  map[string]string{"H_PS_PSSID": "baidu"},
	})

	// 修改请求头
	config.AddHandler(&req.ChangeHeader{
		Pattern: "^http://127.0.0.1:8877/headerTest$",
		Header:  map[string][]string{"X": {"qiandu"}},
	})

	// 修改请求 cookie
	config.AddHandler(&req.ChangeCookie{
		Pattern: "^http://127.0.0.1:8877/cookieTest$",
		Cookie:  map[string]string{"x": "qiandu"},
	})

	glog.DLogger.Fatal(mitmtools.Start(config))
}
```

# 热重载

使用热重载会大大的方便调试  
这里推荐项目：[air](https://github.com/cosmtrek/air)  
注意：这里是 install 而不是 get

```
go install github.com/cosmtrek/air@latest
```

## 配置文件

这里可以直接使用 air init 方式创建，或者使用以下文件

```
# air.toml

[build]
# 指定入口文件
entry = "main.go"
# 监听目录
watch = ["."]
# 指定命令
cmd = "go run main.go"
# 延迟重新构建
delay = 3000
```

# 执行

直接在项目下面输入以下命令即可

```
air
```