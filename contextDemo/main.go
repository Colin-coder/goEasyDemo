package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)
// 参考链接：https://blog.csdn.net/u011957758/article/details/82948750

var generalCtx context.Context

func init() {
	generalCtx = context.Background()
}

func funcTimeoutCtx() {

	// 带超时的，超时过后，case <-ctx.Done(): 会进入
	ctx, cancel := context.WithTimeout(generalCtx, 10*time.Second)
	chiHanBao2(ctx)

	defer cancel()
}

func chiHanBao2(ctx context.Context) {
	// 个数
	n := 0
	for {
		time.Sleep(time.Second)
		select {
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("吃了 %d 个汉堡 \n", n)
		case <-ctx.Done():
			fmt.Println("stop")
			return
		}
	}
}

const (
	trace   = "trace_id"
	session = "session_id"
)

func funcWithValue() {
	ctx := context.WithValue(generalCtx, trace, "123")
	cctx := context.WithValue(ctx, session, "666")

	// cctx继承ctx 也带着ctx的value值
	traceIDSessionID(cctx)
}

func traceIDSessionID(cctx context.Context) {
	fmt.Println(cctx.Value(trace).(string))

	fmt.Println(cctx.Value(session).(string))
}

func funcCancelCtx() {
	ctx, cancel := context.WithCancel(generalCtx)

	// 开启goroutine 把context传给子goroutine ，在主中cancel一下，子收到done的消息，进行一些操作
	eatNum := chiHanBao(ctx)
	for n := range eatNum {
		if n >= 10 {
			cancel()
			break
		}
	}

	fmt.Println("正在统计结果。。。")
	time.Sleep(1 * time.Second)
}

func chiHanBao(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0
	go func() {
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d 秒，吃了 %d 个汉堡 \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("我吃了 %d 个汉堡\n", n)
			}
		}
	}()

	return c
}

func main() {
	//funcCancelCtx()
	//	funcTimeoutCtx()
	funcWithValue()
}
