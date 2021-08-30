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
	chans chan interface{}
}

// @param.num:the number of coroutine
// @param.num:the size of chans buffer
func New(num int,chanSize int) *Coroutine{
	waiter := &sync.WaitGroup{}
	waiter.Add(num)
	coroutine := &Coroutine{
		num:num,
		waiter: waiter,
		chans: make(chan interface{},chanSize),
	}

	return coroutine
}

// Write data to the chans
func (m *Coroutine) Push(data interface{}){
	if data==nil {
		return
	}
	m.chans <- data
}

func (m *Coroutine) SetWorker(worker func(interface{})){
	for i:=0;i<m.num;i++{
		go (func(){
			defer m.waiter.Done()
			for {

				val:= <-m.chans
				if val==nil {
					return
				}
				worker(val)
			}
		})()
	}
}

// Waiting for all the coroutine to finish.
func (m *Coroutine) Wait(){
	for i:=0;i<m.num;i++{
		m.chans <- nil
	}
	m.waiter.Wait()
	close(m.chans)
}




