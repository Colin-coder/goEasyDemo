package main

import (
	"errors"
	"fmt"
)

// recover()用于将panic的信息捕捉。recover必须定义在panic之前的defer语句中。

func tryRecover() {
	defer func() {
		r := recover() // 接受panic的值

		// 把r强制转成error类型
		if err, ok := r.(error); ok { // 转化成功
			fmt.Println("error occur---", err)
		} else { // 转化失败，panic的不是error类型
			panic(fmt.Sprintf("dont know what to do: %v", r))
		}
	}()

	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(errors.New("this is a err"))

	//panic(123)
}

func main() {
	tryRecover()
}
