package handler

import (
	"regexp"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

var ShowLog = false // 是否打印日志

type baseHandler struct {
	proxy.BaseAddon
}

type Addon interface {
	proxy.Addon
	Check() error // 检查输入参数
}

// IsMatch 判断是否匹配指定字符串
func IsMatch(pattern, s string) bool {
	res, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}

	return res
}
