package study

import (
	"fmt"
	"runtime"
)

func chanFribo(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch: // 循环读取
			fmt.Print(num, " ")
		case <-quit:
			//退出当前go程
			runtime.Goexit()
		}
	}
}

/*
斐波那契 current=(current-1)+(current-2)
1 1 2 3 5 8 13 21 34 55 89 144 ...
*/
func PrinterFibonacci(ch chan bool) {
	fmt.Println("斐波那契")
	input := make(chan int)
	go chanFribo(input, ch)
	var a, b int = 1, 1
	for i := 0; i < 20; i++ {
		//fmt.Print(a, " ")
		input <- a
		a, b = b, a+b
		//time.Sleep(time.Second * 1)
	}
	fmt.Println("ending")
	//运行完 通知结束
	ch <- true
}
