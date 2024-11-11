package resp

import (
	"fmt"
	"github.com/Leviathangk/go-glog/glog"
	"github.com/Leviathangk/go-mitmtools/handler"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
	"strings"
)

type ChangeHeader struct {
	handler.BaseHandler
	Pattern string              // url 匹配规则
	Header  map[string][]string // 需要替换的头
}

func (r *ChangeHeader) Responseheaders(f *proxy.Flow) {
	if handler.IsMatch(r.Pattern, f.Request.URL.String()) {
		for key, value := range r.Header {
			if _, ok := f.Response.Header[key]; ok {
				f.Response.Header[key] = value

				if handler.ShowLog || r.ShowLog {
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

func (fin *ChangeCookie) Response(f *proxy.Flow) {
	if handler.IsMatch(fin.Pattern, f.Request.URL.String()) {
		for key, value := range f.Response.Header {
			if strings.ToLower(key) == "set-cookie" {
				for keyIndex, cookie := range value {
					for newK, newV := range fin.Cookie {
						if handler.IsMatch(newK, cookie) {
							cookieSlice := strings.Split(cookie, ";")

							cookie = strings.ReplaceAll(cookie, cookieSlice[0], strings.Split(cookieSlice[0], "=")[0]+"="+newV)

							f.Response.Header[key][keyIndex] = cookie

							if handler.ShowLog || fin.ShowLog {
								glog.DLogger.Debugf("ChangeCookie 已查找到：%s -> %s -> %s\n", newK, cookie, f.Request.URL)
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
