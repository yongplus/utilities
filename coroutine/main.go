package coroutine

import (
	"context"
	"log"
	"runtime"
	"sync"
)

type Coroutine struct {
	waiter        *sync.WaitGroup
	num           int
	ctx           context.Context
	cancel        context.CancelFunc
	chans         chan interface{}
	recvChans     chan interface{}
	havesetworker bool
}

type Chans struct {
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
	if !m.havesetworker {
		panic("Please make sure set worker by calling SetWorker* methods before pushing data~!")
	}
	m.chans <- data
}

//Set a worker with a return result
func (m *Coroutine) SetWorker(worker func(interface{}) interface{}) {
	m.havesetworker = true
	for i := 0; i < m.num; i++ {
		go (func() {
			// 延迟处理的函数
			defer m.waiter.Done()
			for {

				val := <-m.chans
				if val == nil {
					return
				}
				result := func() interface{} {
					defer m._recovery()
					return worker(val)
				}()
				if result != nil && m.recvChans != nil {
					m.recvChans <- result
				}
			}
		})()
	}
}

//Set a worker without a return result
func (m *Coroutine) SetWorker2(worker func(interface{})) {
	m.havesetworker = true
	for i := 0; i < m.num; i++ {
		go (func() {
			defer m.waiter.Done()
			for {

				val := <-m.chans
				if val == nil {
					return
				}
				func() {
					defer m._recovery()
					worker(val)
				}()

			}
		})()
	}
}

//Set the result receive function
func (m *Coroutine) SetListener(recv func(interface{})) {
	if recv == nil {
		return
	}
	m._resetRecvChans()
	go (func() {
		for {
			val := <-m.recvChans
			if val == nil {
				log.Println("null")
				return
			}
			recv(val)
		}
	})()
}

//Set the result receive function
func (m *Coroutine) RecvChans() chan interface{} {
	m._resetRecvChans()
	return m.recvChans
}

func (m *Coroutine) _resetRecvChans() {
	if m.recvChans != nil { // if the recvChans has been set than release it and reset
		tmpChans := m.recvChans
		m.recvChans <- nil
		m.recvChans = nil
		close(tmpChans)
	}
	m.recvChans = make(chan interface{}, 0)
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
		close(m.recvChans)
	}

	close(m.chans)

}

func (m *Coroutine) _recovery() {
	// 发生宕机时，获取panic传递的上下文并打印
	err := recover()
	switch err.(type) {
	case runtime.Error: // 运行时错误
		log.Println("runtime error:", err)
	default: // 非运行时错误
		if err != nil {
			log.Println("error:", err)
		}
	}
}
