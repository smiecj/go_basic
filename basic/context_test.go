package basic

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
func TestTimeoutContext(t *testing.T) {
	timeout10SecondCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	timeout5SecondCtx, _ := context.WithTimeout(timeout10SecondCtx, 5*time.Second)
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
