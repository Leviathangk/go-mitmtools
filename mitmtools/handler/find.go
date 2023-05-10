package handler

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strings"
)

type FindContentRule struct {
	Rule
	Pattern        string // url 匹配规则
	ContentPattern string // 需要查找的内容
}

func (fin *FindContentRule) Response(f *proxy.Flow) {
	if IsMatch(fin.Pattern, f.Request.URL.String()) {
		if IsMatch(fin.ContentPattern, string(f.Response.Body)) {
			glog.Debugf("FindContentRule 已查找到：%s -> %s\n", fin.ContentPattern, f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindContentRule) Check() error {

	if len(fin.ContentPattern) == 0 {
		return fmt.Errorf("FindContentRule 未接收到需要查找的内容！")
	}

	return nil
}

type FindCookieRule struct {
	Rule
	Pattern    string   // url 匹配规则
	KeyPattern []string // 需要查找的内容
}

func (fin *FindCookieRule) Response(f *proxy.Flow) {
	if IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Response.Header {
			if strings.ToLower(key) == "set-cookie" {
				for _, cookie := range value {
					for _, kp := range fin.KeyPattern {
						if IsMatch(kp, cookie) {
							cookie = strings.Split(cookie, ";")[0]
							glog.Debugf("FindCookieRule 已查找到：%s -> %s -> %s\n", kp, cookie, f.Request.URL)
						}
					}
				}
				break
			}
		}
	}
}

// Check 检查是否符合启动要求
func (fin *FindCookieRule) Check() error {

	if len(fin.KeyPattern) == 0 {
		return fmt.Errorf("FindCookieRule 未接收到需要查找的内容！")
	}

	return nil
}

type FindHeaderRule struct {
	Rule
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
