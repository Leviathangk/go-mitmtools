package resp

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"
	"os"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type ReplaceFile struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
	Times       int    // 替换次数（0 代表无限）
	timesRecord int    // 记录当前次数
}

func (r *ReplaceFile) Response(f *proxy.Flow) {

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

		f.Response.Body = r.Content

		if handler.ShowLog || r.ShowLog {
			glog.DLogger.Debugf("ReplaceFile 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (r *ReplaceFile) Check() error {

	// 读取文件
	if len(r.Content) == 0 && !r.hasReadFile {
		if r.FilePath != "" {
			content, err := os.ReadFile(r.FilePath)
			if err != nil {
				return fmt.Errorf("ReplaceFile 文件读取出错：" + err.Error())
			} else {
				r.Content = content
			}
		} else {
			return fmt.Errorf("ReplaceFile 未接收到需要替换的内容！")
		}
		r.hasReadFile = true
	}

	return nil
}

// ReplaceFileIfNoCookie 只有当没有 cookie 的时候才替换
type ReplaceFileIfNoCookie struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (r *ReplaceFileIfNoCookie) Response(f *proxy.Flow) {

	// 替换响应
	if handler.IsMatch(r.Pattern, f.Request.URL.String()) {
		if handler.CookieExists(f) {
			glog.DLogger.Warnf("当前存在 cookie 不进行替换：%s\n", f.Request.Header.Get("cookie"))
			return
		}
		f.Response.Body = r.Content

		if handler.ShowLog || r.ShowLog {
			glog.DLogger.Debugf("ReplaceFile 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (r *ReplaceFileIfNoCookie) Check() error {

	// 读取文件
	if len(r.Content) == 0 && !r.hasReadFile {
		if r.FilePath != "" {
			content, err := os.ReadFile(r.FilePath)
			if err != nil {
				return fmt.Errorf("ReplaceFile 文件读取出错：" + err.Error())
			} else {
				r.Content = content
			}
		} else {
			return fmt.Errorf("ReplaceFile 未接收到需要替换的内容！")
		}
		r.hasReadFile = true
	}

	return nil
}
