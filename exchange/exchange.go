package exchange

import (
	"errors"
	"nico/proto"
)

// Exchange 接口
type Exchange interface {
	ParseNicoRequest() (*proto.NicoRequest, error)
	ParseResponse() error
}

func NewExchange(exchangeID int) (Exchange, error) {
	switch exchangeID {
	case 1:
		return NewAppLinkExchange(), nil
	case 2:
		return NewNicoExchange(), nil
	default:
		return nil, errors.New("unknown exchange ID")
	}
}
