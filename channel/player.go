package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

// 用来等待程序结束
var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main()  {
	// 创建一个无缓存管道
	court := make(chan int)

	// 计数器加2, 表示要等待两个goroutine
	wg.Add(2)

	// 启动两个玩家
	go player("lin", court)
	go player("lv", court)

	// 发球(初始化管道内数据)
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int)  {
	// 在函数退出时通知mian函数工作已经完成
	defer wg.Done()
	for {
		// 等待球被击打过来(ball 管带内数据, ok 管道返回标志)
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 取随机数模仿是否成功击打
		n := rand.Intn(100)
		if n%14 == 0 {

			fmt.Printf("Player %s Missed\n", name)
			// 击打失败关闭管道(比赛结束)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		// 击球数加一, 击球数写入管道
		ball++
		court <- ball
	}
}
