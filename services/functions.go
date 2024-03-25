package services

import "github.com/vaibhavgvk08/tigerhall-kittens/graph/model"

func CreateTiger(id string, input model.CreateTigerInput) *model.Tiger {
	tiger := &model.Tiger{
		ID:                  id,
		Name:                &input.Name,
		Dob:                 &input.Dob,
		LastSeenTimeStamp:   input.LastSeenTimeStamp,
		ImageURL:            input.ImageURL,
		LastSeenCoordinates: make([]*model.Coordinates, 0),
	}
	tiger.LastSeenCoordinates = append(tiger.LastSeenCoordinates, &model.Coordinates{
		Lat:  input.LastSeenCoordinates[0].Lat,
		Long: input.LastSeenCoordinates[0].Long,
	})
	return tiger
}
