package handler

import (
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

var ShowLog = false

type BaseHandler struct {
	proxy.BaseAddon
	ShowLog bool // 单独配置
}

// Addon 基础的处理器
type Addon interface {
	proxy.Addon
	Check() error // 检查输入参数
}
