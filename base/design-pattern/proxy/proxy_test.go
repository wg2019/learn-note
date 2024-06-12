// Package proxy 测试
package proxy

import "testing"

// TestGetWebsite 测试
func TestGetWebsite(t *testing.T) {
	website := GetWebsite(1)
	news := website.News()
	t.Logf("news: %s", news)
	website2 := GetWebsite(2)
	news2 := website2.News()
	t.Logf("news: %s", news2)
}
