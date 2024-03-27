package common

import (
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
)

func CreateTiger(id string, input model.CreateTigerInput) *model.Tiger {
	tiger := &model.Tiger{
		ID:                id,
		Name:              &input.Name,
		Dob:               &input.Dob,
		LastSeenTimeStamp: input.LastSeenTimeStamp,
		ImageURL:          input.ImageURL,
	}
	tiger.LastSeenCoordinates = &model.Coordinates{
		Lat:  input.LastSeenCoordinates.Lat,
		Long: input.LastSeenCoordinates.Long,
	}
	return tiger
}

func CreateTigerFromDBObject(input *TigerDBStruct) *model.Tiger {
	tiger := &model.Tiger{
		ID:                input.ID,
		Name:              &input.Name,
		Dob:               &input.Dob,
		LastSeenTimeStamp: input.LastSeenTimeStamp[0],
		ImageURL:          &input.ImageURL,
	}
	tiger.LastSeenCoordinates = &model.Coordinates{
		Lat:  input.LastSeenCoordinates[0].Lat,
		Long: input.LastSeenCoordinates[0].Long,
	}
	return tiger
}

func CreateTigerSightings(input *TigerDBStruct, err error) *model.TigerSightings {
	if err != nil {
		return nil
	}
	tigerSightings := &model.TigerSightings{}
	tigerSightings.Name = &input.Name
	tigerSightings.Sightings = make([]*model.Sighting, 0)
	for i := 0; i < len(input.LastSeenTimeStamp); i++ {
		tigerSightings.Sightings = append(tigerSightings.Sightings, &model.Sighting{
			Timestamp:   &input.LastSeenTimeStamp[i],
			Coordinates: input.LastSeenCoordinates[i],
		})
	}
	return tigerSightings
}

func CreateTigersListFromDBObject(input []*TigerDBStruct) []*model.Tiger {
	var tigers []*model.Tiger
	for _, eachTiger := range input {
		tiger := &model.Tiger{
			ID:                eachTiger.ID,
			Name:              &eachTiger.Name,
			Dob:               &eachTiger.Dob,
			LastSeenTimeStamp: eachTiger.LastSeenTimeStamp[0],
			ImageURL:          &eachTiger.ImageURL,
		}
		tiger.LastSeenCoordinates = &model.Coordinates{
			Lat:  eachTiger.LastSeenCoordinates[0].Lat,
			Long: eachTiger.LastSeenCoordinates[0].Long,
		}
		tigers = append(tigers, tiger)
	}
	return tigers
}

func CreateResponse(status, error string) *model.Response {
	resp := new(model.Response)
	resp.Status = status
	resp.Error = &error
	return resp
}

func CreateLoginResponse(status, token, error string) *model.LoginResponse {
	resp := new(model.LoginResponse)
	resp.Status = status
	resp.Error = &error
	resp.AccessToken = token
	return resp
}

func CreateEmailObject(emailList, tigerName, lastSeenTimestamp string, lastSeenCoordinates *model.Coordinates) EmailData {
	return EmailData{
		To:                       emailList,
		TigerName:                &tigerName,
		TigerLastSeenTimestamp:   lastSeenTimestamp,
		TigerLastSeenCoordinates: lastSeenCoordinates,
	}
}
