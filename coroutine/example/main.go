package main

import (
	"github.com/yongplus/unity/coroutine"
	"log"
)

func main(){


	//设置url 返回必须是 ip:port text格式 并且是\r\n分割
	crt := coroutine.New(6,1)
	crt.SetWorker(func(val interface{}) {

	})
	for i:=0;i<10000;i++{
		crt.Push(i)
	}
	crt.Wait()
	log.Println("执行完成")
}