package system_design

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestRoutinePool(t *testing.T) {
	rp := &RoutinePool{
		Tasks: make(chan func(), 10000),
		Count: 10,
	}
	rp.Start()
	t.Log("i", exec(rp))
}

func exec(rp *RoutinePool) int32 {
	var i int32
	task := func() {
		atomic.AddInt32(&i, 1)
	}

	for j := 0; j < 1000000; j++ {
		rp.AddTask(task)
	}
	rp.SafeStop()
	return i
}

func execNoPool() int32 {
	var i int32
	var wg sync.WaitGroup
	task := func() {
		atomic.AddInt32(&i, 1)
	}

	for j := 0; j < 1000000; j++ {
		wg.Add(1)
		go func() {
			task()
			wg.Done()
		}()
	}
	wg.Wait()
	return i
}

func BenchmarkRoutinePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rp := &RoutinePool{
			Tasks: make(chan func(), 10000),
			Count: 10,
		}
		rp.Start()
		b.StartTimer()
		b.Log(exec(rp))
		b.StopTimer()
	}
}

func BenchmarkNoRoutinePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(execNoPool())
	}
}
