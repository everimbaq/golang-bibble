package system_design

import (
	"sync"
	"sync/atomic"
)

type RoutinePool struct {
	Tasks      chan func()
	Count      int
	ifStopWork int32
	wg         sync.WaitGroup
}

func (r *RoutinePool) Start() {
	for i := 0; i < r.Count; i++ {
		r.wg.Add(1)
		go func() {
			for {
				task, ok := <-r.Tasks
				if !ok {
					r.wg.Done()
					return
				}
				task()

			}
		}()
	}
}

func (r *RoutinePool) Stop() {
	close(r.Tasks)
	r.wg.Wait()
}

func (r *RoutinePool) SafeStop() {

	if atomic.LoadInt32(&r.ifStopWork) != 1 {
		atomic.StoreInt32(&r.ifStopWork, 1)
	}
	r.Stop()
}

func (r *RoutinePool) AddTask(task func()) bool {
	if atomic.LoadInt32(&r.ifStopWork) == 1 {
		return false
	}
	r.Tasks <- task
	return true
}
