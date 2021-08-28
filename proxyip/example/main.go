package main

import (
	"github.com/franela/goreq"
	proxyip "github.com/yongplus/unity/proxyip"
	"log"
	"strings"
)

func main(){

	//设置url 返回必须是 ip:port text格式 并且是\r\n分割
	proxyip := proxyip.New("http://webapi.http.zhimacangku.com/getip?num=10&type=1&pro=440000&city=0&yys=0&port=1&time=1&ts=0&ys=0&cs=0&lb=1&sb=0&pb=4&mr=1&regions=")

	//获取ip
	ip := proxyip.GetOne()
	if len(ip) == 0{ //如果获取失败
		log.Println(proxyip.Error())
		return
	}

	log.Println("使用代理ip："+ip)
	proxy := "http://" + ip

	url := "https://2021.ip138.com/"
	req := goreq.Request{Proxy: proxy, Method: "POST", Uri: url}
	req.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63030522)")

	rps, err := req.Do()
	if err != nil {
		if strings.Contains(err.Error(),"Timeout"){
			proxyip.DelOne(proxy)
		}

		log.Println(err.Error())
		return
	}
	body, _ := rps.Body.ToString()
	rps.Body.Close()
	//通过请求ip138 查看输出源码中的ip是代理ip即可。
	log.Println(body)

}
