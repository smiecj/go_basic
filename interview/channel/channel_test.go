package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAlphaNumberPrint(t *testing.T) {
	printer := AlphaAndNumberPrinter{}
	printer.Start()
}

// 从一个已经关闭的channel 中获取数据
func TestReadFromClosedChannel(t *testing.T) {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		c <- 1
		close(c)
	}()

	go func() {
		time.Sleep(time.Second)
		i := <-c
		// 第一次: 顺利获取数据
		fmt.Println(i)
		// 第二次: 获取到零值
		i = <-c
		fmt.Println(i)
		wg.Done()
	}()
	wg.Wait()
}
