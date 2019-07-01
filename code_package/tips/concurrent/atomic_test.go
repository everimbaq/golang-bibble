package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var countMtx sync.Mutex
var countSig int32
var count int

func atomAdd() {
	for {
		if atomic.CompareAndSwapInt32(&countSig, 0, 1) {
			count++
			atomic.StoreInt32(&countSig, 0)
			break
		}
	}

}

func mtxAdd() {
	countMtx.Lock()
	count++
	countMtx.Unlock()
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			atomAdd()
		}()
	}
	b.Log("Atomic ", count)
}

func BenchmarkMtx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			mtxAdd()
		}()
	}

	b.Log("mtx ", count)
}

// 验证atomic作用
func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	var i int64 = 0
	var j int64 = 0
	for r := 0; r < 1000000; r++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&i, 1)
			j++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(i, j)
}
