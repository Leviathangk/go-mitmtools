package handler

import (
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"regexp"
)

// IsMatch 判断是否匹配指定字符串
func IsMatch(pattern, s string) bool {
	res, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}

	return res
}

// CookieExists 判读请求头有无 cookie
func CookieExists(f *proxy.Flow) bool {
	return f.Request.Header.Get("cookie") != ""
}
