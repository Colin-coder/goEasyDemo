package main

import (
	"fmt"
	"sync"
)

// 这里使用sync.WaitGroup来实现同步

// 创建任务，并返回任务
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//wg: wg,

		// 函数式编程，把动作进行包装
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

// 执行任务的函数
// 这个任务是，通过channel接受一个传入，打印出来输入
// 并在外部等待所有任务的结束
func doWorker(id int, w worker /*wg *sync.WaitGroup*/) {
	for n := range w.in { // 如果ch 被close，则这个循环结束
		fmt.Printf("worker %d received %c\n", id, n)

		//done <- true
		//wg.Done()
		w.done()
	}
}

// 任务参数
type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup

	done func()
}

func main() {
	chanDemo()
}

func chanDemo() {
	// <-chan  只能发送数据的channel
	// 只能收数据的channel
	var workers [10]worker

	// 系统提供的
	var wg sync.WaitGroup

	// 创建任务
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	// 计数 20
	//wg.Add(20)

	// 每个任务通过channel传入一个参数
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i
	}
	// wait for finished
	//for _, worker := range workers {
	//	<-worker.done
	//}

	//  第二次传入参数，第二次执行任务
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
	}

	// 这里通过另一种 channel 阻塞等待的方式来实现并行goroutine之间的通信
	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}

	// 之前的做法是固定一个时间去等待，
	// 这种肯定不行，无法预估需要等待的具体时间
	//time.Sleep(time.Microsecond)

	// 等待所有
	wg.Wait()
}
