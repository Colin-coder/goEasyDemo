package main

/*
	传统测试 Vs 表格测试

	传统测试：
	1. 错了一个全挂了

	表格驱动测试
	1. 分离测试数据和测试逻辑
	2. 明确的出错信息
	3. 可以部分失败
*/

func Add(a, b int32) int32 {
	return a + b
}
