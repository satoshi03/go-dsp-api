package data

import (
	"github.com/mxmCherry/openrtb"
	"golang.org/x/net/context"
	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/satoshi03/go-dsp-api/common/consts"
	"github.com/satoshi03/go-dsp-api/common/errors"
	"github.com/satoshi03/go-dsp-api/redis"
)

type Index []Ad

const (
	KeyPrefix       = `index`
	BannerRegular   = `banner`
	BannerRectangle = `banner_rect`
	Video           = `video`
	Native          = `native`
)

func getBannerDetailView(w, h uint64) (string, error) {
	// if size is empty, use default
	if w == 0 && h == 0 {
		return BannerRegular, nil
	}
	// In case of regular banner size
	if w == consts.BannerRegularWidth && h == consts.BannerRegularHight {
		return BannerRegular, nil
	}
	// In case of rectangle banner size
	if w == consts.BannerRectangleWidth && h == consts.BannerRectangleHight {
		return BannerRectangle, nil
	}
	// Creative size not match
	return "", errors.InvalidCreativeSizeError
}

func getView(imp *openrtb.Imp) (string, error) {
	switch {
	case imp.Banner != nil:
		return getBannerDetailView(imp.Banner.W, imp.Banner.H)
	case imp.Video != nil:
		return "", errors.NoSupportError{"video"}
	case imp.Native != nil:
		return "", errors.NoSupportError{"native"}
	}
	return "", errors.InvalidViewTypeError
}

func makeKey(imp *openrtb.Imp) (string, error) {
	viewType, err := getView(imp)
	if err != nil {
		return "", err
	}

	// <KeyPrefix>:<View>
	// e.g index:banner
	return KeyPrefix + ":" + viewType, nil
}

func GetIndex(ctx context.Context, imp *openrtb.Imp) (Index, error) {
	// Get key for getting value in redis
	key, err := makeKey(imp)
	if err != nil {
		return nil, err
	}

	// Get value from redis
	cli := redis.GetConn(ctx, consts.CtxRedisKey)
	value, _ := redis.GetCmd(cli, key)
	var out Index
	err = msgpack.Unmarshal([]byte(value), &out)
	if err != nil {
		// Failed to get valid value
		return nil, err
	}

	return out, nil
}
