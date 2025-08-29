package study

import "fmt"

var OrderQueue = map[string]chan string{}

func SendMessage(topic string, value string) bool {
	ch, ok := OrderQueue[topic]
	if ok {
		fmt.Println("发送：", topic, value)
		ch <- value
	}
	return true
}
func Stop(topic string) bool {
	if len(topic) <= 0 {
		return false
	}
	ch, ok := OrderQueue[topic]
	if ok {
		close(ch)
	}
	return ok
}
func CreateQueue(topic string) bool {
	if len(topic) <= 0 {
		return false
	}
	_, ok := OrderQueue[topic]
	if !ok {
		OrderQueue[topic] = make(chan string, 5)
	}
	return true
}

func ConsumeMsg(topic string) {
	for k, v := range OrderQueue {
		if k == topic {
			for i := range v {
				fmt.Println(topic, "订单来啦:", i)
			}
		}
	}
	fmt.Println(topic, "监听结束")
}
