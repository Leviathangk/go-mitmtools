package handler

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"os"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type ReplaceFileRule struct {
	Rule
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (r *ReplaceFileRule) Response(f *proxy.Flow) {

	// 替换响应
	if IsMatch(r.Pattern, f.Request.URL.String()) {
		f.Response.Body = r.Content

		if ShowLog {
			glog.Debugf("ReplaceFileRule 已修改响应结果：%s\n", f.Request.URL)
		}
	}
}

// Check 检查是否符合启动要求
func (r *ReplaceFileRule) Check() error {

	// 读取文件
	if len(r.Content) == 0 && !r.hasReadFile {
		if r.FilePath != "" {
			content, err := os.ReadFile(r.FilePath)
			if err != nil {
				return fmt.Errorf("ReplaceFileRule 文件读取出错：" + err.Error())
			} else {
				r.Content = content
			}
		} else {
			return fmt.Errorf("ReplaceFileRule 未接收到需要替换的内容！")
		}
		r.hasReadFile = true
	}

	return nil
}
