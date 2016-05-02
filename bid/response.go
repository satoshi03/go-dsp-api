package bid

import (
	"github.com/mxmCherry/openrtb"
	"github.com/pborman/uuid"

	"github.com/satoshi03/go-dsp-api/common/consts"
	"github.com/satoshi03/go-dsp-api/data"
)

func makeBidResponse(br *openrtb.BidRequest, ads []*data.Ad) *openrtb.BidResponse {
	sb := makeSeatBid(ads)
	return &openrtb.BidResponse{
		ID:      br.ID,
		SeatBid: sb,
		Cur:     consts.DefaultBidCur,
	}
}

func makeSeatBid(ads []*data.Ad) []openrtb.SeatBid {
	bid := make([]openrtb.Bid, len(ads))
	for i, ad := range ads {
		bid[i] = *makeBid(ad)
	}
	return []openrtb.SeatBid{
		openrtb.SeatBid{
			Bid: bid,
		},
	}
}

func makeBid(ad *data.Ad) *openrtb.Bid {
	return &openrtb.Bid{
		ID:    uuid.NewRandom().String(),
		ImpID: ad.ImpID,
		Price: ad.CalcBidPrice(),
		AdID:  ad.AdID,
		CID:   ad.CampaignID,
		CrID:  ad.CreativeID,
		NURL:  ad.NURL,
		IURL:  ad.IURL,
		AdM:   ad.AdM,
	}
}
