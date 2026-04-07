package exchange_test

import (
	"math"
	"testing"
)

func wilsonScoreCVR(clicks, conversions int, confidence float64) float64 {
	if clicks == 0 {
		return 0.0
	}
	p := float64(conversions) / float64(clicks)
	n := float64(clicks)
	z := confidence
	numerator := p + (z*z)/(2*n) - z*math.Sqrt((p*(1-p)+z*z/(4*n))/n)
	denominator := 1 + (z*z)/n
	return numerator / denominator
}

func TestWilsonScoreCVR(t *testing.T) {
	got := wilsonScoreCVR(200, 1, 1.96)
	if got <= 0 {
		t.Fatalf("expected positive cvr, got %f", got)
	}
}
