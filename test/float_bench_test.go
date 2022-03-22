package test

import (
	"strconv"
	"testing"
)

// 测试保存浮点数有效位小数 - 性能测试
func TestFloatSaveNumBench(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.FormatFloat(10.900, 'f', -1, 64)
	}
}
