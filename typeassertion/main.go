package main

import (
	"fmt"
	"time"
)

// 一个普通的b interface{} 包含两个成员变量，
// 1. 实现者的类型，通过 a.(type) 来获取
// 2. 实现者的指针，它指向实现者

//用于转换函数里interface{}类型的参数
// 任何类型都是interface{}，它可以代表任意类型
func add(a, b interface{}) {

	// 这里用OK接受，是否转换成功
	// 将 a interface{} 类型转换为  av int16
	av, ok := a.(int16)
	fmt.Printf("%T, %v\n", av, av)
	if !ok {
		fmt.Println("error type assertion!")
		return
	}

	switch t := a.(type) {
	case int:
		fmt.Printf("type [%T] add res[%d]\n", t, a.(int)+b.(int))
	case int16:
		fmt.Printf("type [%T] add res[%d]\n", t, a.(int16)+b.(int16))
	case float32:
		fmt.Printf("type [%T] add res[%f]\n", t, a.(float32)+b.(float32))
	case float64:
		fmt.Printf("type [%T] add res[%f]\n", t, a.(float64)+b.(float64))
	default:
		fmt.Printf("type [%T] not support!\n", t)
	}
}
func test1() {
	add(1, 2)
	add(int16(1), int16(2))
	add(float32(1.1), float32(2.2))
	add(float64(1.1), float64(2.2))
	add(true, false)
}

// 用于转换结构体的interface{}类型字段

type NetMsg struct {
	MsgID int16
	Data  interface{}
}

type Cat struct {
	name string
	age  int16
}

type Dog struct {
	name string
	age  int32
}

type human struct {
	name string
	age  int64
}

type Fish struct {
	name string
	age  int32
	evn  int64
}

func test2() {
	// 结构体中Data  interface{} 可替换成其他结构体
	msg1 := NetMsg{1, Cat{"Qian", 1}}
	msg2 := NetMsg{2, Dog{"doge", 8}}
	msg3 := NetMsg{3, Dog{"allu", 18}}
	msg4 := NetMsg{4, human{"nk", 25}}
	msg5 := NetMsg{5, Fish{"fish", 25, 58}}
	msgHandler(msg1)
	time.Sleep(2000 * time.Millisecond)
	msgHandler(msg2)
	time.Sleep(2000 * time.Millisecond)
	msgHandler(msg3)
	time.Sleep(2000 * time.Millisecond)
	msgHandler(msg4)

	time.Sleep(2000 * time.Millisecond)
	msgHandler(msg5)
}

func msgHandler(msg NetMsg) {
	switch msg.MsgID {
	case 1:
		cat := msg.Data.(Cat)
		fmt.Printf("Do Something with Msg 1 %v \n", cat)
	case 2:
		dog := msg.Data.(Dog)
		fmt.Printf("Do Something with Msg 2 %v \n", dog)
	case 5:
		fish := msg.Data.(Fish)
		fmt.Printf("Do Something with Msg 5 %v \n", fish)
	default:
		fmt.Printf("Error MsgID [%d] \n", msg.MsgID)
	}
}

func main() {

	//test1()
	//test2()

	// 打印结构体成员变量名称
	fish := Fish{name: "123"}
	fmt.Printf("%+v", fish)
}
