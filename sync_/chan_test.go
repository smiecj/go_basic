package sync_

import (
	"log"
	"sync"
	"testing"
	"time"
)

// 测试无 buffer chan，是否能一直 for get with ok?
// 预期输出: 每 2s 打印一行获取的整数；只有最后一次获取 ok 才是 false
// 结论: ok
func TestChanGetWithOK(t *testing.T) {
	// 消费者: 一直消费
	c := make(chan int)
	retWaitGroup := sync.WaitGroup{}
	retWaitGroup.Add(1)
	go func() {
		consumeTicker := time.NewTicker(time.Second)
		for range consumeTicker.C {
			currentInt, ok := <-c
			log.Printf("consumer: ok: %v, int: %d", ok, currentInt)
			if !ok {
				log.Printf("consumer finish: ok: %v, int: %d", ok, currentInt)
				break
			}
		}
		consumeTicker.Stop()
		retWaitGroup.Done()
	}()
	// 生产者: 生产10个数据
	go func() {
		produceTicker := time.NewTicker(2 * time.Second)
		hasProduceCount := 0
		for range produceTicker.C {
			c <- 1
			hasProduceCount++
			if hasProduceCount >= 10 {
				break
			}
		}
		produceTicker.Stop()
	}()
	retWaitGroup.Wait()
}

// todo: 测试 有 buffer chan
