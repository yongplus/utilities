package coroutine

import (
	"context"
	"sync"
)

type Coroutine struct {
	waiter    *sync.WaitGroup
	num       int
	ctx       context.Context
	cancel    context.CancelFunc
	chans     chan interface{}
	recvChans chan interface{}
}

// @param.num:the number of coroutine
// @param.num:the size of chans buffer
func New(num int, chanSize int) *Coroutine {
	waiter := &sync.WaitGroup{}
	waiter.Add(num)
	coroutine := &Coroutine{
		num:    num,
		waiter: waiter,
		chans:  make(chan interface{}, chanSize),
	}

	return coroutine
}

// Write data to the chans
func (m *Coroutine) Push(data interface{}) {
	if data == nil {
		return
	}
	m.chans <- data
}

//Set a worker with a return result
func (m *Coroutine) SetWorker(worker func(interface{}) interface{}) {
	for i := 0; i < m.num; i++ {
		go (func() {
			defer m.waiter.Done()
			for {

				val := <-m.chans
				if val == nil {
					return
				}
				result := worker(val)
				if result != nil && m.recvChans != nil {
					m.recvChans <- result
				}
			}
		})()
	}
}

//Set a worker without a return result
func (m *Coroutine) SetWorker2(worker func(interface{})) {
	for i := 0; i < m.num; i++ {
		go (func() {
			defer m.waiter.Done()
			for {

				val := <-m.chans
				if val == nil {
					return
				}
				worker(val)
			}
		})()
	}
}

//Set the result receive function
func (m *Coroutine) SetListener(recv func(interface{})) {
	if recv == nil {
		return
	}
	m.recvChans = make(chan interface{}, 0)
	go (func() {
		for {
			val := <-m.recvChans
			if val == nil {
				return
			}
			recv(val)
		}
	})()
}

// Waiting for all the workers and listener to finish.
func (m *Coroutine) Wait() {
	for i := 0; i < m.num; i++ {
		m.chans <- nil
	}

	m.waiter.Wait()
	//time.Sleep(time.Millisecond*50)
	if m.recvChans != nil {
		m.recvChans <- nil
	}

	close(m.chans)
}
