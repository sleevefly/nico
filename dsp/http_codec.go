package dsp

import (
	"encoding/json"
	"fmt"
	"strings"

	"nico/proto"
)

type TanxHTTPCodec struct {
	BaseURL string
}

func (c *TanxHTTPCodec) BuildRequest(req *proto.NicoRequest, requestKey string) (string, string, map[string]string, any, error) {
	baseURL := "http://127.0.0.1:18081"
	if c.BaseURL != "" {
		baseURL = c.BaseURL
	}
	body := map[string]any{
		"request_id":    req.RequestID,
		"impression_id": req.ImpressionID,
		"device_id":     req.DeviceID,
		"placement_id":  req.PlacementID,
		"user_id":       req.UserID,
		"request_key":   requestKey,
	}
	return "POST", joinURL(baseURL, "/tanx/bid"), map[string]string{"Content-Type": "application/json"}, body, nil
}

func (c *TanxHTTPCodec) ParseResponse(statusCode int, body []byte) (float64, string, error) {
	return ParseGenericDSPResponse(statusCode, body)
}

type AppLinkHTTPCodec struct {
	BaseURL string
}

func (c *AppLinkHTTPCodec) BuildRequest(req *proto.NicoRequest, requestKey string) (string, string, map[string]string, any, error) {
	baseURL := "http://127.0.0.1:18082"
	if c.BaseURL != "" {
		baseURL = c.BaseURL
	}
	body := map[string]any{
		"request_id":  req.RequestID,
		"device_id":   req.DeviceID,
		"request_key": requestKey,
	}
	return "POST", joinURL(baseURL, "/applink/bid"), map[string]string{"Content-Type": "application/json"}, body, nil
}

func (c *AppLinkHTTPCodec) ParseResponse(statusCode int, body []byte) (float64, string, error) {
	return ParseGenericDSPResponse(statusCode, body)
}

type genericDSPResponse struct {
	BidPrice float64 `json:"bid_price"`
	Payload  string  `json:"payload"`
}

func ParseGenericDSPResponse(statusCode int, body []byte) (float64, string, error) {
	if statusCode >= 300 {
		return 0, "", fmt.Errorf("dsp http status=%d", statusCode)
	}
	var out genericDSPResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return 0, "", err
	}
	return out.BidPrice, out.Payload, nil
}

func joinURL(baseURL, path string) string {
	return strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path, "/")
}
