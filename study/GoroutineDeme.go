package study

import (
	"fmt"
	"runtime"
)

/*
*
注意1：main结束，代表go程也会结束
*/
func CreateRoutine() {
	defer fmt.Println("defer CreateRoutine")
	for i := 0; i < 10; i++ {
		fmt.Println("go routine index:", i)
		if i == 5 {
			//退出当前go程后面的逻辑不再执行
			runtime.Goexit()
		}
	}
	//调用了Goexit 这个就不会执行了
	fmt.Println("CreateRoutine go routine end")
}

/*
**
知识点2：runtime.gosched 让出cpu时间片
*/
func TestGosched() {
	runtime.Gosched()
	for i := 0; i < 10; i++ {
		fmt.Println("TestGosched index:", i)
	}
}

var channel1 = make(chan int)

func ChanTestWrite() {
	//make(chan Type)
	fmt.Println("逻辑走完")
	fmt.Println("channel管道开始写入数据")
	//无缓冲 channel
	channel1 <- 1
}

func ChannelTestRead() {
	num := <-channel1
	fmt.Println("读取管道数据")
	fmt.Println("收到数据：", num)
	fmt.Println("继续执行下面的逻辑")
}

// 默认的channel 是双向的 make(chan type)
// 有缓冲的channel  2个元素之前不会阻塞
var channelBuffer = make(chan int, 2)

func ChannelBufferWrite() {
	fmt.Println(len(channelBuffer), "-", cap(channelBuffer))
	for i := 0; i < 5; i++ {
		//channelBuffer <- i
		wt(channelBuffer, i)
		fmt.Println("写入：", i) //这个io延迟严重
	}
	//如果关闭以后，无法写  但是可读，读到的都是0
	close(channelBuffer)
}
func wt(out chan<- int, num int) {
	out <- num
	fmt.Println("写入：", num)
}
func ChannelBufferRead() {
	//for i := 0; i < 5; i++ {
	//	num, ok := <-channelBuffer
	//	if ok == false {
	//		fmt.Println("channel已经关闭")
	//	} else {
	//		fmt.Println("读取：", num)
	//	}
	//}

	for c := range channelBuffer {
		fmt.Println("读取：", c)
	}
	fmt.Println("读取完成")
}

func rd(in <-chan int) {
	fmt.Println("读取：", in)
}
