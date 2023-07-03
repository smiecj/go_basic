package context_

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
)

// 测试超时 ctx 功能
// 两个 超时 ctx 嵌套，先到 deadline 的结束，是否会触发另一个的结束?
// -- 不会，子 context 的结束并不会触发父 context 的结束, 反过来才成立
func TestContextWithTimeout(t *testing.T) {
	timeout10SecondCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	timeout5SecondCtx, _ := context.WithTimeout(timeout10SecondCtx, 5*time.Second)
	if deadline10Second, ok := timeout10SecondCtx.Deadline(); ok {
		log.Info("[TimeoutContext] 10 second deadline: %v", deadline10Second)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		<-timeout10SecondCtx.Done()
		log.Info("[TimeoutContext] ctx: 10 second")
		wg.Done()
	}()

	go func() {
		<-timeout5SecondCtx.Done()
		log.Info("[TimeoutContext] ctx: 5 second")
		wg.Done()
	}()

	wg.Wait()
}

func TestTimeoutGetErr(t *testing.T) {
	timeoutCtx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func(c chan<- int) {
		time.Sleep(time.Second)
		// 注意不能close
		c <- 0
	}(chan1)

	go func(c chan<- int) {
		time.Sleep(3 * time.Second)
		c <- 1 // this value will not print
	}(chan2)

	go func(c1 <-chan int, c2 <-chan int, ctx context.Context) {
		for i := 0; i < 3; i++ {
			select {
			case v := <-c1:
				println(v)
			case v := <-c2:
				println(v)
			case <-ctx.Done():
				print(ctx.Err().Error()) // context deadline exceeded
				return
			}
		}
	}(chan1, chan2, timeoutCtx)

	time.Sleep(5 * time.Second)
}
