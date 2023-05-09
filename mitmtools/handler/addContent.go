package handler

import (
	"fmt"
	"os"

	"github.com/Leviathangk/go-glog/glog"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type AddContentToHeadRule struct {
	Rule
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (a *AddContentToHeadRule) Response(f *proxy.Flow) {

	// 替换响应
	if IsMatch(a.Pattern, f.Request.URL.String()) {
		f.Response.Body = append(a.Content, f.Response.Body...)

		if ShowLog {
			glog.Debugf("AddContentToHeadRule 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddContentToHeadRule) Check() error {

	// 读取文件
	if len(a.Content) == 0 && !a.hasReadFile {
		if a.FilePath != "" {
			content, err := os.ReadFile(a.FilePath)
			if err != nil {
				return fmt.Errorf("AddContentToHeadRule 文件读取出错：" + err.Error())
			} else {
				a.Content = content
			}
		} else {
			return fmt.Errorf("AddContentToHeadRule 未接收到需要替换的内容！")
		}
		a.hasReadFile = true
	}

	return nil
}

type AddContentToTailRule struct {
	Rule
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (a *AddContentToTailRule) Response(f *proxy.Flow) {

	// 替换响应
	if IsMatch(a.Pattern, f.Request.URL.String()) {
		f.Response.Body = append(f.Response.Body, a.Content...)

		if ShowLog {
			glog.Debugf("AddContentToTailRule 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddContentToTailRule) Check() error {

	// 读取文件
	if len(a.Content) == 0 && !a.hasReadFile {
		if a.FilePath != "" {
			content, err := os.ReadFile(a.FilePath)
			if err != nil {
				return fmt.Errorf("AddContentToTailRule 文件读取出错：" + err.Error())
			} else {
				a.Content = content
			}
		} else {
			return fmt.Errorf("AddContentToTailRule 未接收到需要替换的内容！")
		}
		a.hasReadFile = true
	}

	return nil
}
