package main

import (
	"github.com/yongplus/utility/coroutine"
	"log"
	"time"
)

func main() {
	crt := coroutine.New(1, 1)
	/*
		crt.SetWorker(func(val interface{}) interface{} {

			log.Println(val.(int))
			return val.(int) * val.(int)
		})

			crt.SetListener(func(val interface{}) {
				if 999*999 == val {
					time.Sleep(time.Second * 10)
				}
				log.Printf("收到结果1：%d", val)
			})*/

	go (func(valChans chan interface{}) {
		for {
			val := <-valChans
			if val == nil {
				log.Println("退出")
				return
			}
			if 999*999 == val.(int) {
				time.Sleep(time.Second * 10)
			}
			log.Printf("收到结果2：%d", val)
		}
	})(crt.RecvChans())

	for i := 0; i < 1000; i++ {
		crt.Push(i)
	}
	log.Println("等待退出")
	crt.Wait()
	log.Println("执行完成")
}
