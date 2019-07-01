package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSched(t *testing.T) {
	runtime.GOMAXPROCS(4)
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)
		if i == 1 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}

func BenchmarkPrint(b *testing.B) {
	var s string
	for i := 0; i < 10000; i++ {
		s = fmt.Sprintf("%d", i)
	}
	fmt.Println(s)
}
