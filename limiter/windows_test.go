package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestWindowsRate(t *testing.T) {
	rate := NewGWindowsRate()
	for i := 0; i < 15; i++ {
		if limiter := rate.GetSlidingWindowRateLimiter("测试"); limiter != nil {
			// 模拟接口访问
			for i := 0; i < 15; i++ {
				if limiter.Allow() {
					fmt.Println("允许访问接口")
				} else {
					fmt.Println("接口访问受限")
				}
				time.Sleep(time.Second)
			}
		} else {
			rateLimiter := NewSlidingWindowRateLimiter("测试", 10, 5, time.Second*10) // 创建一个接口限流器，时间窗口大小为10，最大请求数为5，时间间隔为10秒
			rate.SetSlidingWindowRateLimiter(rateLimiter)
			if rateLimiter.Allow() {
				fmt.Println("允许访问接口")
			} else {
				fmt.Println("接口访问受限")
			}
			time.Sleep(time.Second)
		}
	}

}
