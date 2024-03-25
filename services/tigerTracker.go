package services

import (
	"encoding/json"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/database"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
	"github.com/vaibhavgvk08/tigerhall-kittens/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type TigerTrackerService struct {
}

var tigerTrackerService *TigerTrackerService

func init() {
	tigerTrackerService = new(TigerTrackerService)
}

func FetchTigerTrackerObject() *TigerTrackerService {
	return tigerTrackerService
}

func (tg *TigerTrackerService) CreateTiger(input model.CreateTigerInput) *model.Tiger {
	doc := bson.M{
		"name":                input.Name,
		"dob":                 input.Dob,
		"lastSeenTimeStamp":   input.LastSeenTimeStamp,
		"lastSeenCoordinates": input.LastSeenCoordinates,
		"imageURL":            input.ImageURL,
	}

	result, err := database.FetchDBManager().Insert(constants.TIGER, doc)
	if err != nil {
		log.Fatal(err)
	}
	var Id string
	if id, success := result.(primitive.ObjectID); success {
		Id = id.Hex()
	}
	//insertedID := result.In.(primitive.ObjectID).String()
	return CreateTiger(Id, input)
}

func (tg *TigerTrackerService) SightTigerLocation(id string, input model.SightingOfTigerInput) *model.Tiger {
	tiger := tg.ListATigers(id)
	recentCoordinates := input.LastSeenCoordinates[0]
	previousCoordinates := tiger.LastSeenCoordinates[0]
	// pre conditions
	dist := utils.Distance(recentCoordinates.Lat, recentCoordinates.Long, previousCoordinates.Lat, previousCoordinates.Long, constants.SIGHTING_DISTANCE_UNITS)
	if dist > constants.SIGHTING_DISTANCE_THRESHOLD {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}
		filter := bson.M{"_id": bson.M{"$eq": objID}}
		lastSeenTimeStampList := append(input.LastSeenTimeStamp, tiger.LastSeenTimeStamp...)
		var latestCoordinates []*model.Coordinates
		cood := &model.Coordinates{
			Lat:  input.LastSeenCoordinates[0].Lat,
			Long: input.LastSeenCoordinates[0].Long,
		}
		latestCoordinates = append(latestCoordinates, cood)
		lastSeenCoordList := append(latestCoordinates, tiger.LastSeenCoordinates...)

		update := bson.D{
			{"$set", bson.D{
				{"lastSeenTimeStamp", lastSeenTimeStampList},
				{"lastSeenCoordinates", lastSeenCoordList},
			}},
		}

		_, err = database.FetchDBManager().Update(constants.TIGER, filter, update)
		if err != nil {
			panic(err)
		}
		return &model.Tiger{
			ID:                  id,
			Name:                tiger.Name,
			Dob:                 tiger.Dob,
			LastSeenTimeStamp:   lastSeenTimeStampList,
			LastSeenCoordinates: nil,
			ImageURL:            &input.ImageURL,
		}
	}
	return nil
}

func (tg *TigerTrackerService) ListAllTigers() []*model.Tiger {
	result, err := database.FetchDBManager().Find(constants.TIGER, bson.D{})
	if err != nil {
		panic(err)
	}
	var tigers []*model.Tiger
	err = json.Unmarshal(result, &tigers)
	if err != nil {
		return nil
	}
	return tigers
}

func (tg *TigerTrackerService) ListATigers(id string) *model.Tiger {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result, err := database.FetchDBManager().Find(constants.TIGER, filter)

	if err != nil {
		panic(err)
	}
	var tigers []*model.Tiger
	err = json.Unmarshal(result, &tigers)
	if err != nil || len(tigers) == 0 {
		return nil
	}
	return tigers[0]
}
