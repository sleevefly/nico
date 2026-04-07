package service

import (
	"fmt"
	"nico/dsp"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type DSPClientConfig struct {
	DSPID      int
	BaseURL    string
	Timeout    time.Duration
	RetryCount int
	Proxy      string
}

type DSPClientFactory struct{}

func NewDSPClientFactory() *DSPClientFactory {
	return &DSPClientFactory{}
}

func (f *DSPClientFactory) BuildClient(cfg DSPClientConfig) *resty.Client {
	client := resty.New()
	if cfg.Timeout > 0 {
		client.SetTimeout(cfg.Timeout)
	}
	if cfg.RetryCount > 0 {
		client.SetRetryCount(cfg.RetryCount)
	}
	if cfg.Proxy != "" {
		client.SetProxy(cfg.Proxy)
	}
	return client
}

func (f *DSPClientFactory) BuildCodec(cfg DSPClientConfig) (DSPHTTPCodec, error) {
	switch cfg.DSPID {
	case 1:
		return &dsp.TanxHTTPCodec{BaseURL: cfg.BaseURL}, nil
	case 2:
		return &dsp.AppLinkHTTPCodec{BaseURL: cfg.BaseURL}, nil
	default:
		return nil, fmt.Errorf("unsupported dsp id=%d for auto codec build", cfg.DSPID)
	}
}

func (f *DSPClientFactory) RegisterFromConfigs(caller *HTTPDSPCaller, configs []DSPClientConfig) error {
	for _, cfg := range configs {
		codec, err := f.BuildCodec(cfg)
		if err != nil {
			return err
		}
		client := f.BuildClient(cfg)
		caller.RegisterRoute(cfg.DSPID, client, codec)
	}
	return nil
}

func LoadDSPClientConfigsFromEnv(dspIDs []int) []DSPClientConfig {
	out := make([]DSPClientConfig, 0, len(dspIDs))
	for _, dspID := range dspIDs {
		prefix := fmt.Sprintf("DSP_%d_", dspID)
		baseURL := os.Getenv(prefix + "BASE_URL")
		if baseURL == "" {
			continue
		}
		timeoutMS := envInt(prefix+"TIMEOUT_MS", 3000)
		retry := envInt(prefix+"RETRY_COUNT", 0)
		proxy := os.Getenv(prefix + "PROXY")
		out = append(out, DSPClientConfig{
			DSPID:      dspID,
			BaseURL:    baseURL,
			Timeout:    time.Duration(timeoutMS) * time.Millisecond,
			RetryCount: retry,
			Proxy:      proxy,
		})
	}
	return out
}

func envInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}
