package main

import (
	"github.com/yongplus/unity/coroutine"
	"log"
)

func main(){

	crt := coroutine.New(6,1)
	crt.SetWorker(func(val interface{}) {
		log.Println(val.(int))
	})

	crt.Push(1)
	crt.Wait()
	log.Println("执行完成")
}