## proxyip

proxyip用于获取代理ip，日常爬数据或者做其它需要N多ip资源时，可通过传入API Url路径，获取ip并管理。

#### 模块安装：
> go get github.com/yongplus/unity/proxyip   

<br>

#### 调用代码：
```go
//设置url 返回必须是 ip:port text格式 并且是\r\n分割
url := "http://webapi.http.zhimacangku.com/getip?num=10&type=1&pro=440000&city=0&yys=0&port=1&time=1&ts=0&ys=0&cs=0&lb=1&sb=0&pb=4&mr=1&regions="
proxyip := proxyip.New(url)

//获取ip
ip := proxyip.GetOne()
if len(ip) == 0{ //如果获取失败
    log.Println(proxyip.Error())
    return
}

//对于不稳定或无效的ip，可调用如下代码将其从ip池中移除
proxyip.DelOne(proxy)

```