package handler

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strings"
)

type FindContentRule struct {
	Rule
	Pattern string // url 匹配规则
	Content string // 需要查找的内容
}

func (fin *FindContentRule) Response(f *proxy.Flow) {
	if IsMatch(fin.Pattern, f.Request.URL.String()) {
		if strings.Index(string(f.Response.Body), fin.Content) != -1 {
			glog.Debugf("FindContentRule 已查找到：%s -> %s\n", string(fin.Content), f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindContentRule) Check() error {

	if len(fin.Content) == 0 {
		return fmt.Errorf("ReplaceFileRule 未接收到需要替换的内容！")
	}

	return nil
}
