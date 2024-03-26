package database

// Database - todo add comments.
type Database interface {
	Insert(string, string, interface{}) (result interface{}, err error)
	Update(string, string, interface{}, interface{}) (result interface{}, err error)
	Delete(string, string, interface{}) (result interface{}, err error)
	Find(string, string, interface{}, int, int, int) (result []byte, err error)

	Close()
}
