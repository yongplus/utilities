package main

import (
	"github.com/yongplus/unity/coroutine"
	"log"
)

func main(){

	crt := coroutine.New(6,1)
	crt.SetWorker(func(val interface{}) {

	})
	for i:=0;i<10000;i++{
		crt.Push(i)
	}
	crt.Wait()
	log.Println("执行完成")
}