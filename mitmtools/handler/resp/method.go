package resp

import "github.com/lqqyt2423/go-mitmproxy/proxy"

// cookieExists 判读请求头有无 cookie
func cookieExists(f *proxy.Flow) bool {
	return f.Request.Header.Get("cookie") != ""
}
