package req

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/mitmtools/handler"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strings"
)

type ChangeHeader struct {
	handler.BaseHandler
	Pattern string              // url 匹配规则
	Header  map[string][]string // 需要替换的头
}

func (r *ChangeHeader) Requestheaders(f *proxy.Flow) {
	if handler.IsMatch(r.Pattern, f.Request.URL.String()) {
		for key, value := range r.Header {
			if _, ok := f.Request.Header[key]; ok {
				f.Request.Header[key] = value

				if handler.ShowLog {
					glog.DLogger.Debugf("ChangeHeader 正在替换指定请求头：%s -> %v\n", key, value)
				}
			}
		}
	}
}

// Check 检查是否符合启动要求
func (r *ChangeHeader) Check() error {

	if r.Header == nil {
		return fmt.Errorf("ChangeHeader 未接收到需要替换的请求头！")
	}

	return nil
}

type ChangeCookie struct {
	handler.BaseHandler
	Pattern string            // url 匹配规则
	Cookie  map[string]string // 需要查找的内容
}

func (fin *ChangeCookie) Requestheaders(f *proxy.Flow) {
	if handler.IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Request.Header {
			if strings.ToLower(key) == "cookie" {
				for keyIndex, cookie := range value {
					for newK, newV := range fin.Cookie {
						if handler.IsMatch(newK, cookie) {
							newCookie := newK + "=" + newV
							f.Request.Header[key][keyIndex] = newCookie

							if handler.ShowLog {
								glog.DLogger.Debugf("ChangeCookie 已查找到：%s -> %s -> %s\n", newK, newCookie, f.Request.URL)
							}
						}
					}
				}
				break
			}
		}
	}
}

// Check 检查是否符合启动要求
func (fin *ChangeCookie) Check() error {

	if len(fin.Cookie) == 0 {
		return fmt.Errorf("ChangeCookie 未接收到需要查找的内容！")
	}

	return nil
}
