package bid

import (
	"log"

	"github.com/mxmCherry/openrtb"
	"golang.org/x/net/context"

	"github.com/satoshi03/go-dsp-api/common/consts"
	"github.com/satoshi03/go-dsp-api/common/errors"
	"github.com/satoshi03/go-dsp-api/data"
)

func bid(ctx context.Context, br *openrtb.BidRequest) []*data.Ad {
	var selected []*data.Ad
	for _, imp := range br.Imp {
		// If imp is invalid, skip it
		if err := validateImp(&imp); err != nil {
			continue
		}
		ad, ok := getAd(ctx, &imp)
		if ok {
			ad.ImpID = imp.ID
			selected = append(selected, &ad)
		}
	}
	return selected
}

func getAd(ctx context.Context, imp *openrtb.Imp) (data.Ad, bool) {
	// Get Index having candidate ad list
	index, err := data.GetIndex(ctx, imp)
	if err != nil {
		// For debug
		log.Println(err)
		return data.Ad{}, false
	}

	// Find valid ad having max score(=ecpm) in index
	var validAds []data.Ad
	for i := range index {
		if err := validateAd(imp, &index[i]); err == nil {
			validAds = append(validAds, index[i])
		} else {
			// For debug
			log.Println(err)
		}
	}
	// Choice high revenue ad in valid ads
	return choiceBestAd(validAds)
}

func choiceBestAd(va []data.Ad) (data.Ad, bool) {
	maxBidPrice := 0.0
	index := -1
	for i := range va {
		bp := va[i].CalcBidPrice()
		if bp > maxBidPrice {
			maxBidPrice = bp
			index = i
		}
	}
	if index == -1 {
		return data.Ad{}, false
	}
	return va[index], true
}

func validateImp(imp *openrtb.Imp) error {
	// Native Ad is not supported
	if imp.Native != nil {
		return errors.NoSupportError{"native"}
	}
	// Video Ad is not supported
	if imp.Video != nil {
		return errors.NoSupportError{"video"}
	}
	// Check bid currency
	if imp.BidFloorCur != "" && imp.BidFloorCur != consts.DefaultBidCur {
		return errors.InvalidCurError
	}
	return nil
}

func validateAd(imp *openrtb.Imp, ad *data.Ad) error {
	// Check bid price greater than bid floor price
	if ad.CalcBidPrice() <= imp.BidFloor {
		return errors.LowPriceError
	}
	return nil
}
