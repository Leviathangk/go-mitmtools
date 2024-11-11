package resp

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Leviathangk/go-mitmtools/handler"

	"github.com/Leviathangk/go-glog/glog"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type AddScriptToHead struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (a *AddScriptToHead) Response(f *proxy.Flow) {
	// 判断是否是 html
	if !IsHtml(f) {
		return
	}

	// 替换响应
	if handler.IsMatch(a.Pattern, f.Request.URL.String()) {
		// 按照标签查找找到就退出
		scripts := []string{"<head>", "<body>"}
		hasReplace := false
		for _, script := range scripts {

			// 查找 head 标签的索引
			pattern := regexp.MustCompile(script)
			index := pattern.FindIndex(f.Response.Body)
			if len(index) == 0 {
				continue
			}

			// 添加 script 标签
			content := append([]byte("<script>"), append([]byte(a.Content), []byte("</script>")...)...)

			// 插入标签
			f.Response.Body = append(f.Response.Body[:index[1]], append(content, f.Response.Body[index[1]:]...)...)

			hasReplace = true

			if handler.ShowLog || a.ShowLog {
				glog.DLogger.Debugf("AddScriptToHead 已修改响应结果：%s\n", f.Request.URL)
			}

			break
		}

		if !hasReplace {
			glog.DLogger.Warnf("AddScriptToHead 未找到标签：%v，未替换结果\n", scripts)
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddScriptToHead) Check() error {

	// 读取文件
	if len(a.Content) == 0 && !a.hasReadFile {
		if a.FilePath != "" {
			content, err := os.ReadFile(a.FilePath)
			if err != nil {
				return fmt.Errorf("AddScriptToHead 文件读取出错：" + err.Error())
			} else {
				a.Content = content
			}
		} else {
			return fmt.Errorf("AddScriptToHead 未接收到需要替换的内容！")
		}
		a.hasReadFile = true
	}

	return nil
}

type AddScriptToTail struct {
	handler.BaseHandler
	Pattern     string // url 匹配规则
	FilePath    string // 需要替换的文件路径（二选一）
	Content     []byte // 需要替换的内容（二选一）
	hasReadFile bool   // 是否已经读取过文件
}

func (a *AddScriptToTail) Response(f *proxy.Flow) {
	// 判断是否是 html
	if !IsHtml(f) {
		return
	}

	// 替换响应
	if handler.IsMatch(a.Pattern, f.Request.URL.String()) {

		// 按照标签查找找到就退出
		scripts := []string{"</body>", "</head>"}
		hasReplace := false
		for _, script := range scripts {

			// 查找 head 标签的索引
			pattern := regexp.MustCompile(script)
			index := pattern.FindIndex(f.Response.Body)
			if len(index) == 0 {
				continue
			}

			// 添加 script 标签
			content := append([]byte("<script>"), append([]byte(a.Content), []byte("</script>")...)...)

			// 插入标签
			f.Response.Body = append(f.Response.Body[:index[0]], append(content, f.Response.Body[index[0]:]...)...)

			hasReplace = true

			if handler.ShowLog || a.ShowLog {
				glog.DLogger.Debugf("AddScriptToTail 已修改响应结果：%s\n", f.Request.URL)
			}

			break
		}

		if !hasReplace {
			glog.DLogger.Warnf("AddScriptToTail 未找到标签：%v，未替换结果\n", scripts)
		}
	}
}

// Check 检查是否符合启动要求
func (a *AddScriptToTail) Check() error {

	// 读取文件
	if len(a.Content) == 0 && !a.hasReadFile {
		if a.FilePath != "" {
			content, err := os.ReadFile(a.FilePath)
			if err != nil {
				return fmt.Errorf("AddScriptToTail 文件读取出错：" + err.Error())
			} else {
				a.Content = content
			}
		} else {
			return fmt.Errorf("AddScriptToTail 未接收到需要替换的内容！")
		}
		a.hasReadFile = true
	}

	return nil
}

// IsHtml 判断是否是 html 页面
func IsHtml(f *proxy.Flow) bool {
	// 获取响应类型
	contentType := f.Response.Header.Values("Content-Type")

	// 判断是否存在
	for _, value := range contentType {
		if strings.Contains(value, "text/html") {
			return true
		}
	}

	return false
}
