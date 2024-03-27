package utils

import (
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"math"
	"time"
)

func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func RemoveAItemFromList(list1 []string, currentElem string) []string {
	var newlist []string
	for _, elem := range list1 {
		if elem != currentElem {
			newlist = append(newlist, elem)
		}
	}
	return newlist
}

func MergeListsWithOutDuplicates(list1, list2 []string) []string {
	// Create a map to store unique elements
	uniqueMap := make(map[string]bool)

	// Add elements from list1 to the map
	for _, item := range list1 {
		uniqueMap[item] = true
	}

	// Add elements from list2 to the map
	for _, item := range list2 {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = true
		}
	}

	// Extract unique elements from the map into a slice
	var mergedList []string
	for item := range uniqueMap {
		mergedList = append(mergedList, item)
	}

	return mergedList
}

func IsCurrentTimeStampLatest(currentTimestampStr, previousTimestampStr string) (bool, error) {
	// Parse the current timestamp
	currentTimestamp, err := time.Parse(constants.TIMSTAMP_FORMAT, currentTimestampStr)
	if err != nil {
		return false, err
	}

	// Parse the previous timestamp
	previousTimestamp, err := time.Parse(constants.TIMSTAMP_FORMAT, previousTimestampStr)
	if err != nil {
		return false, err
	}

	// Compare timestamps
	return currentTimestamp.After(previousTimestamp), nil
}
