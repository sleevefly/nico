package service_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"nico/proto"
	"nico/service"
)

func TestAggregateReturnsWithinWaitWindowAndStoresLateResults(t *testing.T) {
	store := service.NewLateResponseStore(5 * time.Second)
	defer store.Stop()

	caller := func(ctx context.Context, dspID int, _ *proto.NicoRequest, key string) service.DspResult {
		switch dspID {
		case 1:
			time.Sleep(80 * time.Millisecond)
			return service.DspResult{DspID: dspID, RequestKey: key, BidPrice: 10.2}
		case 2:
			time.Sleep(320 * time.Millisecond)
			return service.DspResult{DspID: dspID, RequestKey: key, BidPrice: 20.5}
		default:
			<-ctx.Done()
			return service.DspResult{DspID: dspID, RequestKey: key, Err: ctx.Err()}
		}
	}

	aggr := service.NewDspAggregator(store, caller)
	req := &proto.NicoRequest{RequestID: "r1", ImpressionID: "imp1", DeviceID: "d1"}
	got := aggr.Aggregate(context.Background(), req, []int{1, 2}, 150*time.Millisecond, time.Second)
	if len(got) != 1 {
		t.Fatalf("expected one in-window result, got %d", len(got))
	}
	if got[0].DspID != 1 {
		t.Fatalf("expected dsp 1 in-window, got %d", got[0].DspID)
	}

	time.Sleep(230 * time.Millisecond)
	key := service.RequestKeyFromNicoRequest(req)
	late := store.GetByKey(key)
	if len(late) != 1 || late[0].DspID != 2 {
		t.Fatalf("expected one late result from dsp2, got %+v", late)
	}
}

func TestMergeCachedAndRealtimeWithDedup(t *testing.T) {
	store := service.NewLateResponseStore(5 * time.Second)
	defer store.Stop()

	req := &proto.NicoRequest{RequestID: "r3", ImpressionID: "imp3", DeviceID: "d3"}
	key := service.RequestKeyFromNicoRequest(req)
	store.Put(key, service.DspResult{DspID: 2, RequestKey: key, BidPrice: 2.2})

	caller := func(_ context.Context, dspID int, _ *proto.NicoRequest, key string) service.DspResult {
		switch dspID {
		case 1:
			return service.DspResult{DspID: 1, RequestKey: key, BidPrice: 3.3}
		case 2:
			return service.DspResult{DspID: 2, RequestKey: key, BidPrice: 4.4}
		default:
			return service.DspResult{DspID: dspID, RequestKey: key, Err: errors.New("unknown")}
		}
	}

	aggr := service.NewDspAggregator(store, caller)
	got := aggr.AggregateWithCache(context.Background(), req, []int{1, 2}, 100*time.Millisecond, 100*time.Millisecond)
	if len(got) != 2 {
		t.Fatalf("expected 2 merged unique results, got %d (%+v)", len(got), got)
	}
}

func TestLoadDSPRuntimeConfigFromEnv(t *testing.T) {
	_ = os.Setenv("DSP_WAIT_TIMEOUT_MS", "1500")
	_ = os.Setenv("DSP_SINGLE_TIMEOUT_MS", "2500")
	_ = os.Setenv("DSP_LATE_RESPONSE_TTL_MS", "3500")
	defer os.Unsetenv("DSP_WAIT_TIMEOUT_MS")
	defer os.Unsetenv("DSP_SINGLE_TIMEOUT_MS")
	defer os.Unsetenv("DSP_LATE_RESPONSE_TTL_MS")

	cfg := service.LoadDSPRuntimeConfigFromEnv()
	if cfg.MainWaitTimeout != 1500*time.Millisecond {
		t.Fatalf("unexpected main wait timeout: %v", cfg.MainWaitTimeout)
	}
}
