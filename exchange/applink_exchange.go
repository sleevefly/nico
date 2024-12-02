package exchange

import (
	"fmt"
	"nico/proto"
)

// Exchange2 实现
type AppLinkExchange struct{}

func NewAppLinkExchange() *AppLinkExchange {
	return &AppLinkExchange{}
}

func (e *AppLinkExchange) ParseNicoRequest() (*proto.NicoRequest, error) {
	fmt.Println("Exchange2: Handling Nico Request")
	return nil, nil
}

func (e *AppLinkExchange) ParseResponse() error {
	fmt.Println("Exchange2: Parsing Response")
	return nil
}
