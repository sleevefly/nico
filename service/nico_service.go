package service

import (
	"fmt"
	"net/http"
	"nico/dsp"
	"nico/exchange"
	"nico/proto"
)

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

	dspIds := []int{}
	//请求
	responses := make([]proto.NicoResponse, len(dspIds))
	for _, id := range dspIds {
		newDsp, err := dsp.NewDsp(id)
		if newDsp == nil || err != nil {
			responses = append(responses, proto.NicoResponse{})
			continue
		}
		if err = newDsp.ParseDspRequest(nil); err != nil {
			responses = append(responses, proto.NicoResponse{})
			continue
		}

	}

	//排序 responses

	// 最终的response

	//请求广告
	fmt.Println(nicoRequest)

	err = newExchange.ParseResponse()

}
