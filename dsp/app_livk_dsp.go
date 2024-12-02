package dsp

import (
	"nico/proto"
)

type AppLinkDsp struct {
	BaseDsp
}

func (a AppLinkDsp) ParseDspRequest(req **proto.NicoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (a AppLinkDsp) GetBidFloor() {
	//TODO implement me
	panic("implement me")
}

func (a AppLinkDsp) ParseDspResponse() {
	//TODO implement me
	panic("implement me")
}

func NewAppLinkDsp() *AppLinkDsp {
	return &AppLinkDsp{}
}
