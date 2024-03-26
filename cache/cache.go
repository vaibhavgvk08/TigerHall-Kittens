package cache

// todo : Implement this using redis/memcache later. And add token expire logic.
var LocalCache map[string]string

func init() {
	LocalCache = make(map[string]string)
}

func Put(username, token string) {
	LocalCache[username] = token
}

func Get(username string) string {
	if d, exists := LocalCache[username]; exists {
		return d
	}
	return ""
}
