package golog

import (
	"testing"
)

func BenchmarkDebug(b *testing.B) {
	l := New()
	l.SetRotateSize(10).SetMaxFileNum(5).Setpath("/tmp/logtest")
	for i := 0; i <= b.N; i++ {
		l.Debug("hello koala : ", i)
	}
}
