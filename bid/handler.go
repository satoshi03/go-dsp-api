package bid

import (
	"log"
	"net/http"

	"github.com/guregu/kami"
	"github.com/mxmCherry/openrtb"
	"golang.org/x/net/context"

	"github.com/satoshi03/go-dsp-api/common/consts"
	"github.com/satoshi03/go-dsp-api/common/utils"
	"github.com/satoshi03/go-dsp-api/fluent"
)

func bidHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Parse bid request
	request, err := parseRequest(r)
	if err != nil {
		log.Println(err)
		noBidResponse(w)
		return
	}

	// Validate request
	if err := request.validate(); err != nil {
		log.Println(err)
		noBidResponse(w)
		return
	}

	// Bidding
	ads := bid(ctx, request.BidRequest)
	if len(ads) == 0 {
		noBidResponse(w)
		return
	}

	// Make Response
	resp := makeBidResponse(request.BidRequest, ads)

	// Send Response
	utils.WriteResponse(w, resp, 200)

	// Send log
	sendLog(ctx, request.BidRequest, resp)
}

func sendLog(ctx context.Context, req *openrtb.BidRequest, resp *openrtb.BidResponse) {
	fluent.Send(ctx, consts.CtxFluentKey, "request", makeRequestLog(req, resp))
	for _, bid := range resp.SeatBid[0].Bid {
		fluent.Send(ctx, consts.CtxFluentKey, "bid", makeBidLog(&bid))
	}
}

func noBidResponse(w http.ResponseWriter) {
	utils.WriteResponse(w, nil, 204)
}

func InitHandler() {
	kami.Post("/v1/bid", bidHandler)
}
