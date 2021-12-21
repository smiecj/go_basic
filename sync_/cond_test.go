package sync_

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 测试 Cond
// 场景: 一个写协程，一个校验协程
// 写协程负责将，校验协程在检查写入成功之后才会结束
func TestCond(t *testing.T) {
	lock := sync.Mutex{}
	cond := sync.NewCond(&lock)
	var sig int64 = 1
	finishChan := make(chan int)

	// 检查协程
	go func() {
		cond.L.Lock()
		for atomic.LoadInt64(&sig) != 2 {
			cond.Wait()
		}
		cond.L.Unlock()
		close(finishChan)
	}()

	// 修改协程
	go func() {
		// 保证修改在等待之后
		time.Sleep(2 * time.Second)
		cond.L.Lock()
		atomic.AddInt64(&sig, 1)
		cond.L.Unlock()
		cond.Signal()
	}()

	<-finishChan
}
