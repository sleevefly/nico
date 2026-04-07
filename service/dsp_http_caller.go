package service

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-resty/resty/v2"

	"nico/dsp"
	"nico/proto"
)

type DSPHTTPCodec interface {
	BuildRequest(req *proto.NicoRequest, requestKey string) (method string, url string, headers map[string]string, body any, err error)
	ParseResponse(statusCode int, body []byte) (bidPrice float64, payload string, err error)
}

type DSPHTTPRoute struct {
	Client *resty.Client
	Codec  DSPHTTPCodec
}

type HTTPDSPCaller struct {
	defaultClient *resty.Client
	routes        atomic.Value // map[int]DSPHTTPRoute
}

func NewHTTPDSPCaller(client *resty.Client) *HTTPDSPCaller {
	if client == nil {
		client = resty.New()
	}
	caller := &HTTPDSPCaller{
		defaultClient: client,
	}
	caller.routes.Store(map[int]DSPHTTPRoute{
		1: {Client: client, Codec: &dsp.TanxHTTPCodec{}},
		2: {Client: client, Codec: &dsp.AppLinkHTTPCodec{}},
	})
	return caller
}

func (c *HTTPDSPCaller) RegisterCodec(dspID int, codec DSPHTTPCodec) {
	if codec == nil {
		return
	}
	c.RegisterRoute(dspID, c.defaultClient, codec)
}

func (c *HTTPDSPCaller) RegisterRoute(dspID int, client *resty.Client, codec DSPHTTPCodec) {
	if codec == nil {
		return
	}
	if client == nil {
		client = c.defaultClient
	}
	oldRoutes := c.routes.Load().(map[int]DSPHTTPRoute)
	newRoutes := make(map[int]DSPHTTPRoute, len(oldRoutes)+1)
	for id, route := range oldRoutes {
		newRoutes[id] = route
	}
	newRoutes[dspID] = DSPHTTPRoute{Client: client, Codec: codec}
	c.routes.Store(newRoutes)
}

func (c *HTTPDSPCaller) Call(ctx context.Context, dspID int, req *proto.NicoRequest, requestKey string) DspResult {
	start := time.Now()
	res := DspResult{
		DspID:      dspID,
		RequestKey: requestKey,
		ArrivedAt:  time.Now(),
	}

	routes := c.routes.Load().(map[int]DSPHTTPRoute)
	route, ok := routes[dspID]
	if !ok {
		res.Err = fmt.Errorf("no dsp codec registered for dsp_id=%d", dspID)
		return res
	}
	codec := route.Codec

	method, url, headers, body, err := codec.BuildRequest(req, requestKey)
	if err != nil {
		res.Err = err
		return res
	}
	if url == "" {
		res.Err = errors.New("empty dsp endpoint url")
		return res
	}

	client := route.Client
	if client == nil {
		client = c.defaultClient
	}

	request := client.R().SetContext(ctx).SetHeaders(headers)
	if body != nil {
		request = request.SetBody(body)
	}

	httpResp, err := request.Execute(method, url)
	if err != nil {
		res.Err = err
		res.Latency = time.Since(start)
		res.ArrivedAt = time.Now()
		return res
	}

	bidPrice, payload, err := codec.ParseResponse(httpResp.StatusCode(), httpResp.Body())
	res.BidPrice = bidPrice
	res.Payload = payload
	res.Err = err
	res.Latency = time.Since(start)
	res.ArrivedAt = time.Now()
	return res
}
