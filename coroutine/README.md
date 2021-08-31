## 协程

Coroutine模块用于并发需求开发。

#### Installation：
> go get github.com/yongplus/unity/coroutine

<br>  

#### Methods:
```go

//Init the struct.
New(coroutineNums int,chanBufferSize int)

//Set a worker with a return result.
SetWorker(function(val interface{}){} interface{})

//Set a worker without any return result.
SetWorker2(function(val interface{}){})

//Push the value to the function passed into the SetWorker*
Push(val interface{})

//Set a listener of the result returned from the workers 
SetListener(recv func(interface{}))

//Waiting for the all the workers and listener to finishe and exit. 
Wait() 
```
<br>  
  

#### Example：
```go
//初始化，(协程数，channel缓存数)
crt := coroutine.New(6,1)
//设置执行函数，val是通过push写入的数据，透传给回调函数。
crt.SetWorker(func(val interface{}) {

})
for i:=0;i<10000;i++{
    crt.Push(i) //向内部channel写入数据
}
//等待所有协程完成工作并退出。
crt.Wait()
```