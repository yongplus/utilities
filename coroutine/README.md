## 协程

Coroutine模块用于并发需求开发。

#### 模块安装：
> go get github.com/yongplus/unity/coroutine

<br>  

#### Methods:
```go

//Init the struct.
New(coroutineNums int,chanBufferSize int)

//Set a worker with a return result.
SetWorker(function(val interface{}){} interface{})

//Set a worker without a return result.
SetWorker2(function(val interface{}){})

//Push the value to the function passed into SetWorker*
Push(val interface{})

//Set a istener of the return result from workers 
SetListener(recv func(interface{}))

//Waiting for the all the workers and Listener to finishe and exit. 
Wait() 
```
<br>  
  

#### 调用代码：
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