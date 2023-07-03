package sync_

import (
	"sync"
	"testing"
)

func TestWaitGroupDoneFirst(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Done()
	wg.Wait()
}
