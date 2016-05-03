package bid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mxmCherry/openrtb"

	"github.com/satoshi03/go-dsp-api/common/consts"
	"github.com/satoshi03/go-dsp-api/common/errors"
)

type Request struct {
	BidRequest *openrtb.BidRequest
}

func parseRequest(r *http.Request) (*Request, error) {
	// Read bid request in json format
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// Unmarshal json to BidRequest
	var br openrtb.BidRequest
	err = json.Unmarshal(body, &br)
	if err != nil {
		return nil, err
	}
	return &Request{
		BidRequest: &br,
	}, nil
}

func (r *Request) validate() error {
	// Required attributes must be exsit
	if r.BidRequest.ID == "" {
		return errors.InvalidRequestParamError{"BidRequest.ID", ""}
	}
	if r.BidRequest.Imp == nil {
		return errors.InvalidRequestParamError{"BidRequest.Imp", ""}
	}

	// Currency type is valid
	if !r.validCurrency() {
		return errors.InvalidCurError
	}
	return nil
}

func (r *Request) validCurrency() bool {
	// currency not exist
	if len(r.BidRequest.Cur) == 0 {
		return true
	}
	// currency exist
	for _, cur := range r.BidRequest.Cur {
		if cur == consts.DefaultBidCur {
			// currency including default bid currency
			return true
		}
	}
	// currency NOT including default bid currency
	return false
}
