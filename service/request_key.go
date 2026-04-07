package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"nico/proto"
)

func RequestKeyFromNicoRequest(req *proto.NicoRequest) string {
	if req == nil {
		return ""
	}
	raw := fmt.Sprintf("%s|%s|%s|%s|%s", req.RequestID, req.ImpressionID, req.DeviceID, req.PlacementID, req.UserID)
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
