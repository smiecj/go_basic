// package zookeeper_ zk 客户端测试
package zookeeper_

import (
	"testing"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/smiecj/go_common/util/log"
)

const (
	// server = "172.17.0.1:12181"
	server = "common1.lls.com:2181"
)

func TestConnectServer(t *testing.T) {
	c, _, err := zk.Connect([]string{server}, time.Second)
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
