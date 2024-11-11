package resp

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strings"
)

type FindContent struct {
	handler.BaseHandler
	Pattern        string // url 匹配规则
	ContentPattern string // 需要查找的内容
}

func (fin *FindContent) Response(f *proxy.Flow) {
	if handler.IsMatch(fin.Pattern, f.Request.URL.String()) {
		if handler.IsMatch(fin.ContentPattern, string(f.Response.Body)) {
			glog.DLogger.Debugf("FindContent 已查找到：%s -> %s\n", fin.ContentPattern, f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindContent) Check() error {

	if len(fin.ContentPattern) == 0 {
		return fmt.Errorf("FindContent 未接收到需要查找的内容！")
	}

	return nil
}

type FindCookie struct {
	handler.BaseHandler
	Pattern    string   // url 匹配规则
	KeyPattern []string // 需要查找的内容
}

func (fin *FindCookie) Response(f *proxy.Flow) {
	if handler.IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Response.Header {
			if strings.ToLower(key) == "set-cookie" {
				for _, cookie := range value {
					for _, kp := range fin.KeyPattern {
						if handler.IsMatch(kp, cookie) {
							cookie = strings.Split(cookie, ";")[0]
							glog.DLogger.Debugf("FindCookie 已查找到：%s -> %s -> %s\n", kp, cookie, f.Request.URL)
						}
					}
				}
				break
			}
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindCookie) Check() error {

	if len(fin.KeyPattern) == 0 {
		return fmt.Errorf("FindCookie 未接收到需要查找的内容！")
	}

	return nil
}

type FindHeader struct {
	handler.BaseHandler
	Pattern    string   // url 匹配规则
	KeyPattern []string // 需要查找的内容
}

func (fin *FindHeader) Response(f *proxy.Flow) {
	if handler.IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Response.Header {
			for _, kp := range fin.KeyPattern {
				if handler.IsMatch(kp, key) {
					glog.DLogger.Debugf("FindHeader 已查找到：%s -> %s -> %s\n", kp, value, f.Request.URL)
				}
			}
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindHeader) Check() error {

	if len(fin.KeyPattern) == 0 {
		return fmt.Errorf("FindHeader 未接收到需要查找的内容！")
	}

	return nil
}
