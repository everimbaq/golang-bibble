package system_design

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkRoutinePoolAdvance(b *testing.B) {
	var num int32
	task := func(i interface{}) {
		atomic.AddInt32(&num, 1)
	}

	p := NewPool(1000, task)
	for i := 0; i < b.N; i++ {
		p.Add(i)
	}
	b.Log("num:", num)
}

func BenchmarkTime(b *testing.B) {
	var t int64
	for i := 0; i < b.N; i++ {
		t = time.Now().UnixNano()
	}

	b.Log(t)
}
