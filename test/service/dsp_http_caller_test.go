package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"

	"nico/proto"
	"nico/service"
)

type testCodec struct {
	url string
}

func (c *testCodec) BuildRequest(_ *proto.NicoRequest, requestKey string) (string, string, map[string]string, any, error) {
	return "POST", c.url, map[string]string{"Content-Type": "application/json"}, map[string]any{"request_key": requestKey}, nil
}

func (c *testCodec) ParseResponse(statusCode int, body []byte) (float64, string, error) {
	if statusCode >= 300 {
		return 0, "", fmt.Errorf("bad status: %d", statusCode)
	}
	var out struct {
		BidPrice float64 `json:"bid_price"`
		Payload  string  `json:"payload"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return 0, "", err
	}
	return out.BidPrice, out.Payload, nil
}

func TestHTTPDSPCallerCall(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"bid_price":8.8,"payload":"ok"}`))
	}))
	defer s.Close()

	caller := service.NewHTTPDSPCaller(resty.New())
	caller.RegisterCodec(100, &testCodec{url: s.URL})
	res := caller.Call(context.Background(), 100, &proto.NicoRequest{RequestID: "req"}, "rk")
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
}

func TestHTTPDSPCallerConcurrentCalls(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"bid_price":5.5,"payload":"concurrent"}`))
	}))
	defer s.Close()

	caller := service.NewHTTPDSPCaller(nil)
	caller.RegisterRoute(200, resty.New().SetTimeout(800*time.Millisecond), &testCodec{url: s.URL})

	const n = 80
	var wg sync.WaitGroup
	errCh := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := caller.Call(context.Background(), 200, &proto.NicoRequest{RequestID: "r"}, "k")
			if res.Err != nil {
				errCh <- res.Err
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Fatalf("concurrent call failed: %v", err)
	}
}
