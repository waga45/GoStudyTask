package study

import (
	"fmt"
	"time"
)

func StartTimer() {
	fmt.Println("StartTimer", time.Now())
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C
	myTimer.Stop()
	fmt.Println("时间：", nowTime)
}

func YanChiTimer() {
	myTimer := time.NewTimer(time.Second * 3)
	fmt.Println("创建定时器")
	<-myTimer.C
	fmt.Println("定时器到时，响铃")
}
