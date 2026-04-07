package limiter_test

import (
	"testing"
	"time"

	"nico/limiter"
)

func TestWindowsRateAllow(t *testing.T) {
	rate := limiter.NewGWindowsRate()
	rate.SetSlidingWindowRateLimiterSource("demo", 10, 2, 200*time.Millisecond)
	if !rate.Allow("demo", false) {
		t.Fatalf("first request should be allowed")
	}
}
