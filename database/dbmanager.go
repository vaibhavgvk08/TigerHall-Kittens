package database

type DBManager struct {
	client Database
}

var DBManagerInstance *DBManager

func FetchDBManager() *DBManager {
	return DBManagerInstance
}

func (obj *DBManager) CreateConnection() {

}

func (obj *DBManager) CloseConnection() {
	obj.client.Close()
}

func (obj *DBManager) Insert(entity int, doc interface{}) (result interface{}, err error) {
	dataBase, col := GetDBAndCollByEntity(entity)
	result, err = obj.client.Insert(dataBase, col, doc)
	return result, err
}

func (obj *DBManager) Update(entity int, filter, update interface{}) (result interface{}, err error) {
	dataBase, col := GetDBAndCollByEntity(entity)
	result, err = obj.client.Update(dataBase, col, filter, update)
	return result, err
}

func (obj *DBManager) Delete(entity int, query interface{}) (result interface{}, err error) {
	dataBase, col := GetDBAndCollByEntity(entity)
	result, err = obj.client.Delete(dataBase, col, query)
	return result, err
}

func (obj *DBManager) Find(entity int, query interface{}, sortOrder, skip, limit int) (result []byte, err error) {
	dataBase, col := GetDBAndCollByEntity(entity)
	result, err = obj.client.Find(dataBase, col, query, sortOrder, skip, limit)
	return result, err
}

func init() {
	DBManagerInstance = &DBManager{
		client: Connect("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.1"),
	}
}
