package norepeatsubstring

import (
	"testing"
)

/*

testing.T 测试代码运行正确性

1. go test .  运行所有测试文件，获取代码覆盖报告
（使用IDE查看测试代码覆盖度）

2. go tool cover 查看代码覆盖报告

*/
func TestLengthOfNonRepeatedSubString(t *testing.T) {

	tests := []struct {
		a string
		c int
	}{
		{"aaddva", 3},
		{"sdfa", 4},
		{"asdf", 4},
		{"asdfawerf", 7},
	}

	// 测试逻辑
	for _, test := range tests {
		if actual := LengthOfNonRepeatedSubString(test.a); actual != test.c {
			t.Errorf("error:[a:%s][b:%d][actual:%d]", test.a, test.c, actual)
		}
	}
}

/*
testing.B 性能测试

使用 pprof 优化性能

1. 将bench数据输出到cpu.out
go test -bench . -cpuprofile cpu.out

2. 使用pprof工具分析这个文件
go tool pprof cpu.out

3. 使用 web命令 查看svg图中统计的 耗时操作（使用前需要安装 http://www.graphviz.org/）
(pprof) web

*/

// 进行性能测试 每次运行时间 1189 ns/op
func BenchmarkSubStr(b *testing.B) {
	s := "aiswefoiwbaigbuawerngieugbh"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 11
	b.ResetTimer() // 重新计数，上面的代码不计入统计时间

	for i := 0; i < b.N; i++ {

		actual := LengthOfNonRepeatedSubString2(s)
		if actual != ans {
			b.Errorf("error [actual:%d]", actual)
		}
	}
}
