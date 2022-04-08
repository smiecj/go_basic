package basic

import (
	"testing"

	"github.com/smiecj/go_common/util/log"
)

type Student struct {
	Age int
}

// 直接获取 map 元素
func TestMapGet(t *testing.T) {
	m := map[string]Student{"xiaoming": {Age: 1}}
	// m["xiaoming"].Age = 1
	stu := m["xiaoming"]
	log.Info("[map] get student age: %d", stu.Age)
}
