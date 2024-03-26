package database

import "github.com/vaibhavgvk08/tigerhall-kittens/constants"

type Data struct {
	DBName         string
	CollectionName string
}

var (
	DBMapping = map[int]Data{
		constants.TIGER: {
			DBName:         "TigerDB",
			CollectionName: "TigerCollection",
		},
		constants.USER: {
			DBName:         "UserDB",
			CollectionName: "UserCollection",
		},
	}

	DEFAULT_SORT_ORDER = -1 // Sort in Ascending order.
)
