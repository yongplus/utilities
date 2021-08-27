package proxyip

import (
	"errors"
	"github.com/franela/goreq"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

type ProxyIp struct {
	url string //the api offering the ips
	mutex     *sync.RWMutex
	pool     []string //ips pool
	lasttime int64 //last update, It depends on available time of ips
	minsize int //refresh the pool of ip if the number of ip in the pool less than minsize
	err error
}

func New(url string) *ProxyIp {
	var proxyIp = &ProxyIp{
		mutex: &sync.RWMutex{},
		lasttime: 0,
		minsize: 1,
		url:url,
	}
	return proxyIp
}


func (m *ProxyIp) SetUrl(url string) {
	m.url = url
}

func (m *ProxyIp) GetOne() string {
	m.Update()
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if len(m.pool) < 1{
		return ""
	}
	if len(m.pool) == 1 {
		return m.pool[0]
	}

	i := rand.Intn(len(m.pool) - 1)

	return m.pool[i]
}

func (m *ProxyIp) DelOne(host string) {

	host = strings.Replace(host, "http://", "", -1)
	host = strings.Replace(host, "/", "", -1)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for idx, val := range m.pool {
		if val == host {
			m.pool = append(m.pool[:idx], m.pool[idx+1:]...)
		}
	}
}
func (m *ProxyIp) Update() {

	m.mutex.Lock()
	defer m.mutex.Unlock()
	if (time.Now().Unix()-m.lasttime) < 240 && len(m.pool) > m.minsize{
		return
	}
	rps, err := goreq.Request{Uri: m.url}.Do()
	if err != nil {
		m.setError(errors.New("Requesting the api failed:"+err.Error()))
		return
	}
	body, _ := rps.Body.ToString()
	rps.Body.Close()

	pool := strings.Split(strings.TrimSpace(body), "\r\n")
	//var validID = regexp.MustCompile(`(?:\d{1,3}\.){3}\d{1,3}$`)
	if len(pool) < 1 {
		m.setError(errors.New("Parsing the content of api failed,body:"+body))
		return
	}
	var validIP = regexp.MustCompile(`(?:\d{1,3}\.){3}\d{1,3}:[1-9][0-9]+$`)
	if !validIP.MatchString(pool[0]){
		m.setError(errors.New("Extracting the ips failed,body:"+body))
		return
	}
	m.lasttime = time.Now().Unix()
	m.pool = pool
}

func (m *ProxyIp) setError(err error){
	m.err = err
}

func (m *ProxyIp) Error() error{
	return m.err
}
