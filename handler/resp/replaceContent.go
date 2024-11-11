package resp

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"
	"strings"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type ReplaceContent struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FindContent string // 查找的需要替换的内容：针对部分内容
	ToContent   string // 替换后的新内容
	Times       int    // 替换次数（0 代表无限）
	timesRecord int    // 记录当前次数
}

func (r *ReplaceContent) Response(f *proxy.Flow) {

	// 替换响应
	if handler.IsMatch(r.Pattern, f.Request.URL.String()) {
		if r.Times != 0 {
			if r.timesRecord >= r.Times {
				glog.DLogger.Warnf("当前替换已达到上限：%d\n", r.Times)
				return
			}
			r.timesRecord += 1
			glog.DLogger.Debugf("当前替换次数：%d-%d\n", r.Times, r.timesRecord)
		}

		f.Response.Body = []byte(strings.ReplaceAll(string(f.Response.Body), r.FindContent, r.ToContent))

		if handler.ShowLog || r.ShowLog {
			glog.DLogger.Debugf("ReplaceContent 已修改响应结果：%s\n", f.Request.URL)
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

// ReplaceContentIfNoCookie 只有当没有 cookie 的时候才替换
type ReplaceContentIfNoCookie struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FindContent string // 查找的需要替换的内容：针对部分内容
	ToContent   string // 替换后的新内容
}

func (r *ReplaceContentIfNoCookie) Response(f *proxy.Flow) {

	// 替换响应
	if handler.IsMatch(r.Pattern, f.Request.URL.String()) {
		if handler.CookieExists(f) {
			glog.DLogger.Warnf("当前存在 cookie 不进行替换：%s\n", f.Request.Header.Get("cookie"))
			return
		}
		f.Response.Body = []byte(strings.ReplaceAll(string(f.Response.Body), r.FindContent, r.ToContent))

		if handler.ShowLog || r.ShowLog {
			glog.DLogger.Debugf("ReplaceContent 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (r *ReplaceContentIfNoCookie) Check() error {

	if r.FindContent == "" {
		return fmt.Errorf("ReplaceContent 未接收到需要替换的内容！")
	}

	return nil
}
