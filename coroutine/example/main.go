package main

import (
	Coroutine "github.com/yongplus/unity/coroutine"
	"log"
)

func main(){
	chans := make(chan interface{},1)
	crt := Coroutine.New(10,chans)
	crt.AddWorker(func(val interface{}) {
			x := val.(int)
			log.Println(x)
	})

	for i:=0;i<100;i++{
		chans	<- i
	}
	crt.Wait()
	log.Println("执行完成")
}