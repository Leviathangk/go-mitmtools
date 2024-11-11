package resp

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"
	"os"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type AddContentToHead struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (a *AddContentToHead) Response(f *proxy.Flow) {

	// 替换响应
	if handler.IsMatch(a.Pattern, f.Request.URL.String()) {
		f.Response.Body = append(a.Content, f.Response.Body...)

		if handler.ShowLog || a.ShowLog {
			glog.DLogger.Debugf("AddContentToHead 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddContentToHead) Check() error {

	// 读取文件
	if len(a.Content) == 0 && !a.hasReadFile {
		if a.FilePath != "" {
			content, err := os.ReadFile(a.FilePath)
			if err != nil {
				return fmt.Errorf("AddContentToHead 文件读取出错：" + err.Error())
			} else {
				a.Content = content
			}
		} else {
			return fmt.Errorf("AddContentToHead 未接收到需要替换的内容！")
		}
		a.hasReadFile = true
	}

	return nil
}

type AddContentToTail struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (a *AddContentToTail) Response(f *proxy.Flow) {

	// 替换响应
	if handler.IsMatch(a.Pattern, f.Request.URL.String()) {
		f.Response.Body = append(f.Response.Body, a.Content...)

		if handler.ShowLog || a.ShowLog {
			glog.DLogger.Debugf("AddContentToTail 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddContentToTail) Check() error {

	// 读取文件
	if len(a.Content) == 0 && !a.hasReadFile {
		if a.FilePath != "" {
			content, err := os.ReadFile(a.FilePath)
			if err != nil {
				return fmt.Errorf("AddContentToTail 文件读取出错：" + err.Error())
			} else {
				a.Content = content
			}
		} else {
			return fmt.Errorf("AddContentToTail 未接收到需要替换的内容！")
		}
		a.hasReadFile = true
	}

	return nil
}
