package win

import (
	"log"
	"net/http"

	"github.com/guregu/kami"
	"golang.org/x/net/context"

	"github.com/satoshi03/go-dsp-api/common/utils"
	"github.com/satoshi03/go-dsp-api/fluent"
)

func winHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Parse win notice
	request := parseRequest(ctx, r)

	// Validate Request
	if err := request.validate(); err != nil {
		// Although error request is not valid, send ok response and log request
		log.Println(err)
	}

	// Send Response
	// No send Ad markup since it is included in bid response
	utils.WriteResponse(w, map[string]interface{}{"message": "ok"}, 200)

	// Send log
	fluent.Send(ctx, "fluent", "win", map[string]interface{}{
		"WonPrice":   request.WonPrice,
		"CreativeID": request.CreativeID,
		"ImpID":      request.ImpID,
	})
}

func InitHandler() {
	kami.Get("/v1/win/:crid", winHandler)
}
