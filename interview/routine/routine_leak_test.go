package routine

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
	"github.com/stretchr/testify/require"

	"net/http"
	_ "net/http/pprof"
)

// 场景1: channel 的不合理使用

// 对 channel 只发送不接收，造成 routine 阻塞
func TestChannelSendNotReceive(t *testing.T) {

	// pprof to check routine
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	// 每次创建3个协程，往同一个 channel 发送数据
	create3Routine := func() int {
		c := make(chan int)
		for i := 0; i < 3; i++ {
			go func(i int) {
				c <- i
			}(i)
		}
		// 返回时消费一个数据
		return <-c
	}

	// 每隔 5s 创建协程
	for i := 0; i < 10; i++ {
		create3Routine()
		time.Sleep(10 * time.Second)
		// 打印当前协程数
		log.Info("routine count: %d\n", runtime.NumGoroutine())
	}
}

// 对 channel 只接收不发送
func TestChannelReceiveNotSend(t *testing.T) {

	// pprof to check routine
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	// 每次创建3个协程，从同一个 c 中消费数据
	create3Routine := func() chan<- int {
		c := make(chan int)
		for i := 0; i < 3; i++ {
			go func() {
				<-c
			}()
		}
		return c
	}

	// 每隔 5s 创建协程
	for i := 0; i < 10; i++ {
		// 获取 c 并生产一条数据
		c := create3Routine()
		c <- i
		time.Sleep(10 * time.Second)
		log.Info("routine count: %d\n", runtime.NumGoroutine())
	}
}

// 对空 channel 的误操作
func TestNilChannel(t *testing.T) {
	var c chan int
	require.Nil(t, c)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// close(c) // close: panic

	// producer
	go func() {
		c <- 1 // chansend: block
		wg.Done()
	}()

	// consumer
	go func() {
		_, ok := <-c // chanrecv: block
		fmt.Printf("ok: %v\n", ok)
		wg.Done()
	}()

	wg.Wait()
	close(c)
}

// 网络请求没有通过协程池控制或设置超时时间
func TestNetCallWithoutTimeout(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			http.Get("http://www.google.com")

			// set timeout
			// httpClient := http.DefaultClient
			// httpClient.Timeout = 3 * time.Second
			// http.Get("http://www.google.com")
		}()
		time.Sleep(time.Second)

		fmt.Printf("routine count: %d\n", runtime.NumGoroutine()) // 每次+2，应该是执行 http.Get 的协程本身 和 http请求的协程
	}
}

// 占锁后未及时释放
func TestLockWithoutUnlock(t *testing.T) {
	m := sync.Mutex{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Lock()
			if i%2 == 0 {
				defer m.Unlock()
			}
		}(i)

		fmt.Println(runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}

// waitgroup 未全部释放
func TestWaitGroupNotAllDone(t *testing.T) {
	for i := 0; i < 100; i++ {
		wg := sync.WaitGroup{}
		wg.Add(10)

		// done routine
		doneCount := rand.Intn(11)
		go func(d int, wg *sync.WaitGroup) {
			for i := 0; i < d; i++ {
				// business code
				wg.Done() // suggest: defer wg.Done()
			}
		}(doneCount, &wg)

		// wait routine
		go func(wg *sync.WaitGroup) {
			wg.Wait()
		}(&wg)

		fmt.Printf("done count: %d, routine: %d\n", doneCount, runtime.NumGoroutine())
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		time.Sleep(1 * time.Second)
	}
}
