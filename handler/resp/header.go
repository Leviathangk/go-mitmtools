package resp

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type AddHeader struct {
	handler.BaseHandler
	Pattern string            // url 匹配规则
	Header  map[string]string // 需要添加的 map
}

func (a *AddHeader) Responseheaders(f *proxy.Flow) {
	if handler.IsMatch(a.Pattern, f.Request.URL.String()) {
		for key, value := range a.Header {
			f.Response.Header.Set(key, value)

			if handler.ShowLog || a.ShowLog {
				glog.DLogger.Debugf("AddHeader 正在添加指定请求头：%s -> %s\n", key, value)
			}
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddHeader) Check() error {

	if a.Header == nil {
		return fmt.Errorf("AddHeader 未接收到需要添加的请求头！")
	}

	return nil
}

type RemoveHeader struct {
	handler.BaseHandler
	Pattern string   // url 匹配规则
	Header  []string // 需要移除的 key
}

func (r *RemoveHeader) Responseheaders(f *proxy.Flow) {
	if handler.IsMatch(r.Pattern, f.Request.URL.String()) {
		for _, key := range r.Header {
			if _, ok := f.Response.Header[key]; ok {
				delete(f.Response.Header, key)

				if handler.ShowLog || r.ShowLog {
					glog.DLogger.Debugf("RemoveHeader 正在移除指定请求头：%s\n", key)
				}
			}
		}
	}
}

// Check 检查是否符合启动要求
func (r *RemoveHeader) Check() error {

	if r.Header == nil {
		return fmt.Errorf("RemoveHeader 未接收到需要添加的请求头！")
	}

	return nil
}
