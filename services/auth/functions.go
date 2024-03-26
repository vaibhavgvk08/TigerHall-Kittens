package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vaibhavgvk08/tigerhall-kittens/cache"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GenerateJWT(username string) (string, error) {
	// token already present in cache.
	if token := cache.Get(username); token != "" {
		return token, nil
	}

	expirationTime := time.Now().Add(time.Duration(60) * time.Minute)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "TigerHall",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(constants.SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	// Put token into cache
	cache.Put(username, tokenString)
	return tokenString, nil
}

// authMiddleware is a custom middleware function that checks for authentication tokens
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isExcludedMutation(r) {
			next.ServeHTTP(w, r)
			return
		}
		// Extract authentication token from the request headers
		token := r.Header.Get("AccessToken")

		// Check if the token is missing or invalid
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Here, you can add your authentication logic, such as validating the token,
		// checking against a database, or calling an authentication service.
		// For simplicity, let's assume token validation logic here.

		// Validate the token (this is just a placeholder)
		if !isValidToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the token is valid, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

// isValidToken is a placeholder function to validate the authentication token
//func isValidToken(username, token string) bool {
//	cachedToken := cache.Get(username)
//	return token == cachedToken
//}

func isValidToken(tokenString string) bool {
	// Parse the token string into a JWT token object
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to sign the token
		return constants.SECRET_KEY, nil
	})
	if err != nil {
		return false // Token parsing error or invalid signature
	}

	// Check if the token is valid
	if token.Valid {
		return true // Token is valid
	}
	return false // Token is invalid
}

func isExcludedMutation(r *http.Request) bool {
	// Parse the GraphQL request body to extract the operation name
	// You may need to adjust this logic depending on how you handle GraphQL requests
	var requestBody struct {
		Query string `json:"query"`
	}
	//err := json.NewDecoder(r.Body).Decode(&requestBody)
	data, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(data, &requestBody)
	if err != nil {
		return false
	}
	r.Body = io.NopCloser(bytes.NewReader(data))
	if err != nil {
		return false // Unable to parse request body, include in authentication
	}

	// Check if the GraphQL operation is a login or register mutation
	return strings.Contains(requestBody.Query, "mutation") && strings.Contains(requestBody.Query, "login") ||
		strings.Contains(requestBody.Query, "mutation") && strings.Contains(requestBody.Query, "register")
}
