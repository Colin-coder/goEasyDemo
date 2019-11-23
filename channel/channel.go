package main

import (
	"fmt"
	"time"
)

// 返回一个只能接收数据的 channel，不能发送出数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, ch chan int) {
	for n := range ch { // 如果ch 被close，则这个循环结束
		fmt.Printf("worker %d received %s\n", id, string(rune(n)))
	}
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}

func chanDemo() {
	// <-chan  只能发送数据的channel
	// chan<-  只能收数据的channel
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Microsecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	//c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(1, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}
