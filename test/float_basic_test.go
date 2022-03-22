package test

import (
	"strconv"
	"testing"

	"github.com/smiecj/go_common/util/log"
)

// 测试保存浮点数有效位小数 - 基本功能
func TestFloatSaveNumBasic(t *testing.T) {
	f := 0.000123
	log.Info(strconv.FormatFloat(f, 'f', -2, 64))
	log.Info(strconv.FormatFloat(f, 'f', -1, 64))
}
