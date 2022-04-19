// package prometheus_ prometheus 示例
package prometheus_

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/smiecj/go_common/util/log"
)

const (
	serverPort = 2212

	// exporter 记录时间周期
	exporterRefreshInterval     = 15 * time.Second
	exporterRefreshIntervalLong = time.Minute
)

// exporter 服务定义
type exporter struct {
	gauge     *prometheus.GaugeVec
	count     *prometheus.CounterVec
	histogram *prometheus.HistogramVec
	summary   *prometheus.SummaryVec
	ctx       context.Context
}

// 启动 exporter
func (exporter *exporter) Start() {
	go func() {
		exporter.startGauge()
		exporter.startCount()
		exporter.startHistogram()
		exporter.startSummary()
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	}()
	log.Info("[exporter] start success, port: %d", serverPort)
}

// 开启 gauge 记录
func (exporter *exporter) startGauge() {
	go func() {
		ticker := time.NewTicker(exporterRefreshInterval)
		for {
			select {
			case <-ticker.C:
				exporter.gauge.With(prometheus.Labels{"type": "cpu"}).Set(rand.Float64() * 100)
				exporter.gauge.With(prometheus.Labels{"type": "memory"}).Set(rand.Float64() * 100)
			case <-exporter.ctx.Done():
				exporter.gauge.Reset()
				return
			}
		}
	}()
}

func (exporter *exporter) startCount() {
	go func() {
		ticker := time.NewTicker(exporterRefreshInterval)
		for {
			select {
			case <-ticker.C:
				// 查询接口调用比较多，每次新增10-20次
				exporter.count.With(prometheus.Labels{"url": "/api/search", "ret": "200"}).Add(10 + math.Floor(10*(rand.Float64())))
				// 新增接口调用较少，每次新增0-10次
				exporter.count.With(prometheus.Labels{"url": "/api/add", "ret": "200"}).Add(math.Floor(10 * (rand.Float64())))
				// 查询接口偶尔因为网络问题会报错，每次新增0-5次
				exporter.count.With(prometheus.Labels{"url": "/api/search", "ret": "500"}).Add(math.Floor(5 * (rand.Float64())))
			case <-exporter.ctx.Done():
				exporter.count.Reset()
				return
			}
		}
	}()
}

func (exporter *exporter) startHistogram() {
	go func() {
		ticker := time.NewTicker(exporterRefreshInterval)
		for {
			select {
			case <-ticker.C:
				// 接口的正常耗时: 设置为1-10s范围内
				exporter.histogram.With(prometheus.Labels{"l_url": "/api/search", "l_ret": "200"}).Observe(math.Ceil(10 * (rand.Float64())))
			case <-exporter.ctx.Done():
				exporter.histogram.Reset()
				return
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(exporterRefreshIntervalLong)
		for {
			select {
			case <-ticker.C:
				// 异常接口耗时: 60~120s
				exporter.histogram.With(prometheus.Labels{"l_url": "/api/search", "l_ret": "200"}).Observe(60 + math.Ceil(60*rand.Float64()))
			case <-exporter.ctx.Done():
				exporter.histogram.Reset()
				return
			}
		}
	}()
}

func (exporter *exporter) startSummary() {
	go func() {
		ticker := time.NewTicker(exporterRefreshInterval)
		for {
			select {
			case <-ticker.C:
				// 接口的正常耗时: 设置为1-10s范围内
				exporter.summary.With(prometheus.Labels{"url": "/api/search", "ret": "200"}).Observe(math.Ceil(10 * (rand.Float64())))
			case <-exporter.ctx.Done():
				exporter.summary.Reset()
				return
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(exporterRefreshIntervalLong)
		for {
			select {
			case <-ticker.C:
				// 异常接口耗时: 60~120s
				exporter.summary.With(prometheus.Labels{"url": "/api/search", "ret": "200"}).Observe(60 + math.Ceil(60*rand.Float64()))
			case <-exporter.ctx.Done():
				exporter.summary.Reset()
				return
			}
		}
	}()
}

// 创建 exporter
func NewExporer(ctx context.Context) *exporter {
	exporter := new(exporter)
	exporter.gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "machine_status",
		Help: "机器状态"}, []string{"type"})

	exporter.count = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "api_status",
		Help: "接口调用情况",
	}, []string{"url", "ret"})

	exporter.histogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_timecost_histogram",
		Help:    "接口耗时分位统计",
		Buckets: []float64{0.01, 0.1, 1, 2, 5, 10, 30},
	}, []string{"l_url", "l_ret"})

	exporter.summary = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "api_summary",
		Help: "生产环境的接口耗时汇总",
		// Objectives 和众数的统计相关，两个参数都是百分比
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"url", "ret"})

	exporter.ctx = ctx
	return exporter
}
