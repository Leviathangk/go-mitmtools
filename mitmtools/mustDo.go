package mitmtools

import (
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler"
	"strconv"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type decodeRule struct {
	handler.Rule
}

// Response 响应体解码：避免重复解码
func (r *decodeRule) Response(f *proxy.Flow) {
	f.Response.ReplaceToDecodedBody()
}

type recalculateRule struct {
	handler.Rule
}

// Response 重新计算 Content-Length 长度
func (r *recalculateRule) Response(f *proxy.Flow) {
	f.Response.Header.Set("Content-Length", strconv.Itoa(len(f.Response.Body)))
}
