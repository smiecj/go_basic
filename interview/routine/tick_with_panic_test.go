package routine

import (
	"testing"
	"time"

	"github.com/smiecj/go_common/util/log"
)

// https://github.com/lifei6671/interview-go/blob/master/question/q012.md
// 每秒调用一次 proc 方法，且不退出
// defer + recover
func TestTickWithPanic(t *testing.T) {
	go func() {

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			go func() {
				defer func() {
					if err := recover(); nil != err {
						log.Info("[recover] err: %s", err.(string))
					}
				}()
				proc()
			}()
		}
	}()

	infiniteChan := make(chan int)
	<-infiniteChan
}

func proc() {
	panic("ok")
}
