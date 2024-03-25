package services

import (
	"encoding/json"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/database"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
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

func (tg *TigerTrackerService) sightTigerLocation(input model.CreateTigerInput) {

}

type Tiger struct {
	ID                  string       `json:"_id"`
	Name                string       `json:"name,omitempty"`
	Dob                 string       `json:"dob,omitempty"`
	LastSeenTimeStamp   string       `json:"lastSeenTimeStamp"`
	LastSeenCoordinates *Coordinates `json:"lastSeenCoordinates"`
	ImageURL            string       `json:"imageURL,omitempty"`
}

type Coordinates struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
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

func (tg *TigerTrackerService) ListATigers(input model.CreateTigerInput) {

}
