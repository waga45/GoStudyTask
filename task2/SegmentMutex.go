package task2

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var globalNum int32 = 0
var lock sync.Mutex

func addGlobalNum(th int) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		globalNum += 1
		fmt.Println(th, "-globalNum:", globalNum)
		lock.Unlock()
	}
	defer runtime.Goexit()
}

/*
*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func BaseMutexUse() {
	timer := time.NewTimer(time.Second * 5)
	for i := 0; i < 10; i++ {
		go addGlobalNum(i)
	}
	<-timer.C
	fmt.Println("end")
}

func addWithAtomic(th int) {
	for i := 0; i < 1000; i++ {
		f := atomic.AddInt32(&globalNum, 1)
		fmt.Println(th, "-globalNum:", f)
	}
}
func carWithAtomic(th int) {
	for i := 0; i < 1000; i++ {
		f := atomic.AddInt32(&globalNum, -1)
		fmt.Println(th, "-globalNum:", f)
	}
}

/*
*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func AtomicUse() {
	timer := time.NewTimer(time.Second * 5)
	for i := 0; i < 10; i++ {
		go addWithAtomic(i)
	}
	for i := 10; i < 20; i++ {
		go carWithAtomic(i)
	}
	<-timer.C
	fmt.Println("last onece:", globalNum)
	fmt.Println("end")
}
