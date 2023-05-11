package handler

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler/common"
	"strings"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type ReplaceContent struct {
	common.baseHandler
	Pattern     string // url 匹配规则
	FindContent string // 查找的需要替换的内容：针对部分内容
	ToContent   string // 替换后的新内容
}

func (r *ReplaceContent) Response(f *proxy.Flow) {

	// 替换响应
	if IsMatch(r.Pattern, f.Request.URL.String()) {
		f.Response.Body = []byte(strings.ReplaceAll(string(f.Response.Body), r.FindContent, r.ToContent))

		if ShowLog {
			glog.Debugf("ReplaceContent 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (r *ReplaceContent) Check() error {

	if r.FindContent == "" {
		return fmt.Errorf("ReplaceContent 未接收到需要替换的内容！")
	}

	return nil
}
