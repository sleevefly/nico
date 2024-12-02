package dsp

import (
	"nico/proto"
)

type TanxDsp struct {
	BaseDsp
	RawReqBody []byte
}

func NewTanxDsp() *TanxDsp {
	return &TanxDsp{}
}

func (t TanxDsp) ParseDspRequest(req **proto.NicoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (t TanxDsp) GetBidFloor() {
	//TODO implement me
	panic("implement me")
}

func (t TanxDsp) ParseDspResponse() {
	//TODO implement me
	panic("implement me")
}
