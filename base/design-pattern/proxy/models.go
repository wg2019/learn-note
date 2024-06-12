// Package proxyDef 实现
package proxy

import "sync"

// Qq 腾讯
type Qq struct{}

// News 腾讯新闻
func (q *Qq) News() string {
	return "腾讯新闻信息"
}

// NetEase 网易
type NetEase struct {
}

// News 网易新闻
func (n *NetEase) News() string {
	return "网易新闻信息"
}

// proxyDef 代理定义
type proxyDef struct {
	company    int
	websiteMap map[int]Website
}

// News 代理新闻
func (p *proxyDef) News() string {
	website, ok := p.websiteMap[p.company]
	if ok {
		return website.News()
	}
	return ""
}

// proxy 代理
var proxy *proxyDef

func init() {
	once := new(sync.Once)
	once.Do(func() {
		proxy = new(proxyDef)
		proxy.websiteMap = make(map[int]Website)
		proxy.websiteMap[1] = new(Qq)
		proxy.websiteMap[2] = new(NetEase)
	})
}

// GetWebsite 获取网站
func GetWebsite(company int) Website {
	proxy.company = company
	return proxy
}
