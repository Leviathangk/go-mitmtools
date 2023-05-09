package mitmtools

import (
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type MitmConfig struct {
	Debug             int
	Addr              string
	StreamLargeBodies int64 // 当请求或响应体大于此字节时，转为 stream 模式
	SslInsecure       bool
	CaRootPath        string
	Upstream          string
	ShowLog           bool // 是否打印日志
	handlers          []handler.Addon
}

// 添加规则
func (m *MitmConfig) AddHandler(h handler.Addon) {
	m.handlers = append(m.handlers, h)
}

// Start 启动入口
func Start(opts *MitmConfig, handlers ...handler.Addon) error {
	p, err := proxy.NewProxy(&proxy.Options{
		Debug:             opts.Debug,
		Addr:              opts.Addr,
		StreamLargeBodies: opts.StreamLargeBodies,
		SslInsecure:       opts.SslInsecure,
		CaRootPath:        opts.CaRootPath,
		Upstream:          opts.Upstream,
	})
	if err != nil {
		return err
	}

	// 修改配置
	handler.ShowLog = opts.ShowLog

	// 添加解析响应体
	p.AddAddon(new(decodeRule))

	// 添加规则
	for _, h := range handlers {
		err := h.Check()
		if err != nil {
			return err
		}

		p.AddAddon(h)
	}

	// 添加规则
	for _, h := range opts.handlers {
		err := h.Check()
		if err != nil {
			return err
		}

		p.AddAddon(h)
	}

	// 添加响应体重新计算
	p.AddAddon(new(recalculateRule))

	glog.Fatal(p.Start())

	return nil
}
