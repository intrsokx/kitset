package main

import (
	"fmt"
	"github.com/intrsokx/kitset/upakit/pkg/ratelimiter"
	"sync"
	"time"
)

var lmt = ratelimiter.NewLimiter(100)

//go run jiujiulimiter.go > test.txt
//true 170
//false 30

//true = capacity + waitTime
//一开始，同时进来200个协程，都去获取令牌，其中因为令牌桶容量为120（qps*1.2）的缘故，前120个协程可以直接获取到令牌；
//然后没有拿到令牌的70个协程等待100ms，在此期间，令牌桶中又生成了10个令牌，分配给这70个协程中的前10个。所以最后是有130个协程获取到了令牌。
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			//等待100ms，等待生成10个token的时间，若是还获取不到，则返回false
			flag := lmt.WaitMaxDuration(1, time.Millisecond*100)
			fmt.Println(flag)
		}()
	}

	wg.Wait()

	test()
}

func test() {
	lmt = ratelimiter.NewLimiterByQpm(6)
	fmt.Println(lmt.Rate())
	for i := 0; i < 10; i++ {
		fmt.Println(lmt.WaitMaxDuration(1, 100))
	}
	time.Sleep(time.Second * 10)
	for i := 0; i < 5; i++ {
		fmt.Println(lmt.WaitMaxDuration(1, 100))
	}
}
