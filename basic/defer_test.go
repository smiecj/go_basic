package basic

import (
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
)

// defer 使用示例
// https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01

type Book struct {
	name string
}

func (book Book) printNameByStruct() {
	log.Info("[book struct] name: %s", book.name)
}

func (book *Book) printNameByPointer() {
	log.Info("[book pointer] name: %s", book.name)
}

// 匿名返回值
func returnWithAnonymous() int {
	ret := 1
	defer func() {
		ret = ret + 1
	}()
	ret = ret + 3
	return ret
}

// 带名称的返回值
func returnWithName() (ret int) {
	ret = 1
	// 注意defer 内部的 ret 依然是和外层一样的对象
	defer func() {
		ret = ret + 1
	}()
	ret = ret + 3
	return ret
}

// 测试: defer 和 return 的先后执行顺序
// return 准备返回值 -> defer -> 返回给上层值
func TestDeferWithReturn(t *testing.T) {
	log.Info("[defer] return with anonymous: %d", returnWithAnonymous()) // 4
	log.Info("[defer] return with variable name: %d", returnWithName())  // 5
}

// 测试: 多个 defer 之间的执行顺序
func TestDeferOrder(t *testing.T) {
	for index := 0; index < 10; index++ {
		defer log.Info("[defer] order: current index: %d", index) // first in last out (stack)
	}
}

// 测试: defer 和 for 循环
// 传入下标参数 和 正确的执行方式
// defer 会在 包裹住的方法(function)结束 才被触发, 参考 defer 的触发时机: https://stackoverflow.com/a/45620423
func TestDeferWithLoop(t *testing.T) {
	func() {
		for index := 0; index < 10; index++ {
			// 注意: 和直接 defer fmt.Print... 这种方式不同，这里 defer 记录的是 i 的地址
			defer func() {
				log.Info("[defer] index with bad order: %d", index) // all 10
			}()
		}
	}()

	func() {
		for index := 0; index < 10; index++ {
			defer func(i int) {
				log.Info("[defer] index with right order: %d", i)
			}(index)
		}
	}()

	func() {
		for index := 0; index < 10; index++ {
			index := index
			defer func() {
				log.Info("[defer] index with right order another solution: %d", index)
			}()
		}
	}()
}

// defer + recover
func TestDeferWithRecover(t *testing.T) {
	jobFinishChan := make(chan struct{})
	go func() {
		defer func() {
			if err := recover(); nil != err {
				log.Info("[defer] exec with err: %s", err.(error).Error())
			} else {
				log.Info("[defer] exec without err")
			}
			close(jobFinishChan)
		}()
		time.Sleep(5 * time.Second)
	}()
	<-jobFinishChan
}

// defer with struct and pointer receiver
// pointer recevier 因为传入的是指针，所以会保留最新的值
// struct receiver 在 执行 defer 的时候已经将整个 struct 值拷贝，所以维持不变
func TestDeferWithReceiver(t *testing.T) {
	book := Book{name: "golang tour"}

	defer book.printNameByStruct()
	defer book.printNameByPointer()

	book.name = "thinking in java"
}

// defer with os.Exit
// https://stackoverflow.com/questions/27629380/how-to-exit-a-go-program-honoring-deferred-calls
// 当前阅读: https://stackoverflow.com/a/28473339
// os.Exit(0) 正常退出时，defer 会正常执行，传3、9都不会执行defer，程序会立即退出，和 panic 的过程很不同
func TestDeferWithExit(t *testing.T) {
	defer log.Info("defer...")

	exitCode := 9
	log.Info("[defer] ready to exit with code: %d", exitCode)
	os.Exit(exitCode)
	runtime.Goexit()
	time.Sleep(5 * time.Second)
}

// 扩展: runtime.Goexit() 和 os.Exit 不同: 停止当前协程，运行所有 defer 方法，等待其他协程正常结束后，进程退出
func TestDeferWithGoexit(t *testing.T) {
	defer log.Info("defer...")

	log.Info("[defer] ready to go exit")
	runtime.Goexit()
	time.Sleep(5 * time.Second)
}

// 其他示例: defer + http 访问和错误判断: https://go.dev/play/p/HwV0gDN3NTE
// 其他示例: defer 和 资源释放的时候对错误的检查: https://go.dev/play/p/PZGn5-1TXv4
