package queue

import "fmt"

// 用注释写文档
// go doc  /   godoc -http :6060
// example 使用测试
// 实例代码测试
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
