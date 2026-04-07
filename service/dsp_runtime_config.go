package service

import (
	"os"
	"strconv"
	"time"
)

type DSPRuntimeConfig struct {
	MainWaitTimeout  time.Duration
	SingleDSPTimeout time.Duration
	LateResponseTTL  time.Duration
}

func LoadDSPRuntimeConfigFromEnv() DSPRuntimeConfig {
	return DSPRuntimeConfig{
		MainWaitTimeout:  loadDurationMs("DSP_WAIT_TIMEOUT_MS", DefaultMainWaitTimeout),
		SingleDSPTimeout: loadDurationMs("DSP_SINGLE_TIMEOUT_MS", DefaultSingleDSPTimeout),
		LateResponseTTL:  loadDurationMs("DSP_LATE_RESPONSE_TTL_MS", DefaultLateResponseTTL),
	}
}

func loadDurationMs(key string, fallback time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	ms, err := strconv.Atoi(val)
	if err != nil || ms <= 0 {
		return fallback
	}
	return time.Duration(ms) * time.Millisecond
}
