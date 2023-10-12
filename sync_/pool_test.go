package sync_

import (
	"bytes"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	byteLen = 1024
	counter atomic.Uint32
	longStr = strings.Repeat("S", byteLen)
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			counter.Add(1)
			return bytes.Buffer{}
		},
	}

	go func() {
		time.Sleep(time.Second)
		buf := pool.New().(bytes.Buffer)
		defer pool.Put(buf)
		defer buf.Reset()
		buf.WriteString("hello")
		println(buf.String())
	}()

	go func() {
		time.Sleep(2 * time.Second)
		buf := pool.Get().(bytes.Buffer)
		buf.WriteString("world")
		println(buf.String())
	}()

	time.Sleep(5 * time.Second)
	println(counter.Load())
}

func BenchmarkByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.Buffer{}
		buf.WriteString(longStr)
		io.Copy(io.Discard, &buf)
	}
}

func BenchmarkByteBufferWithPool(b *testing.B) {
	// var newCount int64

	pool := sync.Pool{
		New: func() any {
			// atomic.AddInt64(&newCount, 1)
			return new(bytes.Buffer)
		},
	}

	// fmt.Println()
	for i := 0; i < b.N; i++ {
		buf := pool.Get().(*bytes.Buffer)
		buf.WriteString(longStr)
		io.Copy(io.Discard, buf)
		buf.Reset()
		pool.Put(buf)
	}

	// fmt.Printf("test count: %d, buf new count: %d\n", b.N, newCount)
}
