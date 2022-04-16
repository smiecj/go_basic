// package escape 内存逃逸场景总结
package escape

import (
	"io"
	"testing"
)

type People struct {
	Age int
}

type fakeReader struct {
}

func (reader *fakeReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

// 返回指针，将会逃逸
//go:noinline
func bornPeople() *People {
	p := new(People)
	p.Age = 1
	return p
}

// 发送到 chan, 也会逃逸，即使 chan 只能被内部访问
//go:noinline
func sendChan() {
	testChan := make(chan *People, 1)
	p := new(People)
	testChan <- p
}

// 只有在方法内部用到了指针，不会逃逸
//go:noinline
func bornPeopleNotReturn() {
	p := new(People)
	p.Age = 1
}

// slice 存 结构体，不会逃逸
//go:noinline
func makeStructSlice() {
	str := "test"
	strSlice := make([]string, 0, 10)
	strSlice = append(strSlice, str)
}

// slice 存指针，会导致指针对应的对象逃逸
//go:noinline
func makePointerSlice() {
	// strSlice := make([]*string, 0, 10)
	// strSlice = append(strSlice, str)
	str := "test"
	strSlice := []*string{&str}
	strSlice = append(strSlice, &str)
}

// 调用 interface 方法导致逃逸
// 直接声明成 具体实现类 不会导致逃逸
//go:noinline
func callInterfaceFunc() {
	// reader := fakeReader{}
	var reader io.Reader
	reader = &fakeReader{}
	reader.Read([]byte("str"))
}

func TestEscapePointer(t *testing.T) {
	bornPeople()
	sendChan()
	bornPeopleNotReturn()
	makeStructSlice()

	makePointerSlice()

	callInterfaceFunc()
}
