package limiter

import (
	"sync"
	"time"
)

type GWindowsRate struct {
	WindowsRate map[string]*SlidingWindowRateLimiter
}

// NewGWindowsRate 创建一个全局的滑动窗口管理器
func NewGWindowsRate() *GWindowsRate {
	return &GWindowsRate{
		make(map[string]*SlidingWindowRateLimiter),
	}
}

// SetSlidingWindowRateLimiter 将滑动窗口放进全局管理器中
func (rate *GWindowsRate) SetSlidingWindowRateLimiter(rateLimiter *SlidingWindowRateLimiter) {
	rate.WindowsRate[rateLimiter.name] = rateLimiter
}

// SetSlidingWindowRateLimiterSource 若全局管理器中不存在就创建一个
func (rate *GWindowsRate) SetSlidingWindowRateLimiterSource(name string, windowSize int, maxRequests int, interval time.Duration) {
	rate.WindowsRate[name] = NewSlidingWindowRateLimiter(name, windowSize, maxRequests, interval)
}

// GetSlidingWindowRateLimiter 获取对应名称的滑动窗口
func (rate *GWindowsRate) GetSlidingWindowRateLimiter(name string) *SlidingWindowRateLimiter {
	if val, ok := rate.WindowsRate[name]; ok {
		return val
	}
	return nil
}

func (rate *GWindowsRate) Allow(name string, def bool) bool {
	limiter := rate.GetSlidingWindowRateLimiter(name)
	if limiter == nil {
		if def {
			rate.SetSlidingWindowRateLimiter(NewSlidingWindowRateLimiter(name, 10, 5, time.Second*10))
			return false
		}
		return true
	}
	return limiter.Allow()
}

// SlidingWindowRateLimiter 是一个基于滑动窗口的接口限流器
type SlidingWindowRateLimiter struct {
	name        string
	mu          sync.Mutex
	window      []time.Time   // 时间窗口内的请求时间
	windowSize  int           // 时间窗口的大小
	maxRequests int           // 时间窗口内允许的最大请求数
	interval    time.Duration // 时间窗口的时间间隔
}

// NewSlidingWindowRateLimiter 创建一个新的滑动窗口接口限流器
func NewSlidingWindowRateLimiter(name string, windowSize int, maxRequests int, interval time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		name:        name,
		windowSize:  windowSize,
		maxRequests: maxRequests,
		interval:    interval,
		window:      make([]time.Time, 0, windowSize),
	}
}

// Allow 判断是否允许接口访问
func (s *SlidingWindowRateLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	// 移除时间窗口之前的请求时间
	for len(s.window) > 0 && now.Sub(s.window[0]) > s.interval {
		s.window = s.window[1:]
	}

	// 判断时间窗口内的请求数是否超出限制
	if len(s.window) < s.maxRequests {
		s.window = append(s.window, now)
		return true
	}

	return false
}
