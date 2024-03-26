package common

import "github.com/vaibhavgvk08/tigerhall-kittens/graph/model"

type EmailData struct {
	To                       string
	UserName                 string
	TigerName                *string
	TigerLastSeenTimestamp   string
	TigerLastSeenCoordinates *model.Coordinates
}
