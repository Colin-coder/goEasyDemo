package main

import (
	"bufio"
	"fmt"
	"goEasyDemo/functionStyle/fibonacci"
	"os"
)

// 资源管理方式：defer
// 一个动作对应另一个动作（文件打开关闭，数据库连接关闭，等等）
// open/close  lock/unlock

// defer 函数退出前打印
// 多个defer，栈结构，先进后出的调用顺序

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)

	panic("error occur")
	fmt.Println(4)
}

func tryDeferTm() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("too many")
		}
	}
}
func writeFIle(filename string) {

	// 创建文件
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//err = errors.New("nk error") //自定义的error
	if err != nil {
		// 强制类型转换
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}

	// 关闭文件
	defer file.Close()

	// 写入缓存
	writer := bufio.NewWriter(file)

	// 将缓存中数据写入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	//tryDeferTm()
	writeFIle("fib.txt")
}
