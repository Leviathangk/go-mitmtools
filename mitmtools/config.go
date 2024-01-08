package mitmtools

import (
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler"
	"strconv"
)

const (
	defaultPort              = 8866
	defaultStreamLargeBodies = 1024 * 1024 * 5
	defaultSslInsecure       = true
	defaultShowLog           = true
)

type Config struct {
	Debug             int
	Addr              string
	StreamLargeBodies int64 // 当请求或响应体大于此字节时，转为 stream 模式
	SslInsecure       bool
	CaRootPath        string
	Upstream          string
	ShowLog           bool // 是否打印日志
	handlers          []handler.Addon
}
type SetFunc func(c *Config)

// NewConfig 新建配置
func NewConfig(opt ...SetFunc) *Config {
	config := new(Config)

	for _, o := range opt {
		o(config)
	}

	// 参数检查
	if config.Addr == "" {
		config.Addr = ":8866"
	}
	if config.StreamLargeBodies == 0 {
		config.StreamLargeBodies = 1024 * 1024 * 5
	}

	return config
}

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
	return func(c *Config) {
		c.Upstream = p
	}
}

// SetAddr 设置端口
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

// SetShowLog 日志输出
func SetShowLog(b bool) SetFunc {
	return func(c *Config) {
		c.ShowLog = b
	}
}
