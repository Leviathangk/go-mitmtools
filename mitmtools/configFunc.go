package mitmtools

import (
	"github.com/Leviathangk/go-mitmtools/handler"
	"strconv"
	"strings"
)

// AddHandler 添加规则
func (c *Config) AddHandler(h handler.Addon) {
	c.handlers = append(c.handlers, h)
}

// AddHandler 添加规则
func AddHandler(h handler.Addon) SetFunc {
	return func(c *Config) {
		c.handlers = append(c.handlers, h)
	}
}

// SetProxy 设置代理
func SetProxy(p string) SetFunc {
	if p != "" && !strings.HasPrefix(p, "http") {
		p = "http://" + p
	}

	return func(c *Config) {
		c.Upstream = p
	}
}

// SetPort 设置端口
func SetPort(p int) SetFunc {
	return func(c *Config) {
		c.Addr = "127.0.0.1:" + strconv.Itoa(p)
	}
}

// SetAddr 设置ip、端口
func SetAddr(ip string, p int) SetFunc {
	return func(c *Config) {
		c.Addr = ip + ":" + strconv.Itoa(p)
	}
}

// SetStreamLargeBodies 当请求或响应体大于此字节时，转为 stream 模式
func SetStreamLargeBodies(p int) SetFunc {
	return func(c *Config) {
		c.Addr = ":" + strconv.Itoa(p)
	}
}

// SetSslInsecure ssl 校验
func SetSslInsecure(b bool) SetFunc {
	return func(c *Config) {
		c.SslInsecure = b
	}
}

// SetShowLog 日志输出，一旦配置，将会忽略每个 handler 的默认输出配置
func SetShowLog(b bool) SetFunc {
	return func(c *Config) {
		c.ShowLog = b
	}
}

// SetCaRootPath 设置证书路径，文件夹路径，非文件
func SetCaRootPath(p string) SetFunc {
	return func(c *Config) {
		c.CaRootPath = p
	}
}
