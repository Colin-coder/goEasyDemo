package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {

	// 测试数据
	tests := []struct {
		a, b, c int32
	}{
		{1, 2, 3},
		{0, 2, 2},
		{0, 0, 0},
		{1, -1, 2},
		{math.MaxInt32, 1, math.MinInt32},
	}

	// 测试逻辑
	for _, test := range tests {
		if actual := Add(test.a, test.b); actual != test.c {
			t.Errorf("error:[a:%d][b:%d]", test.a, test.b)
		}
	}
}
