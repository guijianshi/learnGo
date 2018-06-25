package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberWorker = 4  // 消费者数量,即goroutine数量
	taskLoad     = 10 // 待处理数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	tasks := make(chan string, taskLoad)

	wg.Add(numberWorker)

	for gr := 1; gr <= numberWorker; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("task: %d", post)
	}
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {

	defer wg.Done()
	for {
		// 获取任务
		task, ok := <-tasks
		if !ok {
			fmt.Printf("%d 工作人员已经结束工作\n", worker)
			return
		}

		fmt.Printf("%d 号工作人员开始 %s工作\n", worker, task)

		// 随机等待一段时间
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("%d 号工作人员已经完成了 %s 工作\n", worker, task)
	}
}
