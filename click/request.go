package click

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/guregu/kami"
	"golang.org/x/net/context"

	"github.com/satoshi03/go-dsp-api/common/errors"
)

type Request struct {
	WonPrice   float64
	CreativeID string
	ImpID      string
}

func parseRequest(ctx context.Context, r *http.Request) *Request {
	wonPriceStr := r.FormValue("price")
	wonPrice, err := strconv.ParseFloat(wonPriceStr, 64)
	if err != nil {
		wonPrice = 0.0
	}

	impID := r.FormValue("impid")
	creativeID := kami.Param(ctx, "crid")

	return &Request{
		WonPrice:   wonPrice,
		CreativeID: creativeID,
		ImpID:      impID,
	}
}

func (r *Request) validate() error {
	if r.WonPrice <= 0.0 {
		return errors.InvalidRequestParamError{"price", fmt.Sprintf("%f", r.WonPrice)}
	}

	if r.ImpID == "" {
		return errors.InvalidRequestParamError{"impid", r.ImpID}
	}

	if r.CreativeID == "" {
		return errors.InvalidRequestParamError{"crid", r.CreativeID}
	}
	// validation ok
	return nil
}
