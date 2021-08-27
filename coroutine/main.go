package coroutine

import (
	"context"
	"sync"
)

type Coroutine struct{
	waiter *sync.WaitGroup
	num int
	chans chan interface{}
	ctx context.Context
	cancel context.CancelFunc
}

func New(num int,chans chan interface{}) *Coroutine{
	waiter := &sync.WaitGroup{}
	waiter.Add(num)
	ctx,cancel := context.WithCancel(context.Background())
	coroutine := &Coroutine{
		num:num,
		waiter: waiter,
		chans: chans,
		ctx: ctx,
		cancel: cancel,
	}

	return coroutine
}

func (m *Coroutine) AddWorker(worker func(interface{})){


	for i:=0;i<m.num;i++{
		go (func(){
			for {
				select {
					case val:= <-	m.chans:
						worker(val)
					case <-m.ctx.Done():
						m.waiter.Done()
						return
				}
			}
		})()
	}

}

func (m *Coroutine) Wait(){
	m.cancel()
	m.waiter.Wait()
}




