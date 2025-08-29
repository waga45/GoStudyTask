package study

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var cond sync.Cond
var lock sync.Mutex
var globalNum = 100

func addNum(t int, num int) {
	lock.Lock()
	globalNum += num
	fmt.Println("Add线程：", t, " num:", globalNum)
	lock.Unlock()
}

func catNum(t int, num int) {
	lock.Lock()
	defer lock.Unlock()
	globalNum -= num
	fmt.Println("Cat线程：", t, " num:", globalNum)
}
func TestMutex(ch chan bool) {
	fmt.Println("start")
	for i := 0; i < 50; i++ {
		go addNum(i, 100)
	}
	for i := 50; i < 100; i++ {
		go catNum(i, 100)
	}
	time.Sleep(time.Second * 6)

	fmt.Println("end", globalNum)
	ch <- true
	runtime.Goexit()
}

type ConfigMapper struct {
	config map[string]interface{}
	mu     sync.RWMutex
}

func (c *ConfigMapper) GetConfig(name string) (interface{}, bool) {
	//读锁
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.config[name]
	return value, ok
}
func (c *ConfigMapper) SetConfig(name string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.config[name] = value
}

func TestRWMutex() {
	configMapper := &ConfigMapper{config: make(map[string]interface{})}
	configMapper.SetConfig("app_name", "My Application")
	configMapper.SetConfig("version", "1.0.0")
	configMapper.SetConfig("timeout", "30s")

	var wg sync.WaitGroup
	//读取模拟
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				if value, exists := configMapper.GetConfig("app_name"); exists {
					fmt.Printf("Reader %d: %s\n", id, value)
				}
				time.Sleep(10 * time.Millisecond) // 模拟读取耗时
			}
		}(i)
	}
	wg.Add(1)
	//写入模拟
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond) // 等待一段时间后更新
		configMapper.SetConfig("app_name", "Updated Application")
		fmt.Println("配置已更新")
	}()
	wg.Wait()
	v, _ := configMapper.GetConfig("app_name")
	fmt.Println("更新后值：", v)

}
