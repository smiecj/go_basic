// package sync_ 并发包的测试
package sync_

import (
	"sync"
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
	"github.com/stretchr/testify/require"
)

// 测试mutex 写锁
func TestMutex(t *testing.T) {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	hasLockTwice := false

	go func() {
		lock.Lock()
		log.Info("[sync] [error] Can not lock twice")
		hasLockTwice = true
		lock.Unlock()
	}()

	<-time.After(5 * time.Second)
	require.Equal(t, false, hasLockTwice)
}

// 测试 rwmetux 读写锁
func TestRWMutex(t *testing.T) {
	lock := sync.RWMutex{}
	lock.RLock()
	defer lock.RUnlock()
	writeLockWhenReadLock := false
	readLockTwice := false

	go func() {
		lock.Lock()
		writeLockWhenReadLock = true
		lock.Unlock()
	}()

	go func() {
		lock.RLock()
		readLockTwice = true
		lock.RUnlock()
	}()

	<-time.After(5 * time.Second)
	require.Equal(t, false, writeLockWhenReadLock)
	require.Equal(t, true, readLockTwice)
}
