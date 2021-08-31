package main

import (
	"github.com/yongplus/unity/coroutine"
	"log"
	"time"
)

func main() {

	crt := coroutine.New(1, 1)
	crt.SetWorker(func(val interface{}) interface{} {
		log.Println(val.(int))
		return val.(int) * val.(int)
	})

	crt.SetListener(func(val interface{}) {
		if 999*999 == val {
			time.Sleep(time.Second * 10)
		}
		log.Printf("收到结果：%d", val)
	})

	for i := 0; i < 1000; i++ {
		crt.Push(i)
	}
	log.Println("等待退出")
	crt.Wait()
	log.Println("执行完成")
}
