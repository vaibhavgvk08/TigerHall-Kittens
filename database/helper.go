package database

func GetDBAndCollByEntity(entity int) (string, string) {
	if data, found := DBMapping[entity]; found {
		return data.DBName, data.CollectionName
	}
	return "", ""
}
