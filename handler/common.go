package handler

import (
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strconv"
)

type DecodeRule struct {
	proxy.BaseAddon
}

// Response 响应体解码：避免重复解码
func (r *DecodeRule) Response(f *proxy.Flow) {
	f.Response.ReplaceToDecodedBody()
}

type RecalculateRule struct {
	proxy.BaseAddon
}

// Response 重新计算 Content-Length 长度
func (r *RecalculateRule) Response(f *proxy.Flow) {
	f.Response.Header.Set("Content-Length", strconv.Itoa(len(f.Response.Body)))
}
