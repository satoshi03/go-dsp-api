package data

import (
	"github.com/satoshi03/go-dsp-api/common/consts"
)

type Ad struct {
	ImpID      string
	CampaignID string
	CreativeID string
	Price      float64
	AdID       string
	NURL       string
	IURL       string
	AdM        string
	Adomain    map[string]interface{}
	PeCPM      float64
}

// Calculate bid price by considering revenue
func (ad *Ad) CalcBidPrice() float64 {
	// Very simple implementation
	// For revenue maximization, need to find the bid price that can win by lowest price
	return ad.PeCPM * (1.0 - consts.RevenueRatio)
}
