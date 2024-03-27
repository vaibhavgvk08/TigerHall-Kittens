package common

import "github.com/vaibhavgvk08/tigerhall-kittens/graph/model"

type EmailData struct {
	To                       string
	TigerName                *string
	TigerLastSeenTimestamp   string
	TigerLastSeenCoordinates *model.Coordinates
}

type TigerDBStruct struct {
	ID                    string               `json:"_id"`
	Name                  string               `json:"name,omitempty"`
	Dob                   string               `json:"dob,omitempty"`
	LastSeenTimeStamp     []string             `json:"lastSeenTimeStamp"`
	LastSeenCoordinates   []*model.Coordinates `json:"lastSeenCoordinates"`
	ImageURL              string               `json:"imageURL,omitempty"`
	ReporterUserNamesList []string             `json:"reporterUserNamesList"`
}

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
