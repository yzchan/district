package district

import (
	"math/rand"
	"testing"
	"time"
)

func TestDistrict_Search(t *testing.T) {
	d := Instance
	t.Log("Start testing Search func")
	adCode := 510000
	t.Logf("arguments: 四川省[%d]", adCode)
	primary, secondary := d.ShortNames(adCode)
	t.Logf("primary, secondary := d.ShortNames(%d)", adCode)
	t.Logf("primary   = \"%s\"", primary)
	t.Logf("secondary = \"%s\"", secondary)
	if primary != "川" || secondary != "蜀" {
		t.Fatal("error")
	}
	t.Log("test ok")
}

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
