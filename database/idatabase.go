package database

// Database - todo add comments.
type Database interface {
	Insert(string, string, interface{}) (result interface{}, err error)
	Update(string, string, interface{}, interface{}) (result interface{}, err error)
	Delete(string, string, interface{}) (result interface{}, err error)
	Find(string, string, interface{}) (result []byte, err error)

	Close()
}
