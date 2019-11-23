package main

import (
	"fmt"
	"math/rand"
	"time"
)

type channelData struct {
	data     int
	resource int
}

func generator(index int) chan channelData {
	out := make(chan channelData)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- channelData{i, index}
			i++
		}
	}()
	return out
}

// 返回一个只能接收数据的 channel，不能发送出数据
func createWorker(id int) chan<- channelData {
	c := make(chan channelData)
	go worker(id, c)
	return c
}

func worker(id int, ch chan channelData) {
	for n := range ch { // 如果ch 被close，则这个循环结束
		fmt.Printf("worker %d received %d from resource %d\n", id, n.data, n.resource)
	}
}

func main() {
	//var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(0), generator(1)
	w := createWorker(0)
	hasValue := false
	var n channelData
	for {

		var activeChannel chan<- channelData
		if hasValue {
			activeChannel = w
		}
		select { // 哪个channel先来，先处理哪个
		case n = <-c1:
			//fmt.Println("received from c1", n)
			//w <- n
			hasValue = true
		case n = <-c2:
			//fmt.Println("received from c2", n)
			//w <- n
			hasValue = true
			//default: // 非阻塞式的从channel中获取数据
			//	fmt.Println("no value received")
		case activeChannel <- n:
			hasValue = false
		}
	}
}
