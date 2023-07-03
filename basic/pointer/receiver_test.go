package pointer

import (
	"fmt"
	"testing"
)

type others struct{}

type receiver struct {
	name   string
	others others
}

func (r *receiver) setName(name string) {
	r.name = name
}

func (r receiver) getName() string {
	return r.name
}

func TestReceiver(t *testing.T) {
	// 对可寻址的对象，pointer receiver 和 value receiver 方法都能直接调用
	r := receiver{
		name: "haha",
	}
	r.setName("test")
	fmt.Println(r.getName())
}
