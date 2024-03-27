package services

import (
	"encoding/json"
	"errors"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/database"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
	"github.com/vaibhavgvk08/tigerhall-kittens/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (tg *TigerTrackerService) CreateTiger(input model.CreateTigerInput) (*model.Tiger, error) {

	doc := bson.M{
		"name":                  input.Name,
		"dob":                   input.Dob,
		"lastSeenTimeStamp":     []string{input.LastSeenTimeStamp},
		"lastSeenCoordinates":   []model.Coordinates{model.Coordinates{input.LastSeenCoordinates.Lat, input.LastSeenCoordinates.Long}},
		"imageURL":              input.ImageURL,
		"reporterUserNamesList": []string{input.ReporterUserName},
	}

	result, err := database.FetchDBManager().Insert(constants.TIGER, doc)
	if err != nil {
		return nil, errors.New("error in inserting data into DB")
	}
	var Id string
	if id, success := result.(primitive.ObjectID); success {
		Id = id.Hex()
	}

	return common.CreateTiger(Id, input), nil
}

func (tg *TigerTrackerService) SightTigerLocation(id string, input model.SightingOfTigerInput) (*model.Tiger, error) {
	tigerDBObject, err := tg.FetchTigerFromDB(id)
	if err != nil {
		return nil, err
	}
	recentCoordinates := input.LastSeenCoordinates
	previousCoordinates := tigerDBObject.LastSeenCoordinates[0]
	// pre conditions
	dist := utils.Distance(recentCoordinates.Lat, recentCoordinates.Long, previousCoordinates.Lat, previousCoordinates.Long, constants.SIGHTING_DISTANCE_UNITS)
	if dist > constants.SIGHTING_DISTANCE_THRESHOLD {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}
		filter := bson.M{"_id": bson.M{"$eq": objID}}
		lastSeenTimeStampList := append([]string{input.LastSeenTimeStamp}, tigerDBObject.LastSeenTimeStamp...)
		var latestCoordinates []*model.Coordinates
		cood := &model.Coordinates{
			Lat:  input.LastSeenCoordinates.Lat,
			Long: input.LastSeenCoordinates.Long,
		}
		latestCoordinates = append(latestCoordinates, cood)
		lastSeenCoordList := append(latestCoordinates, tigerDBObject.LastSeenCoordinates...)
		// Removing duplicates from usernames of people who sighted a tiger.
		usersWhoSightedTigerList := utils.MergeListsWithOutDuplicates([]string{input.ReporterUserName}, tigerDBObject.ReporterUserNamesList)

		update := bson.D{
			{"$set", bson.D{
				{"lastSeenTimeStamp", lastSeenTimeStampList},
				{"lastSeenCoordinates", lastSeenCoordList},
				{"reporterUserNamesList", usersWhoSightedTigerList},
			}},
		}

		if _, err = database.FetchDBManager().Update(constants.TIGER, filter, update); err != nil {
			return nil, err
		}
		tigerDBObj := &common.TigerDBStruct{
			ID:                    id,
			Name:                  tigerDBObject.Name,
			Dob:                   tigerDBObject.Dob,
			LastSeenTimeStamp:     lastSeenTimeStampList,
			LastSeenCoordinates:   lastSeenCoordList,
			ImageURL:              input.ImageURL,
			ReporterUserNamesList: usersWhoSightedTigerList,
		}
		PrepareAndSendEmails(tigerDBObj, input.ReporterUserName)
		return common.CreateTigerFromDBObject(tigerDBObj), nil
	}
	return nil, errors.New("sighting is not possible if a tiger is not within 5 kilometres of its previous sighting")
}

func (tg *TigerTrackerService) FetchAllTigersFromDB(input *model.InputParams) ([]*model.Tiger, error) {
	result, err := database.FetchDBManager().Find(constants.TIGER, bson.D{}, database.DEFAULT_SORT_ORDER, *input.Offset, *input.Limit)
	if err != nil {
		panic(err)
	}
	var tigers []*common.TigerDBStruct
	err = json.Unmarshal(result, &tigers)
	if err != nil {
		return nil, err
	}
	return common.CreateTigersListFromDBObject(tigers), nil
}

func (tg *TigerTrackerService) FetchTigerFromDB(id string) (*common.TigerDBStruct, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result, err := database.FetchDBManager().Find(constants.TIGER, filter, database.DEFAULT_SORT_ORDER, 0, 0)
	if err == nil {
		var tigersDBObject []*common.TigerDBStruct
		err = json.Unmarshal(result, &tigersDBObject)
		if len(tigersDBObject) == 0 {
			return nil, errors.New("provided tiger Id is invalid")
		}
		if err != nil {
			return nil, errors.New("internal system error :: error in unmarshalling tiger object")
		}
		return tigersDBObject[0], nil
	}
	return nil, err
}
