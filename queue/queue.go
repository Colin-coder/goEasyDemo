package queue

// 将 []int 别名化，作为继承功能使用，继承了slice []int
/*
type Queue []int

func (q *Queue)Push(v int)  {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue)IsEmpty() bool {
	return len(*q) == 0
}
*/

// 泛型编程，Queue目前支持任何类型
// 日志功能测试
//		eg ...
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
