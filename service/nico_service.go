package service

import (
	"context"
	"fmt"
	"net/http"
	"nico/exchange"
	"sync"
)

var (
	defaultAggregatorOnce sync.Once
	defaultAggregator     *DspAggregator
)

func getDefaultAggregator() *DspAggregator {
	defaultAggregatorOnce.Do(func() {
		cfg := LoadDSPRuntimeConfigFromEnv()
		dspIDs := []int{1, 2}
		factory := NewDSPClientFactory()
		if err := factory.RegisterFromConfigs(defaultHTTPCaller, LoadDSPClientConfigsFromEnv(dspIDs)); err != nil {
			fmt.Printf("register dsp routes from env failed: %v\n", err)
		}
		defaultAggregator = NewDspAggregator(NewLateResponseStore(cfg.LateResponseTTL), nil)
	})
	return defaultAggregator
}

func ExchangeService(exchangeId int, req *http.Request) {

	newExchange, err := exchange.NewExchange(exchangeId)
	if err != nil {
		return
	}

	//设置请求的原始request

	//转化为nico的request
	nicoRequest, err := newExchange.ParseNicoRequest()
	if err != nil {
		return
	}

	//设置request

	//获取dsp
	dspIds := []int{1, 2}
	cfg := LoadDSPRuntimeConfigFromEnv()
	aggregator := getDefaultAggregator()
	results := aggregator.AggregateWithCache(context.Background(), nicoRequest, dspIds, cfg.MainWaitTimeout, cfg.SingleDSPTimeout)

	//排序 responses

	// 最终的response

	//请求广告
	fmt.Println(nicoRequest, results)

	err = newExchange.ParseResponse()

}
