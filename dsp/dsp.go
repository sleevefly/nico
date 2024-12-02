package dsp

import (
	"errors"
	"nico/proto"
)

type Dsp interface {
	ParseDspRequest(req **proto.NicoRequest) error
	GetBidFloor()
	ParseDspResponse()
}

type BaseDsp struct {
}

func NewDsp(dspId int) (Dsp, error) {
	switch dspId {
	case 1:
		return NewTanxDsp(), nil
	case 2:
		return NewAppLinkDsp(), nil
	default:
		return nil, errors.New("unknown exchange ID")
	}
}
