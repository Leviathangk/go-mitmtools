package mitmtools

import (
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

// Start 启动入口
func Start(opts *Config, handlers ...handler.Addon) error {
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
	glog.DLogger.ShowCaller = false

	// 添加解析响应体
	p.AddAddon(new(handler.DecodeRule))

	// 添加规则
	for _, h := range handlers {
		err = h.Check()
		if err != nil {
			return err
		}

		p.AddAddon(h)
	}

	// 添加规则
	for _, h := range opts.handlers {
		err = h.Check()
		if err != nil {
			return err
		}

		p.AddAddon(h)
	}

	// 添加响应体重新计算
	p.AddAddon(new(handler.RecalculateRule))

	glog.DLogger.Fatal(p.Start())

	return nil
}
