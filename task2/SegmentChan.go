package task2

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/
func BaseChanUse() {
	var ch = make(chan int)
	var timer = time.NewTimer(time.Second * 10)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("发送：", i)
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case num := <-ch:
				fmt.Println("收到：", num)
			case <-timer.C:
				runtime.Goexit()
			}
		}
	}()
	wg.Wait()
}

func BufferChanUse() {
	var ch = make(chan int, 5)
	fmt.Println(len(ch), cap(ch))
	var timer = time.NewTimer(time.Second * 10)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Println("发送：", i)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case num := <-ch:
				fmt.Println("收到：", num)
			case <-timer.C:
				runtime.Goexit()
			}
		}
	}()
	wg.Wait()
}

type Queue struct {
	item []int
	cd   *sync.Cond
}

// FIFO
func (q *Queue) put(item int) {
	q.cd.L.Lock()
	defer q.cd.L.Unlock()
	q.item = append(q.item, item)

	//notify
	q.cd.Broadcast()
}

func (q *Queue) get(timeOut time.Duration) (int, bool) {
	fmt.Println("start get...")
	q.cd.L.Lock()
	defer q.cd.L.Unlock()
	var expireOut = false
	var timer = time.AfterFunc(timeOut, func() {
		q.cd.L.Lock()
		expireOut = true
		q.cd.Broadcast()
		q.cd.L.Unlock()
	})
	defer timer.Stop()
	for len(q.item) == 0 && !expireOut {
		//wait
		q.cd.Wait()
	}
	if len(q.item) == 0 && expireOut {
		fmt.Println("waiting get timeout...")
		return 0, false
	}
	item := q.item[0]
	q.item = q.item[1:]
	return item, true
}

func ChanAndCond() {
	//var wg = sync.WaitGroup{}
	q := &Queue{
		item: make([]int, 0),
		cd:   sync.NewCond(&sync.Mutex{}),
	}

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		for i := 0; i < 10; i++ {
			q.put(i)
			fmt.Println("压入栈：", i)
			time.Sleep(time.Second * 1)
		}
	}()
	for {
		n, b := q.get(time.Second * 10)
		if !b {
			return
		}
		fmt.Println("出栈：", n)
	}
	//wg.Wait()
}
