package handler

import (
	"fmt"

	"github.com/Leviathangk/go-glog/glog"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type AddResponseHeader struct {
	Rule
	Pattern string            // url 匹配规则
	Header  map[string]string // 需要添加的 map
}

func (a *AddResponseHeader) Responseheaders(f *proxy.Flow) {
	if IsMatch(a.Pattern, f.Request.URL.String()) {
		for key, value := range a.Header {
			f.Response.Header.Add(key, value)

			if ShowLog {
				glog.Debugf("AddResponseHeader 正在添加指定请求头：%s -> %s\n", key, value)
			}
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddResponseHeader) Check() error {

	if a.Header == nil {
		return fmt.Errorf("AddResponseHeader 未接收到需要添加的请求头！")
	}

	return nil
}

type RemoveResponseHeader struct {
	Rule
	Pattern string   // url 匹配规则
	Header  []string // 需要移除的 key
}

func (r *RemoveResponseHeader) Responseheaders(f *proxy.Flow) {
	if IsMatch(r.Pattern, f.Request.URL.String()) {
		for _, key := range r.Header {
			if _, ok := f.Response.Header[key]; ok {
				delete(f.Response.Header, key)

				if ShowLog {
					glog.Debugf("RemoveResponseHeader 正在移除指定请求头：%s\n", key)
				}
			}
		}
	}
}

// Check 检查是否符合启动要求
func (r *RemoveResponseHeader) Check() error {

	if r.Header == nil {
		return fmt.Errorf("RemoveResponseHeader 未接收到需要添加的请求头！")
	}

	return nil
}
