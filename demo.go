package main

import (
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/mitmtools"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler"
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

	// 文件、内容整体替换
	// opts.AddHandler(&handler.ReplaceFileRule{
	// 	Pattern: "https://www.baidu.com/",
	// 	Content: []byte("我不是百度"),
	// })

	// 内容查找替换
	opts.AddHandler(&handler.ReplaceContentRule{
		Pattern:     "^https://www.baidu.com/$",
		FindContent: "百度一下，你就知道",
		ToContent:   "百度一下，你也不知道",
	})

	// 在 html 标签开始后插入 script 标签，并添加 js 代码：body 无就 head
	opts.AddHandler(&handler.AddScriptToHeadRule{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("console.log('我不是百度');"),
	})

	// 在 html 标签结束前插入 script 标签，并添加 js 代码：body 无就 head
	opts.AddHandler(&handler.AddScriptToTailRule{
		Pattern: "^https://www.baidu.com/$",
		Content: []byte("console.log('我不是百度');"),
	})

	// 在头部增加内容
	opts.AddHandler(&handler.AddContentToHeadRule{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Content: []byte("console.log(2);"),
	})

	// 在尾部增加内容
	opts.AddHandler(&handler.AddContentToTailRule{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Content: []byte("console.log(1);"),
	})

	// 添加请求头
	opts.AddHandler(&handler.AddResponseHeader{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Header:  map[string]string{"k": "v"},
	})

	// 删除请求头
	opts.AddHandler(&handler.RemoveResponseHeader{
		Pattern: "^https://tysf.cponline.cnipa.gov.cn/am/js/chunk-d7b9a01a.c7f12daa.js$",
		Header:  []string{"Last-Modified", "Content-Type"},
	})

	glog.Fatal(mitmtools.Start(opts))
}
