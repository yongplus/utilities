## 协程

Coroutine模块用于并发需求开发。

#### Installation：
> go get github.com/yongplus/unity/coroutine

<br>  

#### Methods:
```go

//New an instance.
New(coroutineNums int,chanBufferSize int) *Coroutine

//Set a worker with a return result.
func (m *Coroutine) SetWorker(function(val interface{}) interface{} {})

//Set a worker without any return result.
func (m *Coroutine) SetWorker2(function(val interface{}){})

//Push the value to the function passed into the SetWorker*
func (m *Coroutine) Push(val interface{})

//Set a listener of the result returned from the workers 
func (m *Coroutine) SetListener(recv func(interface{}))

//Get the channel of result returned from the workers
func (m *Coroutine) RecvChans() chan interface{}

//Waiting for the all the workers and listener to finishe and exit. 
Wait() 
```
<br>  
  

#### Example：
```go
//初始化，(协程数，channel缓存数)
crt := coroutine.New(6,1)
//设置执行函数，val是通过push写入的数据，透传给回调函数。
crt.SetWorker2(func(val interface{}) {

})
for i:=0;i<10000;i++{
    crt.Push(i) //向内部channel写入数据
}
//等待所有协程完成工作并退出。
crt.Wait()
```