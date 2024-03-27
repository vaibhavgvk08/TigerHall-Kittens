package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/database"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/auth"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func RegisterUser(request model.CreateUserInput) (*model.Response, error) {
	doc := bson.M{
		"username": request.Username,
		"email":    request.Email,
		"password": auth.GetHash([]byte(request.Password)),
	}
	result, err := database.FetchDBManager().Insert(constants.USER, doc)
	if err != nil {
		errString := fmt.Sprint(err)
		return common.CreateResponse(constants.FAILURE, errString), err
	}
	if id, success := result.(primitive.ObjectID); success {
		log.Println("Success :: User registered ", id)
	}
	return common.CreateResponse(constants.SUCCESS, ""), nil
}

func FetchUser(input model.LoginUserInput) ([]*model.User, error) {
	filter := bson.M{"username": bson.M{"$eq": input.Username}}
	result, err := database.FetchDBManager().Find(constants.USER, filter, database.DEFAULT_SORT_ORDER, 0, 0)

	if err != nil {
		panic(err)
	}
	var userInDB []*model.User
	err = json.Unmarshal(result, &userInDB)
	return userInDB, err
}

func LoginUser(input model.LoginUserInput) (*model.LoginResponse, error) {
	userInDB, err := FetchUser(input)
	if err != nil {
		return common.CreateLoginResponse(constants.FAILURE, "", "unmarshall error"), err
	} else if len(userInDB) == 0 {
		return common.CreateLoginResponse(constants.FAILURE, "", "User not registered or found in DB"), err
	}

	userPass := []byte(input.Password)
	dbPass := []byte(userInDB[0].Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		errString := fmt.Sprint(passErr)
		return common.CreateLoginResponse(constants.FAILURE, "", errString), passErr
	}
	jwtToken, err := auth.GenerateJWT(input.Username)
	if err != nil {
		log.Println(err)
		errString := fmt.Sprint(err)
		return common.CreateLoginResponse(constants.FAILURE, "", errString), err
	}

	return common.CreateLoginResponse(constants.SUCCESS, jwtToken, ""), nil
}

func FetchUsersEmails(usernames []string) ([]string, error) {
	filter := bson.M{"username": bson.M{"$in": usernames}}
	result, err := database.FetchDBManager().Find(constants.USER, filter, database.DEFAULT_SORT_ORDER, 0, 0)

	if err != nil {
		return nil, err
	}
	var userInDB []*model.User
	if err := json.Unmarshal(result, &userInDB); err != nil {
		return nil, err
	} else if len(userInDB) == 0 {
		return nil, errors.New("one or more invalid usernames provided")
	}

	var emailList []string
	for _, user := range userInDB {
		emailList = append(emailList, user.Email)
	}
	return emailList, nil
}
