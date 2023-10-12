package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Jeffail/tunny"

	"github.com/panjf2000/ants/v2"
)

const (
	parallel  = 10
	taskCount = 100
)

func TestTunnyPool(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(taskCount)
	pool := tunny.NewFunc(parallel, func(payload interface{}) interface{} {
		fmt.Printf("%d execute\n", payload)
		time.Sleep(3 * time.Second)
		fmt.Printf("%d finish\n", payload)
		wg.Done()
		return payload
	})
	defer pool.Close()

	for i := 0; i < taskCount; i++ {
		// pool.Process 是同步方法
		go func(i int) {
			pool.Process(i)
		}(i)
	}

	// 总执行时间: 100 / 10 * 3 = 30s
	wg.Wait()
}

func TestAntsPool(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(taskCount)
	pool, _ := ants.NewPoolWithFunc(parallel, func(i interface{}) {
		fmt.Printf("%d execute\n", i)
		time.Sleep(3 * time.Second)
		fmt.Printf("%d finish\n", i)
		wg.Done()
	})
	defer pool.Release()

	for i := 0; i < taskCount; i++ {
		// pool.Invoke 是异步方法
		pool.Invoke(i)
	}

	wg.Wait()
}
