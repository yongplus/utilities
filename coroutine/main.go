package coroutine

import (
	"context"
	"sync"
)

type Coroutine struct{
	waiter *sync.WaitGroup
	num int
	ctx context.Context
	cancel context.CancelFunc
	Chans chan interface{}
}

/**
 * @param.num:the number of coroutine
 * @param.num:the size of chans buffer
 */

func New(num int,chanSize int) *Coroutine{
	waiter := &sync.WaitGroup{}
	waiter.Add(num)
	ctx,cancel := context.WithCancel(context.Background())
	coroutine := &Coroutine{
		num:num,
		waiter: waiter,
		Chans: make(chan interface{},chanSize),
		ctx: ctx,
		cancel: cancel,

	}

	return coroutine
}

/**
 * Write data to the chans
 */
func (m *Coroutine) Push(data interface{}){
	m.Chans <- data
}

func (m *Coroutine) SetWorker(worker func(interface{})){
	for i:=0;i<m.num;i++{
		go (func(){
			defer m.waiter.Done()
			for {
				select {
					case val:= <-	m.Chans:
						worker(val)
					case <-m.ctx.Done():
						return
				}
			}
		})()
	}
}
/**
 * Waiting for all the coroutine to finish.
 */
func (m *Coroutine) Wait(){
	m.cancel()
	m.waiter.Wait()
	close(m.Chans)
}




