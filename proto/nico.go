package proto

type NicoRequest struct {
	RequestID    string
	ImpressionID string
	DeviceID     string
	PlacementID  string
	UserID       string
}

type NicoResponse struct {
	DspID       int
	RequestKey  string
	BidPrice    float64
	Payload     string
	IsFromCache bool
}
