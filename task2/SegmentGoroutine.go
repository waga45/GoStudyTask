package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/
func BaseGoroutine() {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("奇数:", i)
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数:", i)
			}
		}
	}()
	wg.Wait()
}

type Account struct {
	id      int
	balance int
	record  []string
}

// 存钱
func (a *Account) saveBalance(_balance int) int {
	a.balance += _balance
	r := fmt.Sprintf("%s存了：%d", time.DateTime, _balance)
	a.record = append(a.record, r)
	return a.balance
}

/*
*设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
 */
func TaskSchedul() {
	as := []Account{{id: 1, balance: 10000}, {id: 2, balance: 100}, {id: 3, balance: 900}}
	var wg = sync.WaitGroup{}
	for i := 0; i < len(as); i++ {
		wg.Add(1)
		var num = rand.Intn(1000)
		go func(v *Account, num int) {
			defer wg.Done()
			var startTimeSec = time.Now().Second()
			fmt.Println("用户：", v.id, " 当前余额：", v.balance, " 拿着：", num, "准备存钱...")
			v.saveBalance(num)
			time.Sleep(time.Second * 2)
			fmt.Println("用户：", v.id, " 余额：", v.balance)
			var endTimeSec = time.Now().Second()
			fmt.Println("用户：", v.id, " 存钱耗时：", (endTimeSec - startTimeSec), " sec")
		}(&as[i], num)
	}

	wg.Wait()
	fmt.Println("所有账户：", as)
}
