package limiter

import (
	"sync"
	"time"
)

// LeakyBucket 是一个基于漏斗算法的限流器
type LeakyBucket struct {
	rate       float64    // 令牌生成速率，单位：令牌/秒
	capacity   float64    // 漏斗容量
	water      float64    // 漏斗中当前令牌数量
	lastLeakMs int64      // 上次漏水时间（毫秒）
	mu         sync.Mutex // 用于保证并发安全
}

// NewLeakyBucket 创建一个新的漏斗限流器
func NewLeakyBucket(rate float64, capacity float64) *LeakyBucket {
	return &LeakyBucket{
		rate:       rate,
		capacity:   capacity,
		water:      0,
		lastLeakMs: time.Now().UnixNano() / int64(time.Millisecond),
	}
}

// Allow 判断是否允许通过，即是否有令牌可用
func (l *LeakyBucket) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	currentMs := time.Now().UnixNano() / int64(time.Millisecond)
	elapsedMs := float64(currentMs - l.lastLeakMs)

	// 计算当前时间段内产生的令牌数量
	leakedWater := elapsedMs / 1000 * l.rate
	l.water -= leakedWater
	if l.water < 0 {
		l.water = 0
	}

	// 更新漏水时间
	l.lastLeakMs = currentMs

	// 增加令牌到漏斗中
	if l.water < l.capacity {
		l.water++
		return true
	}

	return false
}
