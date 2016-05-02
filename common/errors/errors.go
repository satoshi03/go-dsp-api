package errors

import (
	"errors"
)

var (
	// For bidding validation
	InvalidCurError          = errors.New("Bid currency is not valid")
	InvalidCreativeSizeError = errors.New("Creative size is invalid")
	InvalidViewTypeError     = errors.New("View type is invalid")
	LowPriceError            = errors.New("Bid Price is lower than floor price")

	// For redis access
	RedisKeyCreateError = errors.New("Failed to create redis key")
)

type InvalidRequestParamError struct {
	Param string
	Value string
}

func (e InvalidRequestParamError) Error() string {
	return "Parameter is not valid. Param: " + e.Param + " Value: " + e.Value
}

type NoSupportError struct {
	NSField string
}

func (e NoSupportError) Error() string {
	return e.NSField + " is not supported"
}
