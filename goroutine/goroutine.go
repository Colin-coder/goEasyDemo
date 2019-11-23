package main

import (
	"fmt"
	"time"
)

/*
	go routine

	协程 轻量级线程
	非抢占式多任务处理
*/

func test2() {
	for i := 0; i < 10; i++ {
		go func() {
			for true {
				fmt.Printf("hello from goroutine %d\n", i)
			}
		}()
	}
	time.Sleep(time.Microsecond)
}

// 多goroutine 时 使用 race检测数据冲突问题
// go run -race goroutine.go
func test1() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	fmt.Println("a is ", a)

	time.Sleep(2 * time.Second)
}

func main() {
	test1()
}
