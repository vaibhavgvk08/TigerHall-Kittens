package cache

// todo : Implement this using redis/memcache later. And add token expire logic.
var AutheticationCache map[string]string

func init() {
	AutheticationCache = make(map[string]string)
}

func Put(username, token string) {
	AutheticationCache[username] = token
}

func Get(username string) string {
	if d, exists := AutheticationCache[username]; exists {
		return d
	}
	return ""
}
