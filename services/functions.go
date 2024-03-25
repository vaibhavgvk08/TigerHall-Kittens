package services

import "github.com/vaibhavgvk08/tigerhall-kittens/graph/model"

func CreateTiger(id string, input model.CreateTigerInput) *model.Tiger {
	return &model.Tiger{
		ID:                id,
		Name:              &input.Name,
		Dob:               &input.Dob,
		LastSeenTimeStamp: input.LastSeenTimeStamp,
		LastSeenCoordinates: &model.Coordinates{
			Lat:  input.LastSeenCoordinates.Lat,
			Long: input.LastSeenCoordinates.Long,
		},
		ImageURL: input.ImageURL,
	}
}
