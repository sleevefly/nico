package service_test

import (
	"os"
	"testing"
	"time"

	"nico/service"
)

func TestLoadDSPClientConfigsFromEnv(t *testing.T) {
	_ = os.Setenv("DSP_1_BASE_URL", "http://127.0.0.1:19001")
	_ = os.Setenv("DSP_1_TIMEOUT_MS", "1500")
	_ = os.Setenv("DSP_1_RETRY_COUNT", "2")
	defer os.Unsetenv("DSP_1_BASE_URL")
	defer os.Unsetenv("DSP_1_TIMEOUT_MS")
	defer os.Unsetenv("DSP_1_RETRY_COUNT")

	configs := service.LoadDSPClientConfigsFromEnv([]int{1, 2})
	if len(configs) != 1 {
		t.Fatalf("expected one config, got %d", len(configs))
	}
	if configs[0].Timeout != 1500*time.Millisecond {
		t.Fatalf("unexpected timeout: %v", configs[0].Timeout)
	}
}

func TestDSPClientFactoryRegisterFromConfigs(t *testing.T) {
	caller := service.NewHTTPDSPCaller(nil)
	factory := service.NewDSPClientFactory()
	err := factory.RegisterFromConfigs(caller, []service.DSPClientConfig{
		{DSPID: 1, BaseURL: "http://127.0.0.1:19001", Timeout: time.Second, RetryCount: 1},
		{DSPID: 2, BaseURL: "http://127.0.0.1:19002", Timeout: 2 * time.Second, RetryCount: 2},
	})
	if err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}
}
