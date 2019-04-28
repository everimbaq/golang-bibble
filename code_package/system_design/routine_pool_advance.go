package system_design

import (
	"github.com/everimbaq/golang-bibble/code_package/datastructs_and_algorithms"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type DoFunc func(interface{})
type PoolAdv struct {
	capacity  int32
	doFunc    DoFunc
	mtx       sync.Mutex
	count     int32
	idle      *datastructs_and_algorithms.Stack
	closeChan chan struct{}
}

type IdleWorkers struct {
	w    *Worker
	next *IdleWorkers
}

type Worker struct {
	TaskArgs chan interface{}
	TaskFunc DoFunc
	pool     *PoolAdv
	idleTime int64
}

func NewPool(cap int32, doF DoFunc) *PoolAdv {
	p := &PoolAdv{
		capacity: cap,
		idle:     datastructs_and_algorithms.NewStack(),
		doFunc:   doF,
	}
	go p.CloseIdle()
	return p
}

func (p *PoolAdv) Add(args interface{}) {
	woker := p.getWorker()
	woker.TaskArgs <- args
}

func (p *PoolAdv) getWorker() *Worker {
	p.mtx.Lock()
	for {
		if w := p.PopIdle(); w != nil {
			p.mtx.Unlock()
			return w
		}
		if atomic.LoadInt32(&p.count) <= p.capacity {
			w := p.newWorker()
			p.mtx.Unlock()
			return w
		}
	}
}

func (p *PoolAdv) newWorker() *Worker {
	w := &Worker{
		TaskArgs: make(chan interface{}, 1),
		TaskFunc: p.doFunc,
		pool:     p,
	}
	w.Loop()
	atomic.AddInt32(&p.count, 1)
	return w
}

func (p *PoolAdv) CloseIdle() {
	heartbeat := time.NewTicker(time.Second * 10)
	now := time.Now().Unix()
	for range heartbeat.C {
		p.mtx.Lock()
		var expireIndex = -1
		for i := 0; i < p.idle.GetSize(); i++ {
			w := p.idle.Read(i).(*Worker)
			if now-w.idleTime > 10 {
				expireIndex++
			} else {
				break
			}
		}
		if expireIndex != -1 {
			p.idle.SubN(expireIndex)
		}
		p.mtx.Unlock()
		log.Println("clean expire, remains:", p.idle.GetSize())
	}
}

func (w *Worker) Loop() {
	go func() {
		for arg := range w.TaskArgs {
			if arg == nil {
				atomic.AddInt32(&w.pool.count, -1)
				return
			}
			w.TaskFunc(arg)
			w.idleTime = time.Now().Unix()
			w.pool.idle.Add(w)
		}
	}()
}

func (p *PoolAdv) AddIdle(w *Worker) {
	p.idle.Add(w)
}

func (p *PoolAdv) PopIdle() *Worker {
	idleW := p.idle.Pop()
	if idleW == nil {
		return nil
	}
	return idleW.(*Worker)
}
