package services

import (
	"encoding/json"
	"fmt"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/database"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func RegisterUser(request model.CreateUserInput) *model.Response {
	//response.Header().Set("Content-Type", "application/json")
	doc := bson.M{
		"username": request.Username,
		"email":    request.Email,
		"password": auth.GetHash([]byte(request.Password)),
	}
	result, err := database.FetchDBManager().Insert(constants.USER, doc)
	if err != nil {
		log.Fatal(err)
		errString := fmt.Sprint(err)
		return &model.Response{
			Status: "failure",
			Error:  &errString,
		}
	}
	if id, success := result.(primitive.ObjectID); success {
		log.Println("Success :: User registered ", id)
	}
	return &model.Response{
		Status: "Success",
		Error:  nil,
	}
}

func LoginUser(input model.LoginUserInput) *model.LoginResponse {
	filter := bson.M{"username": bson.M{"$eq": input.Username}}
	result, err := database.FetchDBManager().Find(constants.USER, filter)

	if err != nil {
		panic(err)
	}
	var userInDB []*model.User
	err = json.Unmarshal(result, &userInDB)
	if err != nil || len(userInDB) == 0 {
		panic(err)
	}

	userPass := []byte(input.Password)
	dbPass := []byte(userInDB[0].Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		errString := fmt.Sprint(passErr)
		return &model.LoginResponse{
			Status: "fail",
			Error:  &errString,
		}
	}
	jwtToken, err := auth.GenerateJWT(input.Username)
	if err != nil {
		log.Println(err)
		errString := fmt.Sprint(err)
		return &model.LoginResponse{
			Status: "fail",
			Error:  &errString,
		}
	}
	fmt.Println(jwtToken)
	return &model.LoginResponse{
		Status:      "success",
		Error:       nil,
		AccessToken: jwtToken,
	}
}
