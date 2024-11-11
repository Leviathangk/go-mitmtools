package mitmtools

import (
	"github.com/Leviathangk/go-mitmtools/handler"
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
		config.Addr = ":" + strconv.Itoa(defaultPort)
	}
	if config.StreamLargeBodies == 0 {
		config.StreamLargeBodies = defaultStreamLargeBodies
	}

	return config
}
