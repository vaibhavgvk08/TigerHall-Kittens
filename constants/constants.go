package constants

const (
	TIGER = iota
	USER
)

const (
	SIGHTING_DISTANCE_THRESHOLD = 5
	SIGHTING_DISTANCE_UNITS     = "KMS"

	MESSAGE_QUEUE_BUFFER_SIZE = 10
)

var SECRET_KEY = []byte("gosecretkey")
