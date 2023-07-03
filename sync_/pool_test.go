package sync_

import (
	"bytes"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	counter atomic.Uint32
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
