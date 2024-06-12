// Package decorator 测试
package decorator

import "testing"

func TestFAW_Checklist(t *testing.T) {
	l := NewLiberation()
	t.Logf("原厂清单：%s", l.Checklist())

	for i := 1; i < 10; i++ {
		l = NewFAW(l, i)
	}
	for _, value := range l.Checklist() {
		t.Logf("改装情况：%s", value)
	}
}
