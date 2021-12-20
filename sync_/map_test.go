package sync_

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	mapSize = 50
)

// 测试 sync.Map 并发读写
func TestSyncMap(t *testing.T) {
	m := sync.Map{}

	// 并发读和写
	noticeChan := make(chan int)
	valChan := make(chan interface{}, 100)
	defer close(valChan)
	for index := 0; index < mapSize; index++ {
		go func(i int) {
			<-noticeChan
			m.Store(i, i)
		}(index)
	}

	for index := 0; index < mapSize; index++ {
		go func(i int) {
			<-noticeChan
			val, _ := m.Load(i)
			valChan <- val
		}(index)
	}

	close(noticeChan)
	<-time.After(5 * time.Second)
	require.Equal(t, mapSize, len(valChan))
}
