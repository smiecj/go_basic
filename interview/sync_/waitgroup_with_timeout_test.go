package sync_

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
)

// https://github.com/lifei6671/interview-go/blob/master/question/q013.md
// 让 WaitGroup 支持超时功能
// timeout context + select 实现
func TestWaitWithTimeout(t *testing.T) {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	timeoutContext, _ := context.WithTimeout(context.Background(), timeout)
	waitGroupFinishChan := make(chan int)
	// 这里要注意实际使用的时候，一定要保证 wg 能结束，否则这里会造成内存泄漏
	go func() {
		wg.Wait()
		close(waitGroupFinishChan)
	}()
	select {
	case <-timeoutContext.Done():
		log.Info("[wait_with_timeout] timeout")
		return true
	case <-waitGroupFinishChan:
		log.Info("[wait_with_timeout] chan finish")
		return false
	}
}
