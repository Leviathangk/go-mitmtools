package handler

import (
	"github.com/Leviathangk/go-glog/glog"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type ShowReq struct {
	baseHandler
	Pattern string // url 匹配规则
}

func (r *ShowReq) Response(f *proxy.Flow) {

	// 替换响应
	if IsMatch(r.Pattern, f.Request.URL.String()) {
		glog.Debugf("ShowReq 当前请求：%s\n", f.Request.URL)
	}
}

// Check 检查是否符合启动要求
func (r *ShowReq) Check() error {
	return nil
}
