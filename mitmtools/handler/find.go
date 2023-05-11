package handler

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler/common"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strings"
)

type FindContent struct {
	common.baseHandler
	Pattern        string // url 匹配规则
	ContentPattern string // 需要查找的内容
}

func (fin *FindContent) Response(f *proxy.Flow) {
	if IsMatch(fin.Pattern, f.Request.URL.String()) {
		if IsMatch(fin.ContentPattern, string(f.Response.Body)) {
			glog.Debugf("FindContent 已查找到：%s -> %s\n", fin.ContentPattern, f.Request.URL)
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
	common.baseHandler
	Pattern    string   // url 匹配规则
	KeyPattern []string // 需要查找的内容
}

func (fin *FindCookie) Response(f *proxy.Flow) {
	if IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Response.Header {
			if strings.ToLower(key) == "set-cookie" {
				for _, cookie := range value {
					for _, kp := range fin.KeyPattern {
						if IsMatch(kp, cookie) {
							cookie = strings.Split(cookie, ";")[0]
							glog.Debugf("FindCookie 已查找到：%s -> %s -> %s\n", kp, cookie, f.Request.URL)
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

type FindHeaderRule struct {
	common.baseHandler
	Pattern    string   // url 匹配规则
	KeyPattern []string // 需要查找的内容
}

func (fin *FindHeaderRule) Response(f *proxy.Flow) {
	if IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Response.Header {
			for _, kp := range fin.KeyPattern {
				if IsMatch(kp, key) {
					glog.Debugf("FindHeaderRule 已查找到：%s -> %s -> %s\n", kp, value, f.Request.URL)
				}
			}
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindHeaderRule) Check() error {

	if len(fin.KeyPattern) == 0 {
		return fmt.Errorf("FindHeaderRule 未接收到需要查找的内容！")
	}

	return nil
}
