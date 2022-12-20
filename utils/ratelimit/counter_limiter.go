package ratelimit

import "time"

// CounterSildeWindowLimiter 计数器滑动窗口限流
// 在一段时间区间内，处理请求的最大数量固定，超不部分不做处理
type CounterSildeWindowLimiter struct {
	windowSize int   // 窗口大小
	limit      int   // 窗口时间的最大请求数量
	splitNum   int   // 切分小窗口的数量
	counters []int  // 小窗口数组
	startTime  int64 // 开始时间
	index      int   // 当前小窗口的索引
}

func NewCounterSildeWindowLimiter(windowSize,limit,splitNum int) *CounterSildeWindowLimiter {
	return &CounterSildeWindowLimiter{
		windowSize:windowSize,
		limit: limit,
		splitNum: splitNum,
		counters: make([]int,splitNum),
		startTime: time.Now().UnixNano(),
	}
}
