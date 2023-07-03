// package zookeeper_ zk 客户端测试
package zookeeper_

import (
	"flag"
	"testing"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/smiecj/go_common/util/log"
)

var (
	zkAddress *string
)

func init() {
	zkAddress = flag.String("zk", "localhost:2181", "zookeeper server address")
}

func TestConnectServer(t *testing.T) {
	flag.Parse()

	c, _, err := zk.Connect([]string{*zkAddress}, time.Second)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	log.Info("%+v %+v", children, stat)
	go func() {
		e := <-ch
		log.Info("%+v	", e)
	}()
}
