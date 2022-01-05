package tools

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFind(b *testing.B) {
	b.StopTimer()
	d := Instance
	rand.Seed(time.Now().UnixNano())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = d.Search("上海")
	}
}

func BenchmarkFindParallel(b *testing.B) {
	b.StopTimer()
	d := Instance
	rand.Seed(time.Now().UnixNano())
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = d.Search("上海")
		}
	})
}
