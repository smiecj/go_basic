// package channel channel 相关题目
package channel

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 题目: 开启两个协程，分别打印字母和数字，最终形成字母数字交错打印的效果（可连续）
// https://github.com/lifei6671/interview-go/blob/master/question/q001.md
type AlphaAndNumberPrinter struct{}

// 实现: 典型的 生产者 - 消费者模型，开启三个协程，一个消费者 负责消费一个 channel 的数据，两个生产者，都往一个 channel 中写数据
// 关闭channel: 由于 channel 只能被关闭一次，并且关闭后也无法写入数据，所以必须按照 生产者停止生产 -> 消费者停止消费的顺序完成
// 综上，使用 waitGroup + channel 来实现
// 扩展: 如果要求字母和数字交错打印? -- 需要两个 chan 和 两个消费者，并保证这两个 routine 之间要通信
func (printer *AlphaAndNumberPrinter) Start() {
	c := make(chan int)
	retWaitGroup := sync.WaitGroup{}
	retWaitGroup.Add(1)

	stopWaitGroup := sync.WaitGroup{}
	stopWaitGroup.Add(2)

	// 触发任务开始，避免先创建的 routine 快速执行
	startSignal := make(chan int)

	// 独立协程: 触发消费者的关闭
	go func() {
		stopWaitGroup.Wait()
		close(c)
		// 这里也可以直接 close c, 然后在消费者中 使用 ok 来判断 c 是否关闭即可
		// 有 buffer 的 channel 是否可以用这种方式? -- 在 chan_test 中测试
	}()
	// 消费者
	go func() {
		for {
			currentChar, ok := <-c
			if ok {
				log.Printf("current char: %c", currentChar)
			} else {
				break
			}
		}
		retWaitGroup.Done()
	}()
	// 生产者 - 数字
	go func() {
		baseChar := '0'
		// ticker := time.NewTicker(time.Second)
		hasProduceCount, totalProduceCount := 0, 10
		<-startSignal
		for {
			// for range ticker.C {
			c <- rand.Intn(10) + int(baseChar)
			hasProduceCount++
			if hasProduceCount == totalProduceCount {
				break
			}
		}
		stopWaitGroup.Done()
	}()
	// 生产者 - 字母
	go func() {
		baseChar := 'a'
		// ticker := time.NewTicker(time.Second)
		hasProduceCount, totalProduceCount := 0, 10
		<-startSignal
		for {
			// for range ticker.C {
			c <- rand.Intn(26) + int(baseChar)
			hasProduceCount++
			if hasProduceCount == totalProduceCount {
				break
			}
		}
		stopWaitGroup.Done()
	}()
	log.Printf("print will begin in 5 second ...")
	time.Sleep(5 * time.Second)
	close(startSignal)
	retWaitGroup.Wait()
}
