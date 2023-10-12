package context_

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {
	cancelCtx, cancel := context.WithCancel(context.Background())
	valCtx := context.WithValue(cancelCtx, "key", "value")
	withoutCancelCtx := context.WithoutCancel(valCtx)
	// timeoutCtx, _ := context.WithTimeout(withoutCancelCtx, 5*time.Second)
	timeoutCtx, _ := context.WithTimeoutCause(withoutCancelCtx, 5*time.Second, errors.New("timeout with 5 seconds"))
	_ = context.AfterFunc(timeoutCtx, func() {
		fmt.Println("after timeout execute after func!")
	})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		<-cancelCtx.Done()
		fmt.Println("cancel ctx done")
		wg.Done()
	}()
	go func() {
		// will not cancel by cancelCtx until timeout
		<-timeoutCtx.Done()
		fmt.Println("timeout ctx done")
		fmt.Println(fmt.Sprintf("timeout ctx val: %v", timeoutCtx.Value("key")))
		fmt.Println("timeout ctx cause: " + context.Cause(timeoutCtx).Error())
		wg.Done()
	}()
	cancel()
	wg.Wait()
	fmt.Println("test finish")

	context.TODO()
	context.Background()
}
