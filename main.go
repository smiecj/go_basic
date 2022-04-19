package main

import (
	"context"

	"github.com/smiecj/go_basic/prometheus_"
)

// 主函数
func main() {
	exporter := prometheus_.NewExporer(context.Background())
	exporter.Start()

	// 接口服务: 永久不退出
	ch := make(chan struct{})
	<-ch
}

