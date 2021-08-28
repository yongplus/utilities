### 协程

Coroutine模块用于并发需求开发。

模块安装：
> go get github.com/yongplus/unity/coroutine   

调用代码：   
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