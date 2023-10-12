// package sync_ 并发包的测试
package sync_

import (
	"sync"
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
	"github.com/stretchr/testify/require"
)

// 测试 Mutex 写锁
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

// 测试 RWMetux 读写锁
func TestRWMutex(t *testing.T) {
	lock := sync.RWMutex{}
	lock.RLock()
	defer lock.RUnlock()
	writeLockWhenReadLock := false
	readLockTwice := false

	go func() {
		// 保证读锁 在时间上先执行
		<-time.After(2 * time.Second)
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

// 测试锁拷贝
// -race:
func TestMutexCopy(t *testing.T) {
	mainLock := sync.Mutex{}

	// copy by mutex pointer
	go func(l *sync.Mutex) {
		l.Lock()
	}(&mainLock)

	// unlock in a method with value copy
	go func(l sync.Mutex) {
		time.Sleep(5 * time.Second)
		l.Unlock()
	}(mainLock)

	time.Sleep(time.Second)
	mainLock.Lock()
	mainLock.Unlock()
}

// 锁拷贝的一个更简单的示例
type Lock struct {
	mutex sync.Mutex
}

func (l *Lock) Lock() {
	l.mutex.Lock()
}

func (l Lock) Unlock() {
	l.mutex.Unlock()
}

func TestMutexCopySimple(t *testing.T) {
	l := Lock{}
	l.Lock()
	l.Unlock()
	l.Lock()
}
