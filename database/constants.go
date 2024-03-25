package database

import "github.com/vaibhavgvk08/tigerhall-kittens/constants"

type Data struct {
	DBName         string
	CollectionName string
}

var (
	DBMapping = map[int]Data{
		constants.TIGER: Data{
			DBName:         "TigerDB",
			CollectionName: "TigerCollection",
		},
	}
)
