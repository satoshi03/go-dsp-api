package main

import (
	"github.com/satoshi03/go-dsp-api/bid"
	"github.com/satoshi03/go-dsp-api/click"
	"github.com/satoshi03/go-dsp-api/win"
)

func init() {
	bid.InitHandler()
	win.InitHandler()
	click.InitHandler()
}
