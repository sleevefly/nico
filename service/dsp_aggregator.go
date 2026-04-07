package service

import (
	"context"
	"log"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"nico/proto"
)

const (
	DefaultMainWaitTimeout  = 3 * time.Second
	DefaultSingleDSPTimeout = 3 * time.Second
	DefaultLateResponseTTL  = 30 * time.Second
)

type DspResult struct {
	DspID       int
	RequestKey  string
	BidPrice    float64
	Payload     string
	Err         error
	Latency     time.Duration
	IsFromCache bool
	ArrivedAt   time.Time
}

type DSPCaller func(ctx context.Context, dspID int, req *proto.NicoRequest, requestKey string) DspResult

type DspAggregator struct {
	store  *LateResponseStore
	caller DSPCaller
}

var defaultHTTPCaller = NewHTTPDSPCaller(nil)

func NewDspAggregator(store *LateResponseStore, caller DSPCaller) *DspAggregator {
	if store == nil {
		store = NewLateResponseStore(DefaultLateResponseTTL)
	}
	if caller == nil {
		caller = defaultDSPCaller
	}
	return &DspAggregator{store: store, caller: caller}
}

func (a *DspAggregator) Aggregate(ctx context.Context, req *proto.NicoRequest, dspIDs []int, mainWaitTimeout, singleDSPTimeout time.Duration) []DspResult {
	return a.aggregateInternal(ctx, req, dspIDs, mainWaitTimeout, singleDSPTimeout, false)
}

func (a *DspAggregator) AggregateWithCache(ctx context.Context, req *proto.NicoRequest, dspIDs []int, mainWaitTimeout, singleDSPTimeout time.Duration) []DspResult {
	return a.aggregateInternal(ctx, req, dspIDs, mainWaitTimeout, singleDSPTimeout, true)
}

func (a *DspAggregator) aggregateInternal(ctx context.Context, req *proto.NicoRequest, dspIDs []int, mainWaitTimeout, singleDSPTimeout time.Duration, withCache bool) []DspResult {
	if mainWaitTimeout <= 0 {
		mainWaitTimeout = DefaultMainWaitTimeout
	}
	if singleDSPTimeout <= 0 {
		singleDSPTimeout = DefaultSingleDSPTimeout
	}

	requestKey := RequestKeyFromNicoRequest(req)
	merged := make([]DspResult, 0, len(dspIDs))
	if withCache {
		merged = append(merged, a.store.GetByKey(requestKey)...)
	}

	startAt := time.Now()
	cutoff := startAt.Add(mainWaitTimeout)

	resultsCh := make(chan DspResult, len(dspIDs))
	var wg sync.WaitGroup
	for _, id := range dspIDs {
		wg.Add(1)
		go func(dspID int) {
			defer wg.Done()
			start := time.Now()
			dspCtx, cancel := context.WithTimeout(ctx, singleDSPTimeout)
			defer cancel()

			res := a.caller(dspCtx, dspID, req, requestKey)
			if res.ArrivedAt.IsZero() {
				res.ArrivedAt = time.Now()
			}
			if res.Latency <= 0 {
				res.Latency = time.Since(start)
			}
			if res.RequestKey == "" {
				res.RequestKey = requestKey
			}
			if res.DspID == 0 {
				res.DspID = dspID
			}
			resultsCh <- res
		}(id)
	}

	var mu sync.Mutex
	mainWindowResults := make([]DspResult, 0, len(dspIDs))
	var lateCount atomic.Int64
	collectorDone := make(chan struct{})
	go func() {
		defer close(collectorDone)
		for res := range resultsCh {
			if res.Err != nil {
				continue
			}
			if !res.ArrivedAt.After(cutoff) {
				mu.Lock()
				mainWindowResults = append(mainWindowResults, res)
				mu.Unlock()
				continue
			}
			a.store.Put(requestKey, res)
			lateCount.Add(1)
		}
	}()

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	mainTimer := time.NewTimer(time.Until(cutoff))
	defer mainTimer.Stop()

	select {
	case <-collectorDone:
		// 所有DSP已返回，提前结束。
	case <-mainTimer.C:
		// 主等待窗口到期，后续迟到结果由collector继续回收存储。
	}

	mu.Lock()
	merged = append(merged, mainWindowResults...)
	inWindowCount := len(mainWindowResults)
	mu.Unlock()
	log.Printf("dsp_aggregate request_key=%s dsp_total=%d in_window=%d late_stored=%d cache_hit=%d", requestKey, len(dspIDs), inWindowCount, lateCount.Load(), len(merged)-inWindowCount)
	return mergeAndSort(merged)
}

func mergeAndSort(results []DspResult) []DspResult {
	latestByDSP := make(map[int]DspResult, len(results))
	for _, r := range results {
		prev, ok := latestByDSP[r.DspID]
		if !ok || (!r.IsFromCache && prev.IsFromCache) || (r.ArrivedAt.After(prev.ArrivedAt)) {
			latestByDSP[r.DspID] = r
		}
	}
	merged := make([]DspResult, 0, len(latestByDSP))
	for _, r := range latestByDSP {
		merged = append(merged, r)
	}
	sort.Slice(merged, func(i, j int) bool {
		if merged[i].BidPrice == merged[j].BidPrice {
			return merged[i].DspID < merged[j].DspID
		}
		return merged[i].BidPrice > merged[j].BidPrice
	})
	return merged
}

func defaultDSPCaller(ctx context.Context, dspID int, req *proto.NicoRequest, requestKey string) DspResult {
	return defaultHTTPCaller.Call(ctx, dspID, req, requestKey)
}
