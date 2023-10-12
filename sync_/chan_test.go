package sync_

import (
	"fmt"
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

// 合并两个 channel 的数据 （nil channel 的用法）
func TestMergeTwoChan(t *testing.T) {
	chanOdd := make(chan int)
	chanEven := make(chan int)

	// odd and even producer
	go func(c chan<- int) {
		for i := 1; i < 10; i += 2 {
			chanOdd <- i
		}
		close(c)
	}(chanOdd)
	go func(c chan<- int) {
		for i := 2; i < 10; i += 2 {
			c <- i
		}
		close(c)
	}(chanEven)

	// merge
	chanAll := make(chan int)
	go func(chanToMerge chan<- int, chan1, chan2 <-chan int) {
		defer close(chanToMerge)
		for chan1 != nil || chan2 != nil {
			select {
			case i, ok := <-chan1:
				if !ok {
					chan1 = nil
					continue
				}
				chanToMerge <- i
			case i, ok := <-chan2:
				if !ok {
					chan2 = nil
					continue
				}
				chanToMerge <- i
			}
		}
	}(chanAll, chanOdd, chanEven)

	for i := range chanAll {
		fmt.Println(i)
	}
}

// todo: 测试 有 buffer chan
