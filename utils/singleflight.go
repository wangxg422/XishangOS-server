package utils

import "golang.org/x/sync/singleflight"

var SingleFlight *singleflight.Group

func GetSingleFlight() *singleflight.Group {
	if SingleFlight == nil {
		SingleFlight = &singleflight.Group{}
		return SingleFlight
	}

	return SingleFlight
}
