package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	// 从管道中取值
	runner := <-baton

	fmt.Printf("运动员 %d 号开始赛跑\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("运动员 %d 在跑道上准备接力\n", newRunner)
		go Runner(baton) // 创建goruntine(可是会在管道取值那阻塞,等待管道当中添加值baton <- newRunner)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("运动员 %d 号 完成最后一棒,结束比赛\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("运动员 %d 将交接棒递到 %d 手中\n", runner, newRunner)
	baton <- newRunner
}
