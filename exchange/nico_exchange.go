package exchange

import (
	"fmt"
	"nico/proto"
)

type NicoExchange struct {
}

func NewNicoExchange() *NicoExchange {
	return &NicoExchange{}
}

func (e *NicoExchange) ParseNicoRequest() (*proto.NicoRequest, error) {
	fmt.Println("Exchange1: Handling Nico Request")
	return nil, nil
}

func (e *NicoExchange) ParseResponse() error {
	fmt.Println("Exchange1: Parsing Response")
	return nil
}
